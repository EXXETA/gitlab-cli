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
	expectedReleases []model.Release = []model.Release{
		model.Release{
			Name:    "Release 1",
			TagName: "0.0.1",
		},
		model.Release{
			Name:    "Release 2",
			TagName: "0.0.2",
		},
	}
)

func TestGetAllReleases_shouldReturnReleases(t *testing.T) {
	releasesJSON, err := json.Marshal(expectedReleases)
	result, err := service.NewReleaseService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(strings.NewReader(string(releasesJSON)))}, nil
	}).GetAllReleasesByProjectID("1")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(expectedReleases), len(*result))
}

func TestGetAllReleases_shouldReturnNotFound(t *testing.T) {
	result, err := service.NewReleaseService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{}, fmt.Errorf("{\"message\": \"404\"}")
	}).GetAllReleasesByProjectID("1")

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestGetAllReleases_shouldReturnError(t *testing.T) {
	result, err := service.NewReleaseService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(strings.NewReader("wrong_json_format"))}, nil
	}).GetAllReleasesByProjectID("1")

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestGetReleaseByTagName_shouldReturnRelease(t *testing.T) {
	releaseJSON, err := json.Marshal(expectedReleases[0])
	result, err := service.NewReleaseService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(strings.NewReader(string(releaseJSON)))}, nil
	}).GetReleaseByProjectIDAndTagName("1", "0.0.1")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedReleases[0].TagName, result.TagName)
}

func TestGetReleaseByTagName_shouldReturnNotFound(t *testing.T) {
	result, err := service.NewReleaseService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{}, fmt.Errorf("{\"message\": \"404 Project Not Found\"}")
	}).GetReleaseByProjectIDAndTagName("1", "0.0.1")

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestGetReleaseByTagName_shouldReturnError(t *testing.T) {
	result, err := service.NewReleaseService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(strings.NewReader("wrong_json_format"))}, nil
	}).GetReleaseByProjectIDAndTagName("1", "0.0.1")

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestCreateRelease_shouldReturnRelease(t *testing.T) {
	releaseJSON, err := json.Marshal(expectedReleases[0])
	result, err := service.NewReleaseService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(strings.NewReader(string(releaseJSON)))}, nil
	}).CreateRelease("1", &expectedReleases[0])

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedReleases[0].TagName, result.TagName)
}

func TestCreateRelease_shouldReturnNotFound(t *testing.T) {
	result, err := service.NewReleaseService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{}, fmt.Errorf("{\"message\": \"404 Project Not Found\"}")
	}).CreateRelease("1", &expectedReleases[0])

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestCreateRelease_shouldReturnError(t *testing.T) {
	result, err := service.NewReleaseService(func(method string, path string, body io.Reader) (*http.Response, error) {
		return &http.Response{Body: ioutil.NopCloser(strings.NewReader("wrong_json_format"))}, nil
	}).CreateRelease("1", &expectedReleases[0])

	assert.NotNil(t, err)
	assert.Nil(t, result)
}
