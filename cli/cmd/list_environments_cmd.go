package cmd

import (
	"fmt"

	"errors"
	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		projectId := viper.GetString("project_id")
		if projectId == "" {
			return errors.New(color.RedString("✘ unable to determine project id"))
		}

		environments, _, err := client.ProjectEnvironments(projectId, &gitlab.PaginationOptions{
			PerPage: 100,
		})
		if err != nil {
			return err
		}

		color.Yellow("available environments")
		for _, e := range environments.Items {
			fmt.Printf(" > %4d - %-16s - %s\n", e.Id, e.Name, e.ExternalUrl)
		}

		return nil
	},
}
