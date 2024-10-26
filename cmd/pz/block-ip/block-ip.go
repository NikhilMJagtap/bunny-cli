package pzBlockIPCmd

import (
	"errors"
	"net"
	"strconv"

	"github.com/NikhilMJagtap/bunny-cli/api"
	"github.com/NikhilMJagtap/bunny-cli/client"
	"github.com/spf13/cobra"
)

func validateIP(ip string) (net.IP, error) {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return nil, errors.New("Invalid IP address. Received " + ip)
	}
	return parsedIP, nil
}

var blockIPCommand *cobra.Command = nil
var unblockIPCommand *cobra.Command = nil

func GetBlockIPCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if blockIPCommand == nil {
		blockIPCommand = &cobra.Command{
			Use:   "block-ip pull-zone-id ip",
			Short: "Block an IP from accessing a Pull Zone",
			Long:  "Blocks the specific IP from accessing the Pull Zone",
			Args:  cobra.ExactArgs(2),
			RunE: func(cmd *cobra.Command, args []string) error {
				pullZoneId, err := strconv.Atoi(args[0])
				if err != nil {
					return errors.New("Pull Zone ID must be an integer. Received " + args[0])
				}
				ip, err := validateIP(args[1])
				if err != nil {
					return err
				}
				_, err = api.BlockIP(bunnyClient, uint64(pullZoneId), ip)
				if err != nil {
					return err
				}
				bunnyClient.HandleCommandOutput(cmd, "Blocked IP successfully", []string{})
				return nil
			},
		}
	}
	return blockIPCommand
}

func GetUnblockIPCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if unblockIPCommand == nil {
		unblockIPCommand = &cobra.Command{
			Use:   "unblock-ip pull-zone-id ip",
			Short: "Unblock an IP from accessing a Pull Zone",
			Long:  "Unblocks the specific IP from accessing the Pull Zone",
			Args:  cobra.ExactArgs(2),
			RunE: func(cmd *cobra.Command, args []string) error {
				pullZoneId, err := strconv.Atoi(args[0])
				if err != nil {
					return errors.New("Pull Zone ID must be an integer. Received " + args[0])
				}
				ip, err := validateIP(args[1])
				if err != nil {
					return err
				}
				_, err = api.UnblockIP(bunnyClient, uint64(pullZoneId), ip)
				if err != nil {
					return err
				}
				bunnyClient.HandleCommandOutput(cmd, "Unblocked IP successfully", []string{})
				return nil
			},
		}
	}
	return unblockIPCommand
}
