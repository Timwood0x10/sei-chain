package msgserver

import (
	"context"

	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateQuantityTickSize(goCtx context.Context, msg *types.MsgUpdateQuantityTickSize) (*types.MsgUpdateTickSizeResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validation such that only the user who stored the code can update tick size
	for _, tickSize := range msg.TickSizeList {
		contractAddr := tickSize.ContractAddr
		contractInfo, err := k.GetContract(ctx, contractAddr)
		if err != nil {
			return nil, err
		}

		if msg.Creator != contractInfo.Creator {
			return nil, sdkerrors.ErrUnauthorized
		}

		err = k.SetQuantityTickSizeForPair(ctx, tickSize.ContractAddr, *tickSize.Pair, tickSize.Ticksize)
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgUpdateTickSizeResponse{}, nil
}
