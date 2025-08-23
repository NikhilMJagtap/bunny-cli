package pzPurgeCmd

import (
	"errors"
	"strconv"

	"github.com/NikhilMJagtap/bunny-cli/api"
	"github.com/NikhilMJagtap/bunny-cli/client"
	"github.com/spf13/cobra"
)

var purgeCommand *cobra.Command = nil

func GetPurgeCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if purgeCommand == nil {
		purgeCommand = &cobra.Command{
			Use:   "purge pull-zone-id",
			Short: "Purge Pull Zone cache",
			Long:  "Purge a Pull Zone cache using BunnyCDN API. Optionally, you can pass the CacheTag header to purge specific cache tags.",
			Args:  cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				pullZoneId, err := strconv.Atoi(args[0])
				if err != nil {
					return errors.New("Pull Zone ID must be an integer. Received " + args[0] + " " + strconv.Itoa(pullZoneId))
				}
				cacheTag := cmd.Flags().Lookup("cache-tag").Value.String()
				_, err = api.PurgePullZone(bunnyClient, uint64(pullZoneId), cacheTag)
				if err != nil {
					return err
				}
				bunnyClient.HandleCommandOutput(cmd, "Purge request sent successfully", []string{})
				return nil
			},
		}
	}
	purgeCommand.Flags().String(
		"cache-tag",
		"",
		"The Cache Tag to purge. If a value is provided, cached content with the header CacheTag of this value will be purged.",
	)
	return purgeCommand
}
