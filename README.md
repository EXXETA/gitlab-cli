# gitlab-cli
A simple command-line tool to interact with [GitLab API](https://docs.gitlab.com/ee/api/). Currently only the version 4 of GitLab API is supported.

## Overview

gitlab-cli provides a simple interface to interact with Gitlab servers. It has two main modes:

- `Interactive` provides a list of actions to choose from in an interactive way.

- `Command` provides a single command line for user actions, arguments and flags.

### Supported actions

- projects
    - get list of projects
    - get a project with id
- releases
    - get list of releases of a project
    - get a release with project id and tag name
    - create a new release for a project

### Authentication

Gitlab authentication is done using [Person Access Token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html). To generate a new token, go to "Gitlab > User Settings > Access Token".

## Interactive mode

In the interactive mode, gitlab-cli asks user the actions and parameters for each action in an interactive way.

## Command mode

In the command mode, user should enter the action and all the necessary arguments and flags together to perform the desired action. Example:
- `gitlab-cli project list --url <gitlab_base_url> --token <your_access_token>` to get list of projects
- `gitlab-cli release get --project-id <project_id> --tag-name <tag_name> --url <gitlab_base_url> --token <your_access_token>` to get a release with project id and tag name
- `gitlab-cli release create '{"name":"release 0.0.1","tag_name":"0.0.1","description":"A nice description","ref":"master","released_at":"2019-09-12T11:04:05+02:00"}' --url <gitlab_base_url> --token <your_access_token> --project-id <project_id>` to create a new release

## TODO

- Other actions can be added. Feel free to create pull requests or issues.
- Save the gitlab server info (base url and access token) in a config file.

## Dev it!

### To build:

 - `go build -o gitlab-cli`

 Optionally you can set the version number:
 - `go build -o gitlab-cli -ldflags "-X https://github.com/EXXETA/gitlab-cli/cmd.Version=x.y.z"`

### To test:
- `go test -v ./...`

### To run:
- `go run gitlab-cli`

or 

- `./gitlab-cli`

## Dependencies

- [promptui](https://github.com/manifoldco/promptui)
- [cobra](https://github.com/spf13/cobra)
- [Testify](https://github.com/stretchr/testify)
