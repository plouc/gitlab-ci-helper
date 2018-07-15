package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	environmentsCmd.AddCommand(addEnvironmentsCmd)
}

var addEnvironmentsCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new project environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add-environment")
	},
}
