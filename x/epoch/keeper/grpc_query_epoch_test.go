package keeper_test

import (
	"testing"

	testkeeper "github.com/Timwood0x10/sei-chain/testutil/keeper"
	"github.com/Timwood0x10/sei-chain/x/epoch/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestEpochQuery(t *testing.T) {
	keeper, ctx := testkeeper.EpochKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	epoch := types.DefaultEpoch()
	keeper.SetEpoch(ctx, epoch)

	response, err := keeper.Epoch(wctx, &types.QueryEpochRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryEpochResponse{Epoch: epoch}, response)
}
