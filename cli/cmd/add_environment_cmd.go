package cmd

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func init() {
	environmentsCmd.AddCommand(addEnvironmentsCmd)
}

var addEnvironmentsCmd = &cobra.Command{
	Use:   "add NAME URL",
	Short: "Add a new project environment",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New(color.RedString("requires a NAME and an URL"))
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		projectId := viper.GetString("project_id")
		if projectId == "" {
			return errors.New(color.RedString("✘ unable to determine project id"))
		}

		name := args[0]
		url := args[1]

		color.Yellow("creating environment %s", name)

		env := gitlab.EnvironmentAddPayload{
			Name:        name,
			ExternalUrl: url,
		}

		created, _, err := client.AddProjectEnvironment(projectId, &env)
		if err != nil {
			return errors.New(fmt.Sprintf("an error occurred while creating env %s:\n%v\n", name, err))
		}

		color.Green("✔ successfully created env %s (id: %d)", name, created.Id)

		if viper.GetBool("verbose") {
			created.RenderJson(os.Stdout)
			fmt.Println("")
		}

		return nil
	},
}
