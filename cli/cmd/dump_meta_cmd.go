package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"bytes"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
)

type JobMeta struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Ref     string `json:"ref"`
	RefName string `json:"ref_name"`
	Tag     string `json:"tag"`
	Stage   string `json:"stage"`
	Url     string `json:"url"`
}

type ProjectMeta struct {
	Id  string `json:"id"`
	Dir string `json:"dir"`
}

type ServerMeta struct {
	Name     string `json:"name"`
	Revision string `json:"revision"`
	Version  string `json:"version"`
}

type Meta struct {
	Job     *JobMeta     `json:"job"`
	Project *ProjectMeta `json:"project"`
	Server  *ServerMeta  `json:"server"`
}

func init() {
	RootCmd.AddCommand(dumpMetaCmd)

	dumpMetaCmd.Flags().StringP("file", "f", "ci.json", "The meta file\ncan also be defined with env var: META_FILE")
	viper.BindPFlag("meta_file", dumpMetaCmd.Flags().Lookup("file"))
	viper.BindEnv("meta_file", "META_FILE")
}

var dumpMetaCmd = &cobra.Command{
	Use:   "dump-meta",
	Short: "Dump meta information about CI in a file",
	RunE: func(cmd *cobra.Command, args []string) error {
		file := viper.GetString("meta_file")

		color.Yellow("creating meta file %s\n", file)

		meta := &Meta{
			Job: &JobMeta{
				Id:      os.Getenv("CI_JOB_ID"),
				Name:    os.Getenv("CI_JOB_NAME"),
				Ref:     os.Getenv("CI_COMMIT_SHA"),
				RefName: os.Getenv("CI_COMMIT_REF_NAME"),
				Tag:     os.Getenv("CI_COMMIT_TAG"),
				Stage:   os.Getenv("CI_JOB_STAGE"),
				Url:     os.Getenv("CI_JOB_URL"),
			},
			Project: &ProjectMeta{
				Id:  os.Getenv("CI_PROJECT_ID"),
				Dir: os.Getenv("CI_PROJECT_DIR"),
			},
			Server: &ServerMeta{
				Name:     os.Getenv("CI_SERVER_NAME"),
				Revision: os.Getenv("CI_SERVER_REVISION"),
				Version:  os.Getenv("CI_SERVER_VERSION"),
			},
		}

		metaJson, err := json.Marshal(meta)
		if err != nil {
			return errors.New(fmt.Sprintf("an error occurred while marshalling meta\n%v", err))
		}

		var indented bytes.Buffer
		json.Indent(&indented, metaJson, "", "    ")

		if viper.GetBool("verbose") {
			os.Stdout.Write(indented.Bytes())
			fmt.Println("")
		}

		err = ioutil.WriteFile(file, indented.Bytes(), 0644)
		if err != nil {
			return errors.New(fmt.Sprintf("unable to create meta file %s\n%v", file, err))
		}

		color.Green("✔ successfully wrote meta to: %s\n", file)

		return nil
	},
}
