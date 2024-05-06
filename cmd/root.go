package cmd

import (
	"dc/internal"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dc-dl",
	Short: "Download free and purchased comics",
	Long:  `A CLI app which allows you to download free and purchased comics from dhaka comics`,
	Run:   internal.SaveComic,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var dlurl string
	rootCmd.Flags().StringVar(&dlurl, "url", "", "Pass comic url you want to download")
}
