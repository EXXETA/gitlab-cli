package cmd

import (
	"fmt"

	"github.com/EXXETA/gitlab-cli/config"
	"github.com/EXXETA/gitlab-cli/interactive"
	"github.com/EXXETA/gitlab-cli/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&service.GitlabBaseURL, "url", "u", "", "GitLab server base url")
	rootCmd.PersistentFlags().StringVarP(&service.PrivateToken, "token", "t", "", "Personal access token to authenticate against GitLab")

	// initialize the config file and read server settings
	config.InitConfig("")
	service.GitlabBaseURL = config.GetConfigValue("baseUrl")
	service.PrivateToken = config.GetConfigValue("privateToken")
}

// rootCmd defines the root command of the gitlabacli command mode
var rootCmd = &cobra.Command{
	Use:   "gitlabcli",
	Short: "A command line tool to use GitLab api",
	Long:  "gitlabcli provides a command-line tool to interact with GitLab servers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Long)
		// start the interactive mode
		interactive.Execute()
	},
}

// Execute initializes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
