package scanner

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/schollz/progressbar/v2"
	"kicktoken_viewer/internal/repository/eth"
	"kicktoken_viewer/internal/repository/sqlite"
)

func TxKickCatcher(ethClient *eth.EthereumConn, sqliteClient *sqlite.Client, tokenAddress string, bunchBlockFilterSize int64) error {

	fmt.Println("START: KickToken tx finder")

	contractAddress := common.HexToAddress(tokenAddress)

	// actual topics
	topics := eth.GetTopics()

	// actual events
	events := eth.GetLogEvents()

	// current block number
	currentBlockNumber, err := ethClient.GetCurrentEthereumBlockNumber()
	if err != nil {
		return err
	}

	// last scanned block number
	fromBlockNumber, err := sqliteClient.GetLastBlock()
	if err != nil {
		return err
	}

	length := int(currentBlockNumber) - int(fromBlockNumber)

	if length <= 0 {
		fmt.Println("DONE: KickToken tx finder")
		return nil
	}

	bar := progressbar.New(length)

	for {

		toBlockNumber := fromBlockNumber + bunchBlockFilterSize

		if toBlockNumber > currentBlockNumber {
			toBlockNumber = currentBlockNumber
		}

		logs, err := ethClient.GetFilteredLogs(fromBlockNumber, toBlockNumber, contractAddress, topics)
		if err != nil {
			return err
		}

		txsFromBc := eth.EventLogsToMap(logs, events)

		err = sqliteClient.InsertTxs(txsFromBc)
		if err != nil {
			return err
		}

		err = sqliteClient.UpdateLastBlock(toBlockNumber)
		if err != nil {
			return err
		}

		err = bar.Add(int(toBlockNumber) - int(fromBlockNumber))
		if err != nil {
			return err
		}

		if toBlockNumber == currentBlockNumber {
			break
		}

		fromBlockNumber = toBlockNumber

	}

	err = bar.Finish()
	if err != nil {
		return err
	}
	fmt.Println("")

	fmt.Println("DONE: KickToken tx finder")

	return nil
}
