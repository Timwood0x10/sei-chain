package msgserver_test

import (
	"testing"

	keepertest "github.com/Timwood0x10/sei-chain/testutil/keeper"
	"github.com/Timwood0x10/sei-chain/x/dex/keeper/msgserver"
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestUnsuspendContract(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	keeper.SetParams(ctx, types.DefaultParams())
	keeper.SetContract(ctx, &types.ContractInfoV2{
		ContractAddr:     keepertest.TestContract,
		Creator:          keepertest.TestAccount,
		RentBalance:      types.DefaultContractUnsuspendCost,
		Suspended:        true,
		SuspensionReason: "bad",
	})
	server := msgserver.NewMsgServerImpl(*keeper)
	_, err := server.UnsuspendContract(wctx, &types.MsgUnsuspendContract{
		Creator:      keepertest.TestAccount,
		ContractAddr: keepertest.TestContract,
	})
	require.Nil(t, err)
	contract, err := keeper.GetContract(ctx, keepertest.TestContract)
	require.Nil(t, err)
	require.Equal(t, types.ContractInfoV2{
		ContractAddr: keepertest.TestContract,
		Creator:      keepertest.TestAccount,
	}, contract)
}

func TestUnsuspendContractInvalid(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	contract := types.ContractInfoV2{
		ContractAddr:     keepertest.TestContract,
		Creator:          keepertest.TestAccount,
		RentBalance:      types.DefaultContractUnsuspendCost - 1,
		Suspended:        true,
		SuspensionReason: "bad",
	}
	keeper.SetParams(ctx, types.DefaultParams())
	server := msgserver.NewMsgServerImpl(*keeper)
	keeper.SetContract(ctx, &contract)
	_, err := server.UnsuspendContract(wctx, &types.MsgUnsuspendContract{
		Creator:      keepertest.TestAccount,
		ContractAddr: keepertest.TestContract,
	})
	require.NotNil(t, err)
	gotContract, err := keeper.GetContract(ctx, keepertest.TestContract)
	require.Nil(t, err)
	require.Equal(t, contract, gotContract)

	contract.RentBalance += 1
	contract.Suspended = false
	keeper.SetContract(ctx, &contract)
	_, err = server.UnsuspendContract(wctx, &types.MsgUnsuspendContract{
		Creator:      keepertest.TestAccount,
		ContractAddr: keepertest.TestContract,
	})
	require.NotNil(t, err)
	gotContract, err = keeper.GetContract(ctx, keepertest.TestContract)
	require.Nil(t, err)
	require.Equal(t, contract, gotContract)
}
