package wasm

import (
	"encoding/json"

	"github.com/Timwood0x10/sei-chain/wasmbinding/bindings"
	"github.com/Timwood0x10/sei-chain/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EncodeDexPlaceOrders(rawMsg json.RawMessage, sender sdk.AccAddress) ([]sdk.Msg, error) {
	encodedPlaceOrdersMsg := bindings.PlaceOrders{}
	if err := json.Unmarshal(rawMsg, &encodedPlaceOrdersMsg); err != nil {
		return []sdk.Msg{}, types.ErrEncodeDexPlaceOrders
	}
	placeOrdersMsg := types.MsgPlaceOrders{
		Creator:      sender.String(),
		Orders:       encodedPlaceOrdersMsg.Orders,
		ContractAddr: encodedPlaceOrdersMsg.ContractAddr,
		Funds:        encodedPlaceOrdersMsg.Funds,
	}
	return []sdk.Msg{&placeOrdersMsg}, nil
}

func EncodeDexCancelOrders(rawMsg json.RawMessage, sender sdk.AccAddress) ([]sdk.Msg, error) {
	encodedCancelOrdersMsg := bindings.CancelOrders{}
	if err := json.Unmarshal(rawMsg, &encodedCancelOrdersMsg); err != nil {
		return []sdk.Msg{}, types.ErrEncodeDexCancelOrders
	}
	cancelOrdersMsg := types.MsgCancelOrders{
		Creator:       sender.String(),
		Cancellations: encodedCancelOrdersMsg.Cancellations,
		ContractAddr:  encodedCancelOrdersMsg.ContractAddr,
	}
	return []sdk.Msg{&cancelOrdersMsg}, nil
}
