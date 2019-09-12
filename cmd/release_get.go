package cmd

import (
	"fmt"

	"github.com/EXXETA/gitlab-cli/service"
	"github.com/EXXETA/gitlab-cli/util"
	"github.com/spf13/cobra"
)

var (
	tagName string
)

func init() {
	releaseGetCmd.Flags().StringVarP(&tagName, "tag-name", "t", "", "tag name")

	releaseGetCmd.MarkFlagRequired("tag-name")
	releaseCmd.AddCommand(releaseGetCmd)
}

var releaseGetCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Get a release",
	Long:    "Get a release by project id and tag name",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := service.NewReleaseService(service.GitlabClientImpl).GetReleaseByProjectIDAndTagName(projectID, tagName)
		if err != nil {
			fmt.Println(err)
		} else {
			util.PrintPrettyJSON(result)
		}
	},
}
