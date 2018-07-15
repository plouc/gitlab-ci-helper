package cmd

import (
	"fmt"

	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
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

		environments, _, err := client.ProjectEnvironments(viper.GetString("project_id"), &gitlab.PaginationOptions{
			PerPage: 100,
		})
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		for _, e := range environments.Items {
			fmt.Printf(" > %4d - %-16s - %s\n", e.Id, e.Name, e.ExternalUrl)
		}
	},
}
