package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	environmentsCmd.AddCommand(removeEnvironmentsCmd)
}

var removeEnvironmentsCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a project environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove-environment")
	},
}
