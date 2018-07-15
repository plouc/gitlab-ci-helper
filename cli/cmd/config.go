package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addProjectId(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("project-id", "p", "", "project id\ndefault to environment variable: CI_PROJECT_ID")
	viper.BindPFlag("project_id", cmd.PersistentFlags().Lookup("project-id"))
	viper.BindEnv("project_id", "CI_PROJECT_ID")
}
