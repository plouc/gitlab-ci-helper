package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	awsCmd.AddCommand(awsS3ArchiveCmd)
}

var awsS3ArchiveCmd = &cobra.Command{
	Use:     "s3-archive",
	Aliases: []string{"s3-arch"},
	Short:   "Send archive to a S3 bucket",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aws-s3-archive")
	},
}
