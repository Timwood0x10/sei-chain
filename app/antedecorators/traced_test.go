package antedecorators_test

import (
	"testing"

	"github.com/Timwood0x10/sei-chain/app/antedecorators"
	"github.com/Timwood0x10/sei-chain/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestTracedDecorator(t *testing.T) {
	output = ""
	anteDecorators := []sdk.AnteFullDecorator{
		sdk.DefaultWrappedAnteDecorator(FakeAnteDecoratorOne{}),
		sdk.DefaultWrappedAnteDecorator(FakeAnteDecoratorTwo{}),
		sdk.DefaultWrappedAnteDecorator(FakeAnteDecoratorThree{}),
	}
	tracedDecorators := utils.Map(anteDecorators, func(d sdk.AnteFullDecorator) sdk.AnteFullDecorator {
		return sdk.DefaultWrappedAnteDecorator(antedecorators.NewTracedAnteDecorator(d, nil))
	})
	chainedHandler, _ := sdk.ChainAnteDecorators(tracedDecorators...)
	chainedHandler(sdk.NewContext(nil, tmproto.Header{}, false, nil), FakeTx{}, false)
	require.Equal(t, "onetwothree", output)
}
