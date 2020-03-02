package scanner

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/schollz/progressbar/v2"
	"kicktoken_viewer/internal/repository/eth"
	"kicktoken_viewer/internal/repository/sqlite"
	"sync"
)

func TxMethodFinder(ethClient *eth.EthereumConn, sqliteClient *sqlite.Client) error {

	fmt.Println("START: KickToken tx method finder")

	txs, err := sqliteClient.GetEthBcTxEmptyMethod()
	if err != nil {
		return err
	}

	if len(txs) == 0 {
		fmt.Println("DONE: KickToken tx method finder")
		return nil
	}

	bar := progressbar.New(len(txs))

	var wg sync.WaitGroup

	if len(txs) < 1000 {
		wg.Add(1)
		go updateTrMethod(txs, ethClient, sqliteClient, &wg, bar)
	} else {
		wg.Add(10)
		go updateTrMethod(txs[:int(float64(len(txs))*0.1)], ethClient, sqliteClient, &wg, bar)
		go updateTrMethod(txs[int(float64(len(txs))*0.1):int(float64(len(txs))*0.2)], ethClient, sqliteClient, &wg, bar)
		go updateTrMethod(txs[int(float64(len(txs))*0.2):int(float64(len(txs))*0.3)], ethClient, sqliteClient, &wg, bar)
		go updateTrMethod(txs[int(float64(len(txs))*0.3):int(float64(len(txs))*0.4)], ethClient, sqliteClient, &wg, bar)
		go updateTrMethod(txs[int(float64(len(txs))*0.4):int(float64(len(txs))*0.5)], ethClient, sqliteClient, &wg, bar)

		go updateTrMethod(txs[int(float64(len(txs))*0.5):int(float64(len(txs))*0.6)], ethClient, sqliteClient, &wg, bar)
		go updateTrMethod(txs[int(float64(len(txs))*0.6):int(float64(len(txs))*0.7)], ethClient, sqliteClient, &wg, bar)
		go updateTrMethod(txs[int(float64(len(txs))*0.7):int(float64(len(txs))*0.8)], ethClient, sqliteClient, &wg, bar)
		go updateTrMethod(txs[int(float64(len(txs))*0.8):int(float64(len(txs))*0.9)], ethClient, sqliteClient, &wg, bar)
		go updateTrMethod(txs[int(float64(len(txs))*0.9):], ethClient, sqliteClient, &wg, bar)
	}

	wg.Wait()

	err = bar.Finish()
	if err != nil {
		return err
	}
	fmt.Println("")

	fmt.Println("DONE: KickToken tx method finder")

	return nil
}

func updateTrMethod(txs []string, ethClient *eth.EthereumConn, sqliteClient *sqlite.Client, wg *sync.WaitGroup, bar *progressbar.ProgressBar) {
	for _, txHash := range txs {

		txMethod, err := ethClient.GetTransactionMethod(common.HexToHash(txHash))
		if err != nil {
			ethClient.Log.Fatal().Err(err).Str("txHash", txHash).Str("txMethod", txMethod).Caller().Msg("")
		}

		err = sqliteClient.UpdateTxMethod(txHash, txMethod)
		if err != nil {
			ethClient.Log.Fatal().Err(err).Str("txHash", txHash).Str("txMethod", txMethod).Caller().Msg("")
		}

		err = bar.Add(1)
		if err != nil {
			ethClient.Log.Fatal().Err(err).Str("txHash", txHash).Str("txMethod", txMethod).Caller().Msg("")
		}
	}

	wg.Done()
}
