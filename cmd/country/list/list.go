package countryListCmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/NikhilMJagtap/bunny-cli/api"
	"github.com/NikhilMJagtap/bunny-cli/client"
	log "github.com/sirupsen/logrus"
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
				log.Debug("Running list country command")
				data, err := api.ListCountries(bunnyClient)
				if err != nil {
					log.Error("Failed to list countries: " + err.Error())
					return err
				}
				log.Debug("Country list fetched successfully.")
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
