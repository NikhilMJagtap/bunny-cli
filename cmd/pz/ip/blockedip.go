package pzBlockIPCmd

import (
	"errors"
	"strconv"

	"github.com/NikhilMJagtap/bunny-cli/api"
	"github.com/NikhilMJagtap/bunny-cli/client"
	utils "github.com/NikhilMJagtap/bunny-cli/utils"
	"github.com/spf13/cobra"
)

var blockIPCommand *cobra.Command = nil

func GetBlockIPCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if blockIPCommand == nil {
		blockIPCommand = &cobra.Command{
			Use:     "blocked-ip {add|remove} pull-zone-id ip",
			Short:   "Adds/Removes blocked IP from accessing a Pull Zone",
			Long:    "Adds/Removes blocked IP from accessing a Pull Zone. Wildcard can be used for IPv4.",
			PreRunE: utils.AddRemoveValidator,
			Args:    cobra.ExactArgs(3),
			RunE: func(cmd *cobra.Command, args []string) error {
				pullZoneId, err := strconv.Atoi(args[1])
				if err != nil {
					return errors.New("Pull Zone ID must be an integer. Received " + args[1])
				}
				ip, err := utils.ValidateIP(args[2])
				if err != nil {
					return err
				}
				switch args[0] {
				case "add":
					_, err = api.BlockIP(bunnyClient, uint64(pullZoneId), ip)
					if err != nil {
						return err
					}
					bunnyClient.HandleCommandOutput(cmd, "Added blocked IP successfully", []string{})
				case "remove":
					_, err = api.UnblockIP(bunnyClient, uint64(pullZoneId), ip)
					if err != nil {
						return err
					}
					bunnyClient.HandleCommandOutput(cmd, "Removed blocked IP successfully", []string{})
				}
				return nil
			},
		}
	}
	return blockIPCommand
}
