package cmd

import (
	"fmt"

	"github.com/EXXETA/gitlab-cli/service"
	"github.com/EXXETA/gitlab-cli/util"
	"github.com/spf13/cobra"
)

func init() {
	projectCmd.AddCommand(projectListCmd)
}

var projectListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "Get list of projects",
	Long:    "Get list of projects",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := service.NewProjectService(service.GitlabClientImpl).GetAllProjects()
		if err != nil {
			fmt.Println(err)
		} else {
			util.PrintPrettyJSON(result)
		}
	},
}
