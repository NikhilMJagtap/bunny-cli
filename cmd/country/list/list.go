package countryListCmd

import (
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
			Long:  "A basic command that lists all the countries available with BunnyCDN.",
			Example: `
bunny-cli country list
            `,
			RunE: func(cmd *cobra.Command, args []string) error {
				data, err := api.ListRegions(bunnyClient)
				if err != nil {
					return err
				}
				columns := []string{
					"Id", "CountryCode", "Name", "ContinentCode", "RegionCode", "PricePerGigabyte",
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
