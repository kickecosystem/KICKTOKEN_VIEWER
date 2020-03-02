package types

import "github.com/shopspring/decimal"

type Balances struct {
	AmountLiquid decimal.Decimal
	AmountFrozen decimal.Decimal
}
