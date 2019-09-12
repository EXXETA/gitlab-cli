package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/EXXETA/gitlab-cli/model"
	"github.com/EXXETA/gitlab-cli/service"
	"github.com/EXXETA/gitlab-cli/util"
	"github.com/spf13/cobra"
)

func init() {
	releaseCmd.AddCommand(releaseCreateCmd)
}

var releaseCreateCmd = &cobra.Command{
	Use:     "create [release]",
	Aliases: []string{"c"},
	Short:   "Create a new release",
	Long:    "Create a new release for a project. The argument is the releasse object in JSON format",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// parse the release object from argument
		newRelease := &model.Release{}
		if err := json.Unmarshal([]byte(args[0]), newRelease); err != nil {
			fmt.Printf("ERROR cannot parse release object. %s\n", err)
			return
		}

		result, err := service.NewReleaseService(service.GitlabClientImpl).CreateRelease(projectID, newRelease)
		if err != nil {
			fmt.Println(err)
		} else {
			util.PrintPrettyJSON(result)
		}
	},
}
