package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const integrationName = "aws"

func Register() (string, *cobra.Command, []string) {
	return integrationName, awsCmd, []string{
		"region",
		"endpoint",
		"profile",
	}
}

var awsCmd = &cobra.Command{
	Use:   integrationName,
	Short: fmt.Sprintf("%s commands", integrationName),
}
