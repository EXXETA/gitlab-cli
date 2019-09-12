package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/EXXETA/gitlab-cli/service"
	"github.com/EXXETA/gitlab-cli/util"
)

var (
	projectID string
)

func init() {
	projectGetCmd.Flags().StringVarP(&projectID, "id", "i", "", "project id or URL-encoded path of the project.")
	projectGetCmd.MarkFlagRequired("id")
	projectCmd.AddCommand(projectGetCmd)
}

var projectGetCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Get a project",
	Long:    "Get a prject by id",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := service.NewProjectService(service.GitlabClientImpl).GetProjectByID(projectID)
		if err != nil {
			fmt.Println(err)
		} else {
			util.PrintPrettyJSON(result)
		}
	},
}
