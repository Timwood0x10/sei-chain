package query_test

import (
	"testing"

	keepertest "github.com/Timwood0x10/sei-chain/testutil/keeper"
	"github.com/Timwood0x10/sei-chain/x/dex/keeper/query"
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetOrderCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wrapper := query.KeeperWrapper{Keeper: keeper}
	wctx := sdk.WrapSDKContext(ctx)
	keeper.SetOrderCount(ctx, keepertest.TestContract, keepertest.TestPair.PriceDenom, keepertest.TestPair.AssetDenom, types.PositionDirection_LONG, sdk.NewDec(1), 5)
	price := sdk.NewDec(1)
	query := types.QueryGetOrderCountRequest{
		ContractAddr:      keepertest.TestContract,
		PriceDenom:        keepertest.TestPriceDenom,
		AssetDenom:        keepertest.TestAssetDenom,
		PositionDirection: types.PositionDirection_LONG,
		Price:             &price,
	}
	resp, err := wrapper.GetOrderCount(wctx, &query)
	require.Nil(t, err)
	require.Equal(t, uint64(5), resp.Count)
}
