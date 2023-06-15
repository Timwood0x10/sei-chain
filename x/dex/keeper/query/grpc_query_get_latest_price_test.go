package query_test

import (
	"testing"
	"time"

	keepertest "github.com/Timwood0x10/sei-chain/testutil/keeper"
	"github.com/Timwood0x10/sei-chain/x/dex/keeper/query"
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetLatestPrice(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	keepertest.SeedPriceSnapshot(ctx, keeper, "100", 1)
	keepertest.SeedPriceSnapshot(ctx, keeper, "101", 2)
	keepertest.SeedPriceSnapshot(ctx, keeper, "99", 3)

	ctx = ctx.WithBlockTime(time.Unix(4, 0))
	wctx := sdk.WrapSDKContext(ctx)
	wrapper := query.KeeperWrapper{Keeper: keeper}
	resp, err := wrapper.GetLatestPrice(wctx, &types.QueryGetLatestPriceRequest{
		ContractAddr: keepertest.TestContract,
		PriceDenom:   keepertest.TestPair.PriceDenom,
		AssetDenom:   keepertest.TestPair.AssetDenom,
	})
	require.Nil(t, err)
	require.Equal(t, sdk.MustNewDecFromStr("99"), resp.Price.Price)
}
