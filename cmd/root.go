package cmd

import (
	"fmt"

	"github.com/EXXETA/gitlab-cli/interactive"
	"github.com/EXXETA/gitlab-cli/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&service.GitlabBaseURL, "url", "u", "", "GitLab server base url")
	rootCmd.PersistentFlags().StringVarP(&service.PrivateToken, "token", "t", "", "Personal access token to authenticate against GitLab")
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
