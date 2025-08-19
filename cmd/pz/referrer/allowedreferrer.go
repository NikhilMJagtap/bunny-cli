package pzReferrerCmd

import (
	"errors"
	"strconv"

	"github.com/NikhilMJagtap/bunny-cli/api"
	"github.com/NikhilMJagtap/bunny-cli/client"
	"github.com/NikhilMJagtap/bunny-cli/utils"
	"github.com/spf13/cobra"
)

var allowedReferrerCommand *cobra.Command = nil

func GetAllowedReferrerCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if allowedReferrerCommand == nil {
		allowedReferrerCommand = &cobra.Command{
			Use:     "allowed-referrer {add|remove} pull-zone-id referrer",
			Short:   "Adds/Removes allowed referrer to a Pull Zone",
			Long:    "Adds/Removes allowed referrer to a Pull Zone",
			Args:    cobra.ExactArgs(3),
			PreRunE: utils.AddRemoveValidator,
			RunE: func(cmd *cobra.Command, args []string) error {
				pullZoneId, err := strconv.Atoi(args[1])
				if err != nil {
					return errors.New("Pull Zone ID must be an integer. Received " + args[1])
				}
				referrer := args[2]
				switch args[0] {
				case "add":
					_, err = api.AddAllowedReferrer(bunnyClient, uint64(pullZoneId), referrer)
					if err != nil {
						return err
					}
					bunnyClient.HandleCommandOutput(cmd, "Allowed referrer added successfully", []string{})
				case "remove":
					_, err = api.RemoveAllowedReferrer(bunnyClient, uint64(pullZoneId), referrer)
					if err != nil {
						return err
					}
					bunnyClient.HandleCommandOutput(cmd, "Allowed referrer removed successfully", []string{})
				}
				return nil
			},
		}
	}
	return allowedReferrerCommand
}
