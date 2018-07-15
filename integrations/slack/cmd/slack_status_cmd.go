package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	slackCmd.AddCommand(slackStatusCmd)
}

var slackStatusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Sends job status to slack",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("slack-status")
	},
}
