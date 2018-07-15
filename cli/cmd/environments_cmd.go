package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(environmentsCmd)

	addProjectId(environmentsCmd)
}

var environmentsCmd = &cobra.Command{
	Use:     "environments",
	Aliases: []string{"envs"},
	Short:   "Project environments commands",
}
