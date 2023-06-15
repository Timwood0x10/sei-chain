package tx

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Timwood0x10/sei-chain/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
)

//nolint:deadcode,unused // I assume we'll use this later.
const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(CmdPlaceOrders())
	cmd.AddCommand(CmdCancelOrders())
	cmd.AddCommand(CmdRegisterContract())
	cmd.AddCommand(CmdRegisterPairs())
	cmd.AddCommand(CmdUnregisterContract())
	cmd.AddCommand(CmdContractDepositRent())
	cmd.AddCommand(CmdUpdatePriceTickSize())
	cmd.AddCommand(CmdUpdateQuantityTickSize())
	cmd.AddCommand(NewAddAssetProposalTxCmd())
	cmd.AddCommand(CmdUnsuspendContract())
	// this line is used by starport scaffolding # 1

	return cmd
}
