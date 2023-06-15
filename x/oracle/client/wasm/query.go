package wasm

import (
	oraclekeeper "github.com/Timwood0x10/sei-chain/x/oracle/keeper"
	"github.com/Timwood0x10/sei-chain/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OracleWasmQueryHandler struct {
	oracleKeeper oraclekeeper.Keeper
}

func NewOracleWasmQueryHandler(keeper *oraclekeeper.Keeper) *OracleWasmQueryHandler {
	return &OracleWasmQueryHandler{
		oracleKeeper: *keeper,
	}
}

func (handler OracleWasmQueryHandler) GetExchangeRates(ctx sdk.Context) (*types.QueryExchangeRatesResponse, error) {
	querier := oraclekeeper.NewQuerier(handler.oracleKeeper)
	c := sdk.WrapSDKContext(ctx)
	return querier.ExchangeRates(c, &types.QueryExchangeRatesRequest{})
}

func (handler OracleWasmQueryHandler) GetOracleTwaps(ctx sdk.Context, req *types.QueryTwapsRequest) (*types.QueryTwapsResponse, error) {
	querier := oraclekeeper.NewQuerier(handler.oracleKeeper)
	c := sdk.WrapSDKContext(ctx)
	return querier.Twaps(c, req)
}
