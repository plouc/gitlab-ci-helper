package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const integrationName = "hipchat"

func Register() (string, *cobra.Command, []string) {
	return integrationName, hipChatCmd, []string{
		"token",
		"host",
	}
}

var hipChatCmd = &cobra.Command{
	Use:   integrationName,
	Short: fmt.Sprintf("%s commands", integrationName),
}
