package country

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/NikhilMJagtap/bunny-cli/client"
	countryListCmd "github.com/NikhilMJagtap/bunny-cli/cmd/country/list"
	"github.com/spf13/cobra"
)

var countryCommand *cobra.Command = nil

func GetCountryCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if countryCommand == nil {
		countryCommand = &cobra.Command{
			Use:   "country <command> [flags]",
			Short: "Manage country for BunnyCDN",
			Long: heredoc.Doc(`
				A group of commands to manage countries for BunnyCDN.
			`),
			Example: heredoc.Doc(`
				$ bunny-cli country --help
			`),
			GroupID: "geo",
		}
	}
	countryCommand.AddCommand(countryListCmd.GetListCommand(bunnyClient))
	return countryCommand
}
