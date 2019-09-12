package service_test

import (
	"encoding/json"
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
	expectedVersion model.Version = model.Version{
		Version:  "8.13.0-pre",
		Revision: "4e963fe",
	}
)

func TestGetGitlabVersion_shouldReturnVersion(t *testing.T) {
	versionJSON, err := json.Marshal(expectedVersion)
	result, err := service.NewVersionService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(strings.NewReader(string(versionJSON)))}, nil
	}).GetGitlabVersion()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedVersion.Version, result.Version)
}

func TestGetGitlabVersion_shouldReturnError(t *testing.T) {
	result, err := service.NewVersionService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(strings.NewReader("wrong_json_format"))}, nil
	}).GetGitlabVersion()

	assert.NotNil(t, err)
	assert.Nil(t, result)
}
