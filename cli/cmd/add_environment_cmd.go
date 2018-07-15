package cmd

import (
	"fmt"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	environmentsCmd.AddCommand(addEnvironmentsCmd)
}

var addEnvironmentsCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new project environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add-environment")

		e := gitlab.EnvironmentAddPayload{
			Name: "test",
		}

		created, _, err := client.AddProjectEnvironment(viper.GetString("project_id"), &e)
		if err != nil {
			fmt.Printf("an error occurred while creating env:\n%v\n", err)
			return
		}

		fmt.Printf("Successfully created env (id: %d)", created.Id)
	},
}
