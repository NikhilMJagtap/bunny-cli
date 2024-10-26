package pzGetCmd

import (
	"errors"
	"strconv"

	"github.com/NikhilMJagtap/bunny-cli/api"
	"github.com/NikhilMJagtap/bunny-cli/client"
	"github.com/spf13/cobra"
)

var getPZCommand *cobra.Command = nil

func GetGetPZCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if getPZCommand == nil {
		getPZCommand = &cobra.Command{
			Use:   "get pull_zone_id [flags]",
			Short: "Get specific Pull Zone from BunnyCDN",
			Long:  "",
			Args:  cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				pullZoneId, err := strconv.Atoi(args[0])
				if err != nil {
					return errors.New("Pull Zone ID must be an integer. Received " + args[0])
				}
				data, err := api.GetPullZone(bunnyClient, uint64(pullZoneId))
				if err != nil {
					return err
				}
				columns := []string{
					"Id", "Name", "OriginUrl",
				}
				err = bunnyClient.HandleCommandOutput(
					cmd,
					data,
					columns,
				)
				return err
			},
		}
	}

	return getPZCommand
}
