package mint

import (
	"github.com/Timwood0x10/sei-chain/x/mint/keeper"
	"github.com/Timwood0x10/sei-chain/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func HandleUpdateMinterProposal(ctx sdk.Context, k *keeper.Keeper, p *types.UpdateMinterProposal) error {
	err := types.ValidateMinter(*p.Minter)
	if err != nil {
		return err
	}
	k.SetMinter(ctx, *p.Minter)
	return nil
}
