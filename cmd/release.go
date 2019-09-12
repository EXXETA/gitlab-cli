package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	releaseCmd.PersistentFlags().StringVarP(&projectID, "project-id", "i", "", "project id or URL-encoded path of the project.")

	releaseCmd.MarkPersistentFlagRequired("project-id")
	rootCmd.AddCommand(releaseCmd)
}

var releaseCmd = &cobra.Command{
	Use:     "release",
	Aliases: []string{"r"},
	Short:   "Work with Gitlab releases",
	Long:    "Work with Gitlab releases",
}
