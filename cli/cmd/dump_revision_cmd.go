package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	RootCmd.AddCommand(dumpRevisionCmd)
}

var dumpRevisionCmd = &cobra.Command{
	Use:     "dump-rev",
	Aliases: []string{"dr"},
	Short:   "Dump a REVISION file with the current sha1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dump-revision")
	},
}
