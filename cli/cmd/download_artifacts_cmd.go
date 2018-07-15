package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(downloadArtifactsCmd)
}

var downloadArtifactsCmd = &cobra.Command{
	Use:     "download-artifacts",
	Aliases: []string{"dl-artifacts", "dl"},
	Short:   "Download job artifacts and extract them to specified path if provided",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download-artifacts")
	},
}
