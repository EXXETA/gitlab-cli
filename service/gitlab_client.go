package service

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// GitlabClient sends http requests to Gitlab api
type GitlabClient func(method string, path string, body io.Reader) (*http.Response, error)

var (
	// GitlabBaseURL base url of the Gitlab server
	GitlabBaseURL string

	gitlabAPIPath = "/api/v4"

	// PrivateToken private token for authenticating the user
	PrivateToken string
)

// GitlabClientImpl an implementation of GitlabClient to send http requests to Gitlab api
func GitlabClientImpl(method string, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, GitlabBaseURL+gitlabAPIPath+path, body)
	if err != nil {
		return nil, fmt.Errorf("ERROR cannot connect to Gitlab api. %s", err)
	}

	// set the required headers
	req.Header.Set("PRIVATE-TOKEN", PrivateToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ERROR cannot connect to Gitlab api. %s", err)
	}

	// check for error
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("ERROR cannot read response body. %s", err)
		}
		return nil, fmt.Errorf("ERROR Gitlab returns error. %s", respBody)
	}

	return resp, nil
}
