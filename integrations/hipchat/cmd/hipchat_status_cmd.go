package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	hipChatCmd.AddCommand(hipChatStatusCmd)
}

var hipChatStatusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Sends job status to hipchat",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hipchat-status")
	},
}
