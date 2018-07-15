package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const integrationName = "slack"

func Register() (string, *cobra.Command, []string) {
	return integrationName, slackCmd, []string{
		"token",
	}
}

var slackCmd = &cobra.Command{
	Use:   integrationName,
	Short: fmt.Sprintf("%s commands", integrationName),
}
