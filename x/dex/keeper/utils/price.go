package utils

import (
	"github.com/Timwood0x10/sei-chain/x/dex/exchange"
	"github.com/Timwood0x10/sei-chain/x/dex/keeper"
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func SetPriceStateFromExecutionOutcome(
	ctx sdk.Context,
	keeper *keeper.Keeper,
	contractAddr types.ContractAddress,
	pair types.Pair,
	outcome exchange.ExecutionOutcome,
) {
	if outcome.TotalQuantity.IsZero() {
		return
	}

	avgPrice := outcome.TotalNotional.Quo(outcome.TotalQuantity)
	priceState := types.Price{
		Pair:                       &pair,
		Price:                      avgPrice,
		SnapshotTimestampInSeconds: uint64(ctx.BlockTime().Unix()),
	}
	keeper.SetPriceState(ctx, priceState, string(contractAddr))
}
