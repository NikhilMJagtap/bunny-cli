package region

import (
	"github.com/NikhilMJagtap/bunny-cli/client"
	regionListCmd "github.com/NikhilMJagtap/bunny-cli/cmd/region/list"
	"github.com/spf13/cobra"
)

var regionCommand *cobra.Command = nil

func GetRegionCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if regionCommand == nil {
		regionCommand = &cobra.Command{
			Use:     "region <command> [flags]",
			Short:   "Manage regions for BunnyCDN",
			Long:    "",
			GroupID: "region",
		}
	}
	regionCommand.AddCommand(regionListCmd.GetListCommand(bunnyClient))
	return regionCommand
}
