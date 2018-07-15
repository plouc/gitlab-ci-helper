package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

const integrationName = "hipchat"

func Register() (string, *cobra.Command, []string) {
	return integrationName, hipChatCmd, []string{
		"token",
		"host",
	}
}

var hipChatCmd = &cobra.Command{
	Use:     integrationName,
	Short:   fmt.Sprintf("%s commands", integrationName),
}
