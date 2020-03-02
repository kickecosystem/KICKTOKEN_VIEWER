package csv

import (
	"encoding/csv"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
	"kicktoken_viewer/internal/types"
	"os"
	"strconv"
)

func GenerateBalancesCsv(addressesBalances map[string]types.Balances, log *zerolog.Logger) error {
	fileCsv, err := os.Create("kicktoken_viewer.csv")
	if err != nil {
		return err
	}
	defer func() {
		err := fileCsv.Close()
		if err != nil {
			log.Error().Caller().Err(err).Msg("")
		}
	}()

	// csv writer
	writer := csv.NewWriter(fileCsv)
	defer writer.Flush()

	// csv titles
	err = writer.Write([]string{"Address", "Total KickTokens", "Frozen KickToken", "Liquid KickTokens"})
	if err != nil {
		return err
	}

	// balances
	for address, balances := range addressesBalances {
		err := writer.Write([]string{"https://etherscan.io/address/" + address, balances.AmountLiquid.Add(balances.AmountFrozen).String(), balances.AmountFrozen.String(), balances.AmountLiquid.String()})
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateTotalSupplyCsv(log *zerolog.Logger, amountFrozenTotal decimal.Decimal, amountLiquidTotal decimal.Decimal, totalAddressesCount int) error {
	fileCsv, err := os.Create("kicktoken_total.csv")
	if err != nil {
		return err
	}
	defer func() {
		err := fileCsv.Close()
		if err != nil {
			log.Error().Caller().Err(err).Msg("")
		}
	}()

	// csv writer
	writer := csv.NewWriter(fileCsv)
	defer writer.Flush()

	// csv titles
	err = writer.Write([]string{"Total address", "Total KickTokens", "Total Frozen KickToken", "Total Liquid KickTokens"})
	if err != nil {
		return err
	}

	// balances
	err = writer.Write([]string{strconv.Itoa(totalAddressesCount), amountFrozenTotal.Add(amountLiquidTotal).String(), amountFrozenTotal.String(), amountLiquidTotal.String()})
	if err != nil {
		return err
	}

	return nil
}
