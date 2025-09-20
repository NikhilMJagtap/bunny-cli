package cmd

import (
	"os"

	"github.com/NikhilMJagtap/bunny-cli/client"
	"github.com/NikhilMJagtap/bunny-cli/cmd/country"
	"github.com/NikhilMJagtap/bunny-cli/cmd/pz"
	"github.com/NikhilMJagtap/bunny-cli/cmd/region"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var rootCmd = &cobra.Command{
	Use:   "bunny-cli <command> [flags]",
	Short: "CLI tool for Bunny.net",
	Long:  "CLI tool for Bunny.net",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logLevel, err := cmd.Flags().GetString("log-level")
		if err != nil {
			log.Error("Failed to read log level. Using default log level.")
			return
		}
		level, err := log.ParseLevel(logLevel)
		if err != nil {
			log.Error("Invalid log level: " + logLevel + ". Using default log level.")
			return
		}
		log.SetLevel(level)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func AddGroups() {
	rootCmd.AddGroup(&cobra.Group{
		ID:    "geo",
		Title: "Region & Country",
	})
	rootCmd.AddGroup(&cobra.Group{
		ID:    "pullzone",
		Title: "Pull Zone",
	})
}

func init() {
	AddGroups()
	rootCmd.SetOut(os.Stdout)
	rootCmd.PersistentFlags().Bool("table", false, "prints the results as a table")
	rootCmd.PersistentFlags().String(
		"log-level",
		log.WarnLevel.String(),
		"set the log level. Allowed values are: debug, info, warn, error. The default is warn.",
	)
	bunnyClient := client.GetBunnyClient()
	rootCmd.AddCommand(region.GetRegionCommand(bunnyClient))
	rootCmd.AddCommand(country.GetCountryCommand(bunnyClient))
	rootCmd.AddCommand(pz.GetPZCommand(bunnyClient))
	log.SetFormatter(
		&log.TextFormatter{
			DisableTimestamp: true,
		},
	)
}

func GetRootCommand() *cobra.Command {
	return rootCmd
}
