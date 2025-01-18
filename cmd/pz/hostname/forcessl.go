package pzHostnameCmd

import (
	"errors"
	"strconv"

	"github.com/NikhilMJagtap/bunny-cli/api"
	"github.com/NikhilMJagtap/bunny-cli/client"
	"github.com/NikhilMJagtap/bunny-cli/utils"
	"github.com/spf13/cobra"
)

var forceSSLCommand *cobra.Command = nil

func GetForceSSLCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if forceSSLCommand == nil {
		forceSSLCommand = &cobra.Command{
			Use:     "force-ssl {enable|disable} pull-zone-id hostname",
			Short:   "Enables/Disables force SSL for a Pull Zone",
			Long:    "Enables/Disables force SSL for a Pull Zone",
			Args:    cobra.ExactArgs(3),
			PreRunE: utils.EnableDisableValidator,
			RunE: func(cmd *cobra.Command, args []string) error {
				pullZoneId, err := strconv.Atoi(args[1])
				if err != nil {
					return errors.New("Pull Zone ID must be an integer. Received " + args[1])
				}
				hostname := args[2]
				_, err = api.ForceSSL(bunnyClient, uint64(pullZoneId), hostname, args[0] == "enable")
				if err != nil {
					return err
				}
				bunnyClient.HandleCommandOutput(cmd, "Force SSL updated successfully", []string{})
				return nil
			},
		}
	}
	return forceSSLCommand
}
