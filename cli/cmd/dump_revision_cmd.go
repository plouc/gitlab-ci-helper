package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func init() {
	RootCmd.AddCommand(dumpRevisionCmd)

	dumpRevisionCmd.Flags().StringP("file", "f", "REVISION", "The revision file\ncan also be defined with env var: REVISION_FILE")
	viper.BindPFlag("revision_file", dumpRevisionCmd.Flags().Lookup("file"))
	viper.BindEnv("revision_file", "REVISION_FILE")

	dumpRevisionCmd.Flags().StringP("ref", "r", "", "The sha1, default to env var: CI_COMMIT_SHA")
	viper.BindPFlag("revision_ref", dumpRevisionCmd.Flags().Lookup("ref"))
	viper.BindEnv("revision_ref", "CI_COMMIT_SHA")
}

var dumpRevisionCmd = &cobra.Command{
	Use:     "dump-revision",
	Aliases: []string{"dump-rev"},
	Short:   "Dump a REVISION file with the current sha1",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("creating revision file")

		file := viper.GetString("revision_file")
		ref := viper.GetString("revision_ref")

		fp, err := os.Create(file)
		defer fp.Close()
		if err != nil {
			return errors.New(fmt.Sprintf("unable to create revision file %s\n%v", file, err))
		}

		_, err = fp.Write([]byte(ref))
		if err != nil {
			return errors.New(fmt.Sprintf("unable to write to revision file %s\n%v", file, err))
		}

		fmt.Printf("ref: %s written to file: %s\n", ref, file)

		return nil
	},
}
