package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	hipChatCmd.AddCommand(hipChatStatusCmd)
}

var hipChatStatusCmd = &cobra.Command{
	Use:   "status ROOM MESSAGE",
	Short: "Sends job status to hipchat",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hipchat-status")
	},
}
