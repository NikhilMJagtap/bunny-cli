package cmd

import (
    "os"

    "github.com/NikhilMJagtap/bunny-cli/client"
    "github.com/NikhilMJagtap/bunny-cli/cmd/pz"
    "github.com/NikhilMJagtap/bunny-cli/cmd/region"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "bunny-cli <command> [flags]",
    Short: "CLI tool for Bunny.net",
    Long:  "CLI tool for Bunny.net",
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
    rootCmd.SetOut(os.Stdout)
    rootCmd.PersistentFlags().Bool("table", false, "prints the results as a table")
    bunnyClient := client.GetBunnyClient()
    rootCmd.AddCommand(region.GetRegionCommand(bunnyClient))
    rootCmd.AddCommand(pz.GetPZCommand(bunnyClient))

}
