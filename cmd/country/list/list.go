package countryListCmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/NikhilMJagtap/bunny-cli/api"
	"github.com/NikhilMJagtap/bunny-cli/client"
	"github.com/spf13/cobra"
)

var listCountryCommand *cobra.Command = nil

func GetListCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if listCountryCommand == nil {
		listCountryCommand = &cobra.Command{
			Use:   "list [flags]",
			Short: "List countries available with BunnyCDN",
			Long: heredoc.Doc(`
				A basic command that lists all the countries available with BunnyCDN.
			`),
			Example: heredoc.Doc(`
				$ bunny-cli country list
				$ bunny-cli country list --table
			`),
			RunE: func(cmd *cobra.Command, args []string) error {
				data, err := api.ListCountries(bunnyClient)
				if err != nil {
					return err
				}
				columns := []string{
					"IsoCode", "Name", "IsEU", "TaxRate",
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
	return listCountryCommand
}
