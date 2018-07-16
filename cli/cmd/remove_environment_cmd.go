package cmd

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

func init() {
	environmentsCmd.AddCommand(removeEnvironmentsCmd)
}

var removeEnvironmentsCmd = &cobra.Command{
	Use:     "remove ENV_ID",
	Aliases: []string{"rm"},
	Short:   "Remove a project environment",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New(color.RedString("requires an ENV_ID"))
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		projectId := viper.GetString("project_id")

		envId, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New(fmt.Sprintf("unable to parse environment id %s:\n%v", args[0], err))
		}

		color.Yellow("removing env %d from project %s", envId, projectId)

		_, err = client.RemoveProjectEnvironment(projectId, envId)
		if err != nil {
			return err
		}

		color.Green("✔ successfully removed env")

		return nil
	},
}
