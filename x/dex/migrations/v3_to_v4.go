package migrations

import (
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

func PriceSnapshotUpdate(ctx sdk.Context, paramStore paramtypes.Subspace) error {
	err := migratePriceSnapshotParam(ctx, paramStore)
	return err
}

func migratePriceSnapshotParam(ctx sdk.Context, paramStore paramtypes.Subspace) error {
	defaultParams := types.Params{
		PriceSnapshotRetention: types.DefaultPriceSnapshotRetention,
	}
	paramStore.SetParamSet(ctx, &defaultParams)
	return nil
}
