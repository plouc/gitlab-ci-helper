package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	environmentsCmd.AddCommand(listEnvironmentsCmd)
}

var listEnvironmentsCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List project's environments",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list-environments")

		fmt.Println("project_id", viper.Get("project_id"))
	},
}
