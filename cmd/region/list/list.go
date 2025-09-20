package regionListCmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/NikhilMJagtap/bunny-cli/api"
	"github.com/NikhilMJagtap/bunny-cli/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var listRegionCommand *cobra.Command = nil

func GetListCommand(bunnyClient *client.BunnyClient) *cobra.Command {
	if listRegionCommand == nil {
		listRegionCommand = &cobra.Command{
			Use:   "list [flags]",
			Short: "List regions available with BunnyCDN",
			Long: heredoc.Doc(`
				A basic command that lists all the regions available with BunnyCDN.
			`),
			Example: heredoc.Doc(`
				$ bunny-cli region list
				$ bunny-cli region list --table
			`),
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Debug("Running list region command")
				data, err := api.ListRegions(bunnyClient)
				if err != nil {
					log.Error("Failed to list regions: " + err.Error())
					return err
				}
				log.Debug("Region list fetched successfully.")
				columns := []string{
					"Id", "Name", "RegionCode", "CountryCode", "ContinentCode",
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
	return listRegionCommand
}
