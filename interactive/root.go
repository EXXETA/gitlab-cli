package interactive

import (
	"fmt"
	"os"
	"strings"

	"github.com/EXXETA/gitlab-cli/service"
	"github.com/manifoldco/promptui"
)

var (
	rootActions         = []string{"projects", "releases", "exit"}
	exxetaGitlabBaseURL = "https://gitlabci.exxeta.com"
)

// Execute main method of the interactive mode
func Execute() {

	// gitlab base url
	if service.GitlabBaseURL == "" {
		gitlabURL := executeNotMaskedPrompt("Please enter your GitLab base url", "")
		if gitlabURL != "" {
			// remove the ending "/" from the URL if it has any
			service.GitlabBaseURL = strings.TrimSuffix(gitlabURL, "/")
		}
	}

	// person access token
	if service.PrivateToken == "" {
		personalToken := executeMaskedPrompt("Please enter your GitLab personal access token (You can find it under 'GitLab > User Settings > Access Tokens')", "")
		if personalToken != "" {
			service.PrivateToken = personalToken
		}
	}

	// access the "/version" endpoint to make sure the url and the token are valid
	fmt.Println("Checking server availability and access token...")
	_, error := service.NewVersionService(service.GitlabClientImpl).GetGitlabVersion()
	if error != nil {
		fmt.Println(error)
		return
	}

	// execute the action until user wants to exit
	for {
		if rootAction() {
			break
		}
	}
}

// returns true to exit from root actions
func rootAction() bool {
	index, _ := executeSelect("What would you like to do?", rootActions)
	if index != -1 {
		switch index {
		case 0:
			projectExecute()
		case 1:
			releaseExecute()
		case 2:
			return true
		}
	}
	return false
}

// executeMaskedPrompt wraps executePrompt to mask the input value
func executeMaskedPrompt(label string, defaultValue string) string {
	return executePrompt(label, defaultValue, true)
}

// executeNotMaskedPrompt wraps executePrompt to not mask the input value
func executeNotMaskedPrompt(label string, defaultValue string) string {
	return executePrompt(label, defaultValue, false)
}

func executePrompt(label string, defaultValue string, maskInput bool) string {
	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
	}

	if maskInput {
		prompt.Mask = '*'
	}

	result, err := prompt.Run()

	// handle "Control + C / D"
	if err == promptui.ErrInterrupt || err == promptui.ErrEOF {
		os.Exit(-1)
	}

	if err != nil {
		fmt.Println("ERROR prompt failed", err)
		return ""
	}

	return result
}

func executeSelect(label string, items []string) (int, string) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	index, result, err := prompt.Run()

	// handle "Control + C / D"
	if err == promptui.ErrInterrupt || err == promptui.ErrEOF {
		os.Exit(-1)
	}

	if err != nil {
		fmt.Println("ERROR prompt failed", err)
		return -1, ""
	}

	return index, result
}
