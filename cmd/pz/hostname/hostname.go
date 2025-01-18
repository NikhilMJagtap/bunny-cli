package pzHostnameCmd

import (
	"errors"
	"strconv"

	"github.com/NikhilMJagtap/bunny-cli/api"
	"github.com/NikhilMJagtap/bunny-cli/client"
	"github.com/NikhilMJagtap/bunny-cli/utils"
	"github.com/spf13/cobra"
)

var hostnameCommand *cobra.Command = nil

func GetHostnameCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if hostnameCommand == nil {
		hostnameCommand = &cobra.Command{
			Use:     "hostname {add|remove} pull-zone-id hostname",
			Short:   "Adds/Removes hostname to a Pull Zone",
			Long:    "Adds/Removes hostname to a Pull Zone",
			Args:    cobra.ExactArgs(3),
			PreRunE: utils.AddRemoveValidator,
			RunE: func(cmd *cobra.Command, args []string) error {
				pullZoneId, err := strconv.Atoi(args[1])
				if err != nil {
					return errors.New("Pull Zone ID must be an integer. Received " + args[1])
				}
				hostname := args[2]
				switch args[0] {
				case "add":
					_, err = api.AddHostname(bunnyClient, uint64(pullZoneId), hostname)
				case "remove":
					_, err = api.RemoveHostname(bunnyClient, uint64(pullZoneId), hostname)
				}
				if err != nil {
					return err
				}
				bunnyClient.HandleCommandOutput(cmd, "Hostname updated successfully", []string{})
				return nil
			},
		}
	}
	return hostnameCommand
}
