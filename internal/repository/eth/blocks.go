package eth

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

func (ethClient *EthereumConn) GetCurrentEthereumBlockNumber() (int64, error) {
	contextLocal, cancel := context.WithTimeout(context.Background(), time.Duration(100)*time.Second)
	defer cancel()

	headers, err := ethClient.HeaderByNumber(contextLocal, nil)
	if err != nil {
		ethClient.Log.Error().Caller().Err(err).Msg("")
		return 0, err
	}
	return headers.Number.Int64(), nil
}

func (ethClient *EthereumConn) GetFilteredLogs(fromBlock int64, toBlock int64, tokenAddress common.Address, topics [][]common.Hash) ([]types.Log, error) {

	contextLocal, cancel := context.WithTimeout(context.Background(), time.Duration(100)*time.Second)
	defer cancel()

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{tokenAddress},
		Topics:    topics,
	}

	logs, err := ethClient.FilterLogs(contextLocal, query)
	if err != nil {
		return nil, err
	}
	return logs, nil
}
