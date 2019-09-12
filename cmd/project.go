package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(projectCmd)
}

var projectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Work with Gitlab projects",
	Long:    "Work with Gitlab projects",
}
