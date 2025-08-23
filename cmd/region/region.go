package region

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/NikhilMJagtap/bunny-cli/client"
	regionListCmd "github.com/NikhilMJagtap/bunny-cli/cmd/region/list"
	"github.com/spf13/cobra"
)

var regionCommand *cobra.Command = nil

func GetRegionCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if regionCommand == nil {
		regionCommand = &cobra.Command{
			Use:   "region <command> [flags]",
			Short: "Manage regions for BunnyCDN",
			Long: heredoc.Doc(`
				A group of commands to manage regions for BunnyCDN.
			`),
			Example: heredoc.Doc(`
				$ bunny-cli region --help
			`),
			GroupID: "geo",
		}
	}
	regionCommand.AddCommand(regionListCmd.GetListCommand(bunnyClient))
	return regionCommand
}
