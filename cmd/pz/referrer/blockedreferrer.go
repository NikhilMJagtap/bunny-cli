package pzReferrerCmd

import (
	"errors"
	"strconv"

	"github.com/NikhilMJagtap/bunny-cli/api"
	"github.com/NikhilMJagtap/bunny-cli/client"
	"github.com/NikhilMJagtap/bunny-cli/utils"
	"github.com/spf13/cobra"
)

var blockedReferrerCommand *cobra.Command = nil

func GetBlockedReferrerCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if blockedReferrerCommand == nil {
		blockedReferrerCommand = &cobra.Command{
			Use:     "blocked-referrer {add|remove} pull-zone-id referrer",
			Short:   "Adds/Removes blocked referrer to a Pull Zone",
			Long:    "Adds/Removes blocked referrer to a Pull Zone",
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
					_, err = api.AddBlockedReferrer(bunnyClient, uint64(pullZoneId), referrer)
					if err != nil {
						return err
					}
					bunnyClient.HandleCommandOutput(cmd, "Blocked referrer added successfully", []string{})
				case "remove":
					_, err = api.RemoveBlockedReferrer(bunnyClient, uint64(pullZoneId), referrer)
					if err != nil {
						return err
					}
					bunnyClient.HandleCommandOutput(cmd, "Blocked referrer removed successfully", []string{})
				}
				return nil
			},
		}
	}
	return blockedReferrerCommand
}
