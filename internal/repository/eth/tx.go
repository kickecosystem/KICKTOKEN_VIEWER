package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

func (ethClient *EthereumConn) GetTransactionMethod(txHash common.Hash) (string, error) {

	contextLocal, cancel := context.WithTimeout(context.Background(), time.Duration(100)*time.Second)
	defer cancel()

	tx, _, err := ethClient.TransactionByHash(contextLocal, txHash)
	if err != nil {
		return "", err
	}

	txInput := tx.Data()
	methodBytes := txInput[:4]
	method, err := ethClient.TokenAbi.MethodById(methodBytes)
	if err != nil {
		return "UNDEFINED", nil
	}

	return method.Name, nil
}
