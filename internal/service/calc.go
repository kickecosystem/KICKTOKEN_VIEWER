package service

import (
	"github.com/shopspring/decimal"
	"kicktoken_viewer/internal/repository/sqlite"
	"kicktoken_viewer/internal/types"
)

func calcBalances(allTrxLogs []sqlite.TxFromDb, tokenAddress string) map[string]types.Balances {
	addressesFrozen := make(map[string]decimal.Decimal)
	addressesLiquid := make(map[string]decimal.Decimal)

	for _, trLog := range allTrxLogs {

		// frozen balances
		if trLog.EventName == "Transfer" && trLog.Method == "transferFrozenToken" {
			addressesFrozen[trLog.AddressSecond] = addressesFrozen[trLog.AddressSecond].Add(trLog.Amount)
		}
		if trLog.EventName == "MintFrozen" {
			addressesFrozen[trLog.AddressFirst] = addressesFrozen[trLog.AddressFirst].Add(trLog.Amount)
		}
		if trLog.EventName == "Freeze" {
			addressesFrozen[trLog.AddressFirst] = addressesFrozen[trLog.AddressFirst].Add(trLog.Amount)
		}
		if trLog.EventName == "Transfer" && (trLog.Method == "destroyFrozen" || trLog.Method == "transferFrozenToken") {
			addressesFrozen[trLog.AddressFirst] = addressesFrozen[trLog.AddressFirst].Sub(trLog.Amount)
		}
		if trLog.EventName == "Melt" {
			addressesFrozen[trLog.AddressFirst] = addressesFrozen[trLog.AddressFirst].Sub(trLog.Amount)
		}

		// liquid balances
		if trLog.EventName == "Transfer" && trLog.Method != "mintFrozenTokens" && trLog.Method != "mintBatchFrozenTokens" && trLog.Method != "transferFrozenToken" {
			addressesLiquid[trLog.AddressSecond] = addressesLiquid[trLog.AddressSecond].Add(trLog.Amount)
		}
		if trLog.EventName == "Transfer" && trLog.Method != "destroyFrozen" && trLog.Method != "transferFrozenToken" {
			addressesLiquid[trLog.AddressFirst] = addressesLiquid[trLog.AddressFirst].Sub(trLog.Amount)
		}

	}

	zero := decimal.NewFromFloat(0)
	for address, balanceFrozen := range addressesFrozen {
		if balanceFrozen.Cmp(zero) == 0 {
			delete(addressesFrozen, address)
		}
	}
	delete(addressesFrozen, tokenAddress)
	for address, balanceLiquid := range addressesLiquid {
		if balanceLiquid.Cmp(zero) == 0 {
			delete(addressesLiquid, address)
		}
	}
	delete(addressesLiquid, tokenAddress)

	var addressesBalances = make(map[string]types.Balances)
	for k, v := range addressesLiquid {
		balances := addressesBalances[k]
		balances.AmountLiquid = v
		addressesBalances[k] = balances
	}
	for k, v := range addressesFrozen {
		balances := addressesBalances[k]
		balances.AmountFrozen = v
		addressesBalances[k] = balances
	}

	return addressesBalances
}
