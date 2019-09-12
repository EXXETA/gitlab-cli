package interactive

import (
	"fmt"

	"github.com/EXXETA/gitlab-cli/service"
	"github.com/EXXETA/gitlab-cli/util"
)

var (
	projectActions = []string{"get all projects", "get a project by id", "back"}
)

func projectExecute() {
	// execute the action until user wants to got back
	for {
		if projectAction() {
			break
		}
	}
}

// returns true to exit from project actions
func projectAction() bool {

	index, _ := executeSelect("What would you like to do?", projectActions)
	if index != -1 {
		var result, err interface{}
		projectService := service.NewProjectService(service.GitlabClientImpl)
		switch index {
		case 0:
			result, err = projectService.GetAllProjects()
		case 1:
			projectID := executeNotMaskedPrompt("Please enter the project id", "")
			result, err = projectService.GetProjectByID(projectID)
		case 2:
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
