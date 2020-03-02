package service

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
	"kicktoken_viewer/internal/repository/csv"
	"kicktoken_viewer/internal/repository/eth"
	"kicktoken_viewer/internal/repository/sqlite"
	"kicktoken_viewer/internal/scanner"
)

func GetBalances(log *zerolog.Logger, ethUrl string) {

	sqliteClient, err := sqlite.InitSqliteClient(log)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	defer func() {
		err = sqliteClient.Close()
		if err != nil {
			log.Error().Err(err).Msg("")
		}
	}()

	err = sqliteClient.MigrateUp()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	ethConn, err := eth.InitEthereum(log, ethUrl)
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("can't init Ethereum")
	}

	err = sqliteClient.SyncLastBlock(8433406)
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("can't sync last blocks")
	}

	tokenAddress := "0xc12d1c73ee7dc3615ba4e37e4abfdbddfa38907e"
	err = scanner.TxKickCatcher(ethConn, sqliteClient, tokenAddress, 300)
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("")
	}

	err = scanner.TxMethodFinder(ethConn, sqliteClient)
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("")
	}

	fmt.Println("START: CSV report generating")

	txAmount, err := sqliteClient.GetTxCount()
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("")
	}

	allTrxLogs, err := sqliteClient.GetAllTransactions(int(txAmount))
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("")
	}

	addressesBalances := calcBalances(allTrxLogs, tokenAddress)

	err = csv.GenerateBalancesCsv(addressesBalances, log)
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("")
	}

	var amountFrozenTotal = decimal.NewFromFloat(0)
	var amountLiquidTotal = decimal.NewFromFloat(0)
	for _, balances := range addressesBalances {
		amountFrozenTotal = amountFrozenTotal.Add(balances.AmountFrozen)
		amountLiquidTotal = amountLiquidTotal.Add(balances.AmountLiquid)
	}
	err = csv.GenerateTotalSupplyCsv(log, amountFrozenTotal, amountLiquidTotal, len(addressesBalances))
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("")
	}

	fmt.Println("GENERATED FILES: kicktoken_viewer.csv and kicktoken_total.csv")
}
