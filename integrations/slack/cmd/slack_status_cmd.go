package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	slackCmd.AddCommand(slackStatusCmd)
}

var slackStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Sends job status to slack",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("slack-status")
	},
}
