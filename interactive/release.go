package interactive

import (
	"fmt"
	"time"

	"github.com/EXXETA/gitlab-cli/model"
	"github.com/EXXETA/gitlab-cli/service"
	"github.com/EXXETA/gitlab-cli/util"
)

var (
	releaseActions = []string{"get all releases of a project", "get a release by tag name", "create a release", "back"}
)

func releaseExecute() {
	// execute the action until user wants to got back
	for {
		if releaseAction() {
			break
		}
	}
}

// returns true to exit from release actions
func releaseAction() bool {
	index, _ := executeSelect("What would you like to do?", releaseActions)
	if index != -1 {
		var result, err interface{}
		releaseService := service.NewReleaseService(service.GitlabClientImpl)
		switch index {
		case 0:
			projectID := executeNotMaskedPrompt("Please enter project id", "")
			result, err = releaseService.GetAllReleasesByProjectID(projectID)
		case 1:
			projectID := executeNotMaskedPrompt("Please enter project id", "")
			tag := executeNotMaskedPrompt("Please enter release tag name", "")
			result, err = releaseService.GetReleaseByProjectIDAndTagName(projectID, tag)
		case 2:
			projectID := executeNotMaskedPrompt("Please enter project id", "")
			release := &model.Release{}
			release.Name = executeNotMaskedPrompt("Please enter the release name", "")
			release.TagName = executeNotMaskedPrompt("Please enter the release tag name", "")
			release.Description = executeNotMaskedPrompt("Please enter the release description", "")
			release.ReleasedAt = executeNotMaskedPrompt("Please enter the release time (optional)", time.Now().Format(time.RFC3339))
			release.Reference = executeNotMaskedPrompt("Please enter the release reference", "master")
			result, err = releaseService.CreateRelease(projectID, release)
		case 3:
			return true
		}

		if err != nil {
			fmt.Println(err)
		} else {
			util.PrintPrettyJSON(result)
		}
	}
	return false
}
