package pz

import (
	"github.com/NikhilMJagtap/bunny-cli/client"
	pzBlockIPCmd "github.com/NikhilMJagtap/bunny-cli/cmd/pz/block-ip"
	pzGetCmd "github.com/NikhilMJagtap/bunny-cli/cmd/pz/get"
	pzListCmd "github.com/NikhilMJagtap/bunny-cli/cmd/pz/list"
	"github.com/spf13/cobra"
)

var pzCommand *cobra.Command = nil

func GetPZCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if pzCommand == nil {
		pzCommand = &cobra.Command{
			Use:   "pz <command> [flags]",
			Short: "Manage Pull Zones for BunnyCDN",
			Long:  "",
		}
	}
	pzCommand.AddCommand(pzListCmd.GetListCommand(bunnyClient))
	pzCommand.AddCommand(pzGetCmd.GetGetPZCommand(bunnyClient))
	pzCommand.AddCommand(pzBlockIPCmd.GetBlockIPCommand(bunnyClient))
	pzCommand.AddCommand(pzBlockIPCmd.GetUnblockIPCommand(bunnyClient))
	return pzCommand
}
