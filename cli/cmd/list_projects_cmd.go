package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	RootCmd.AddCommand(listProjectsCmd)
}

var listProjectsCmd = &cobra.Command{
	Use:     "list-projects",
	Aliases: []string{"projects"},
	Short:   "List available projects",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list-projects")
	},
}
