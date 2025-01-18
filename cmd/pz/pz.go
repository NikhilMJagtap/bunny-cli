package pz

import (
	"github.com/NikhilMJagtap/bunny-cli/client"
	pzGetCmd "github.com/NikhilMJagtap/bunny-cli/cmd/pz/get"
	pzHostnameCmd "github.com/NikhilMJagtap/bunny-cli/cmd/pz/hostname"
	pzBlockIPCmd "github.com/NikhilMJagtap/bunny-cli/cmd/pz/ip"
	pzListCmd "github.com/NikhilMJagtap/bunny-cli/cmd/pz/list"
	pzReferrerCmd "github.com/NikhilMJagtap/bunny-cli/cmd/pz/referrer"
	"github.com/spf13/cobra"
)

var pzCommand *cobra.Command = nil

func GetPZCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if pzCommand == nil {
		pzCommand = &cobra.Command{
			Use:     "pz <command> [flags]",
			Short:   "Manage Pull Zones for BunnyCDN",
			Long:    "",
			GroupID: "pullzone",
		}
	}
	pzCommand.AddCommand(pzListCmd.GetListCommand(bunnyClient))
	pzCommand.AddCommand(pzGetCmd.GetGetPZCommand(bunnyClient))
	pzCommand.AddCommand(pzBlockIPCmd.GetBlockIPCommand(bunnyClient))
	pzCommand.AddCommand(pzReferrerCmd.GetAllowedReferrerCommand(bunnyClient))
	pzCommand.AddCommand(pzReferrerCmd.GetBlockedReferrerCommand(bunnyClient))
	pzCommand.AddCommand(pzHostnameCmd.GetHostnameCommand(bunnyClient))
	pzCommand.AddCommand(pzHostnameCmd.GetForceSSLCommand(bunnyClient))
	return pzCommand
}
