package migrations

import (
	"github.com/Timwood0x10/sei-chain/x/dex/keeper"
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func V12ToV13(ctx sdk.Context, dexkeeper keeper.Keeper) error {
	// This isn't the cleanest migration since it could potentially revert any dex params we have changed
	// but we haven't, so we'll just do this.
	defaultParams := types.DefaultParams()
	dexkeeper.SetParams(ctx, defaultParams)
	return nil
}
