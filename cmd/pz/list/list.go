package pzListCmd

import (
    "github.com/NikhilMJagtap/bunny-cli/api"
    "github.com/NikhilMJagtap/bunny-cli/client"
    "github.com/spf13/cobra"
)

var listPZCommand *cobra.Command = nil

func GetListCommand(bunnyClient *client.BunnyClient) *cobra.Command {
    if listPZCommand == nil {
        listPZCommand = &cobra.Command{
            Use:   "list [flags]",
            Short: "List Pull Zones available with BunnyCDN",
            Long:  "",
            RunE: func(cmd *cobra.Command, args []string) error {

                data, err := api.ListPullZones(bunnyClient, parseListPullZoneOpts(cmd))
                if err != nil {
                    return err
                }
                switch data.(type) {
                case map[string]interface{}:
                    data = data.(map[string]interface{})["Items"]
                }
                columns := []string{
                    "Id", "Name", "OriginUrl",
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

    listPZCommand.Flags().IntP("page", "p", 0, "Page number")
    listPZCommand.Flags().IntP("per-page", "P", 1000, "Number of Pull Zones per page")
    listPZCommand.Flags().StringP("search", "s", "", "The search term that will be used to filter the results")
    listPZCommand.Flags().Bool("include-certificate", false, "Determines if the result hostnames should contain the SSL certificate. False by default.")

    return listPZCommand
}

func parseListPullZoneOpts(cmd *cobra.Command) *api.ListPullZoneOpts {
    opts := &api.ListPullZoneOpts{}
    page, _ := cmd.Flags().GetInt("page")
    opts.Page = uint32(page)
    perPage, _ := cmd.Flags().GetInt("per-page")
    opts.PerPage = uint32(perPage)
    search, _ := cmd.Flags().GetString("search")
    opts.Search = search
    includeCertificate, _ := cmd.Flags().GetBool("include-certificate")
    opts.IncludeCertificate = includeCertificate
    return opts
}
