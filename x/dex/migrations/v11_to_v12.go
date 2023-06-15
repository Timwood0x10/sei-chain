package migrations

import (
	"github.com/Timwood0x10/sei-chain/x/dex/keeper"
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func V11ToV12(ctx sdk.Context, dexkeeper keeper.Keeper) error {
	defaultParams := types.DefaultParams()
	dexkeeper.SetParams(ctx, defaultParams)
	return nil
}
