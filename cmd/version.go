package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version of the application
	Version = "0.0.1"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the gitlabcli",
	Long:  "All software has versions. This is gitlabcli's :)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}
