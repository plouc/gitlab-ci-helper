package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	awsCmd.AddCommand(awsS3ExtractCmd)
}

var awsS3ExtractCmd = &cobra.Command{
	Use:     "s3-extract",
	Aliases: []string{"s3-ext"},
	Short:   "Extract archive from a S3 bucket",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aws-s3-extract")
	},
}
