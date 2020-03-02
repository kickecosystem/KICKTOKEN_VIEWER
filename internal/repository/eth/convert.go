package eth

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"kicktoken_viewer/internal/repository/sqlite"
	"math/big"
)

func EventLogsToMap(eventLogs []types.Log, events map[string]string) map[string][]sqlite.TxFromBc {

	txsSummed := make(map[string][]sqlite.TxFromBc)

	for _, vLog := range eventLogs {

		tokenAmountPenny := new(big.Int).SetBytes(vLog.Data)

		if tokenAmountPenny.Cmp(big.NewInt(0)) == 0 {
			continue
		}

		addressFirst, addressSecond := getAddressesFromLog(vLog.Topics)

		txHex := vLog.TxHash.Hex()
		bc := sqlite.TxFromBc{
			AddressFirst:  addressFirst,
			AddressSecond: addressSecond,
			Amount:        decimal.NewFromBigInt(tokenAmountPenny, -8),
			BlockNumber:   vLog.BlockNumber,
			EventName:     events[vLog.Topics[0].String()],
		}

		txsSummed[txHex] = append(txsSummed[txHex], bc)
	}
	return txsSummed
}
