package cmd

import (
	"fmt"

	"github.com/EXXETA/gitlab-cli/service"
	"github.com/EXXETA/gitlab-cli/util"
	"github.com/spf13/cobra"
)

func init() {
	releaseCmd.AddCommand(releaseListCmd)
}

var releaseListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "Get list of releases",
	Long:    "Get list of releases",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := service.NewReleaseService(service.GitlabClientImpl).GetAllReleasesByProjectID(projectID)
		if err != nil {
			fmt.Println(err)
		} else {
			util.PrintPrettyJSON(result)
		}
	},
}
