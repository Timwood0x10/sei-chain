package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Timwood0x10/sei-chain/testutil/keeper"
	"github.com/Timwood0x10/sei-chain/x/epoch/keeper"
	"github.com/Timwood0x10/sei-chain/x/epoch/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EpochKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
