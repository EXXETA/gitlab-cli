package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/EXXETA/gitlab-cli/service"
)

func TestGitlabAPIRequest_wrongGitlabURLFormat_shouldReturnError(t *testing.T) {
	service.GitlabBaseURL = "wrong_url_format"

	_, err := service.GitlabClientImpl("GET", "something", nil)

	assert.NotNil(t, err)
}

func TestGitlabAPIRequest_wrongGitlabURL_shouldReturnError(t *testing.T) {
	service.GitlabBaseURL = "https://wrong.gitlab.base.url"

	_, err := service.GitlabClientImpl("GET", "something", nil)

	assert.NotNil(t, err)
}
