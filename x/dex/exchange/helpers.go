package exchange

import (
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ExecutionOutcome struct {
	TotalNotional sdk.Dec
	TotalQuantity sdk.Dec
	Settlements   []*types.SettlementEntry
	MinPrice      sdk.Dec
	MaxPrice      sdk.Dec
}

func (o *ExecutionOutcome) Merge(other *ExecutionOutcome) ExecutionOutcome {
	return ExecutionOutcome{
		TotalNotional: o.TotalNotional.Add(other.TotalNotional),
		TotalQuantity: o.TotalQuantity.Add(other.TotalQuantity),
		Settlements:   append(o.Settlements, other.Settlements...),
		MinPrice:      sdk.MinDec(o.MinPrice, other.MinPrice),
		MaxPrice:      sdk.MaxDec(o.MaxPrice, other.MaxPrice),
	}
}
