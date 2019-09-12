package service

import (
	"encoding/json"
	"fmt"

	"github.com/EXXETA/gitlab-cli/model"
)

// VersionService represents the available functionalities to interace with Gitlab version
type VersionService struct {
	gitlabClient GitlabClient
}

// NewVersionService creates an instance of the version service
func NewVersionService(gc GitlabClient) *VersionService {
	return &VersionService{gitlabClient: gc}
}

// GetGitlabVersion gets version of the Gitlab
func (versionService *VersionService) GetGitlabVersion() (*model.Version, error) {
	resp, err := versionService.gitlabClient("GET", "/version", nil)

	if err != nil {
		return nil, err
	}

	// close the body after process is done
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	version := &model.Version{}

	if err := dec.Decode(version); err != nil {
		return nil, fmt.Errorf("ERROR cannot parse version object %s", err)
	}

	return version, nil
}
