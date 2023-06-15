package main

import (
	"os"

	"github.com/Timwood0x10/sei-chain/app/params"
	"github.com/Timwood0x10/sei-chain/cmd/seid/cmd"

	"github.com/Timwood0x10/sei-chain/app"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
