package keeper

import (
	"github.com/Timwood0x10/sei-chain/x/epoch/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AfterEpochEnd(ctx sdk.Context, epoch types.Epoch) {
	k.hooks.AfterEpochEnd(ctx, epoch)
}

func (k Keeper) BeforeEpochStart(ctx sdk.Context, epoch types.Epoch) {
	k.hooks.BeforeEpochStart(ctx, epoch)
}
