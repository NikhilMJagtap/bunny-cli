package regionListCmd

import (
    "github.com/NikhilMJagtap/bunny-cli/api"
    "github.com/NikhilMJagtap/bunny-cli/client"
    "github.com/spf13/cobra"
)

var listRegionCommand *cobra.Command = nil

func GetListCommand(bunnyClient *client.BunnyClient) *cobra.Command {
    if listRegionCommand == nil {
        listRegionCommand = &cobra.Command{
            Use:   "list [flags]",
            Short: "List regions available with BunnyCDN",
            Long:  "",
            RunE: func(cmd *cobra.Command, args []string) error {
                data, err := api.ListRegions(bunnyClient)
                if err != nil {
                    return err
                }
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
