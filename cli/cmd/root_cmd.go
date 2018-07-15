package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/plouc/go-gitlab-client/gitlab"
	aws "github.com/rande/gitlab-ci-helper/integrations/aws/cmd"
	hipchat "github.com/rande/gitlab-ci-helper/integrations/hipchat/cmd"
	slack "github.com/rande/gitlab-ci-helper/integrations/slack/cmd"
)

var configFile string
var client *gitlab.Gitlab

func init() {
	cobra.OnInitialize(initConfig)

	// global options
	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is .ci.yaml)")
	RootCmd.PersistentFlags().Bool("no-color", false, "disable color output")
	viper.BindPFlag("no-color", RootCmd.PersistentFlags().Lookup("no-color"))

	// GitLab specific options
	RootCmd.PersistentFlags().String( "host", "", "gitlab host\ndefault to env var: GITLAB_HOST\nor 'host' key in config file if exists")
	viper.BindPFlag("host", RootCmd.PersistentFlags().Lookup("host"))
	viper.BindEnv("host", "GITLAB_HOST")

	RootCmd.PersistentFlags().String( "token", "", "gitlab token\ndefault to env var: GITLAB_TOKEN\nor 'token' key in config file if exists")
	viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token"))
	viper.BindEnv("token", "GITLAB_TOKEN")

	RootCmd.PersistentFlags().String( "api-path", "", "gitlab api path\ndefault to env var: GITLAB_API_PATH\nor'api_path' key in config file if exists,\notherwise: '/api/v4'")
	viper.BindPFlag("api_path", RootCmd.PersistentFlags().Lookup("api-path"))
	viper.BindEnv("api_path", "GITLAB_API_PATH")
	viper.SetDefault("api_path", "/api/v4")

	// Third-party integrations
	registerIntegration(aws.Register())
	registerIntegration(hipchat.Register())
	registerIntegration(slack.Register())
}

func registerIntegration(name string, cmd *cobra.Command, config []string) {
	RootCmd.AddCommand(cmd)

	for _, key := range config {
		flagKey := fmt.Sprintf("%s-%s", name, key)
		envKey := fmt.Sprintf("%s_%s", strings.ToUpper(name), strings.ToUpper(key))
		keyPath := fmt.Sprintf("x.%s.%s", name, key)

		keyUsage := fmt.Sprintf(
			"%s %s\ndefault to environment variable: %s\nor '%s' path in config file if exists",
			name, key, envKey, keyPath,
		)

		cmd.PersistentFlags().String(flagKey, "", keyUsage)
		viper.BindPFlag(keyPath, cmd.PersistentFlags().Lookup(flagKey))
		viper.BindEnv(keyPath, envKey)
	}
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".ci")
	}

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
}

var RootCmd = &cobra.Command{
	Use:   "gitlab-ci-helper",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("no-color") {
			color.NoColor = true
		}

		client = gitlab.NewGitlab(
			viper.GetString("host"),
			viper.GetString("api_path"),
			viper.GetString("token"),
		)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
