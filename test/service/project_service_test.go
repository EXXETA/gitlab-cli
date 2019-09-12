package service_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/EXXETA/gitlab-cli/model"
	"github.com/EXXETA/gitlab-cli/service"
)

var (
	expectedProjects []model.Project = []model.Project{
		model.Project{
			ID:   1,
			Name: "Project 1",
		},
		model.Project{
			ID:   2,
			Name: "Project 2",
		},
	}
)

func TestGetAllProjects_shouldReturnProjects(t *testing.T) {
	projectsJSON, err := json.Marshal(expectedProjects)
	result, err := service.NewProjectService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(strings.NewReader(string(projectsJSON)))}, nil
	}).GetAllProjects()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(expectedProjects), len(*result))
}

func TestGetAllProjects_shouldReturnError(t *testing.T) {
	result, err := service.NewProjectService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(strings.NewReader("wrong_json_format"))}, nil
	}).GetAllProjects()

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestGetProjectByID_shouldReturnProject(t *testing.T) {
	projectJSON, err := json.Marshal(expectedProjects[0])
	result, err := service.NewProjectService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(strings.NewReader(string(projectJSON)))}, nil
	}).GetProjectByID("1")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedProjects[0].ID, result.ID)
}

func TestGetProjectByID_shouldReturnNotFound(t *testing.T) {
	result, err := service.NewProjectService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{}, fmt.Errorf("{\"message\": \"404 Project Not Found\"}")
	}).GetProjectByID("1")

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestGetProjectByID_shouldReturnError(t *testing.T) {
	result, err := service.NewProjectService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(strings.NewReader("wrong_json_format"))}, nil
	}).GetProjectByID("1")

	assert.NotNil(t, err)
	assert.Nil(t, result)
}
