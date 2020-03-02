package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"kicktoken_viewer/internal/repository/eth/token"
	"strings"
	"time"
)

type EthereumConn struct {
	*ethclient.Client
	Log      *zerolog.Logger
	TokenAbi abi.ABI
}

func InitEthereum(log *zerolog.Logger, url string) (*EthereumConn, error) {
	contextLocal, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	ethClient, err := ethclient.DialContext(contextLocal, url)
	if err != nil {
		return nil, err
	}

	tokenAbi, err := abi.JSON(strings.NewReader(token.TokenABI))
	if err != nil {
		return nil, err
	}

	return &EthereumConn{Client: ethClient, Log: log, TokenAbi: tokenAbi}, nil
}
