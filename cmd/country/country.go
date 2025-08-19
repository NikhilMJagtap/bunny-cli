package country

import (
	"github.com/NikhilMJagtap/bunny-cli/client"
	countryListCmd "github.com/NikhilMJagtap/bunny-cli/cmd/country/list"
	"github.com/spf13/cobra"
)

var countryCommand *cobra.Command = nil

func GetCountryCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if countryCommand == nil {
		countryCommand = &cobra.Command{
			Use:     "country <command> [flags]",
			Short:   "Manage country for BunnyCDN",
			Long:    "",
			GroupID: "geo",
		}
	}
	countryCommand.AddCommand(countryListCmd.GetListCommand(bunnyClient))
	return countryCommand
}
