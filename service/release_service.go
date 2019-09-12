package service

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/EXXETA/gitlab-cli/model"
)

// ReleaseService represents the available functionalities to interace with Gitlab releases
type ReleaseService struct {
	gitlabClient GitlabClient
}

// NewReleaseService creates an instance of the release service
func NewReleaseService(gc GitlabClient) *ReleaseService {
	return &ReleaseService{gitlabClient: gc}
}

// GetAllReleasesByProjectID gets all releases of a project
func (releaseService *ReleaseService) GetAllReleasesByProjectID(projectID string) (*[]model.Release, error) {

	resp, err := releaseService.gitlabClient("GET", "/projects/"+projectID+"/releases", nil)

	if err != nil {
		return nil, err
	}

	// close the body after process is done
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	releaseList := &[]model.Release{}

	if err := dec.Decode(releaseList); err != nil {
		return nil, fmt.Errorf("ERROR cannot parse release list. %s", err)
	}

	return releaseList, nil
}

// GetReleaseByProjectIDAndTagName gets a release by project id and tag name
func (releaseService *ReleaseService) GetReleaseByProjectIDAndTagName(projectID string, tagName string) (*model.Release, error) {

	resp, err := releaseService.gitlabClient("GET", "/projects/"+projectID+"/releases/"+tagName, nil)

	if err != nil {
		return nil, err
	}

	// close the body after process is done
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	release := &model.Release{}

	if err := dec.Decode(release); err != nil {
		return nil, fmt.Errorf("ERROR cannot parse release object. %s", err)
	}

	return release, nil
}

// CreateRelease creates a new release for a project
func (releaseService *ReleaseService) CreateRelease(projectID string, release *model.Release) (*model.Release, error) {

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	if err := enc.Encode(release); err != nil {
		return nil, fmt.Errorf("ERROR cannot encode release object. %s", err)
	}

	resp, err := releaseService.gitlabClient("POST", "/projects/"+projectID+"/releases", &buf)

	if err != nil {
		return nil, err
	}

	// close the body after process is done
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	newRelease := &model.Release{}

	if err := dec.Decode(newRelease); err != nil {
		return nil, fmt.Errorf("ERROR cannot parse release object. %s", err)
	}

	return newRelease, nil
}
