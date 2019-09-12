package service

import (
	"encoding/json"
	"fmt"

	"github.com/EXXETA/gitlab-cli/model"
)

// ProjectService represents the available functionalities to interace with Gitlab projects
type ProjectService struct {
	gitlabClient GitlabClient
}

// NewProjectService creates an instance of the project service
func NewProjectService(gc GitlabClient) *ProjectService {
	return &ProjectService{gitlabClient: gc}
}

// GetAllProjects gets all projects
func (projectService *ProjectService) GetAllProjects() (*[]model.Project, error) {
	resp, err := projectService.gitlabClient("GET", "/projects", nil)

	if err != nil {
		return nil, err
	}

	// close the body after process is done
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	projectList := &[]model.Project{}

	if err := dec.Decode(projectList); err != nil {
		return nil, fmt.Errorf("ERROR cannot parse project list. %s", err)
	}

	return projectList, nil
}

// GetProjectByID gets a project by id
func (projectService *ProjectService) GetProjectByID(projectID string) (*model.Project, error) {
	resp, err := projectService.gitlabClient("GET", "/projects/"+projectID, nil)

	if err != nil {
		return nil, err
	}

	// close the body after process is done
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	project := &model.Project{}

	if err := dec.Decode(project); err != nil {
		return nil, fmt.Errorf("ERROR cannot parse project object. %s", err)
	}

	return project, nil
}
