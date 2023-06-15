package keeper

import (
	"github.com/Timwood0x10/sei-chain/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) IsVoteTarget(ctx sdk.Context, denom string) bool {
	_, err := k.GetVoteTarget(ctx, denom)
	return err == nil
}

func (k Keeper) GetVoteTargets(ctx sdk.Context) (voteTargets []string) {
	k.IterateVoteTargets(ctx, func(denom string, denomInfo types.Denom) bool {
		voteTargets = append(voteTargets, denom)
		return false
	})

	return voteTargets
}
