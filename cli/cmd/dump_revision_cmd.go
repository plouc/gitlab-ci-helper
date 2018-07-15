package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(dumpRevisionCmd)
}

var dumpRevisionCmd = &cobra.Command{
	Use:     "dump-revision",
	Aliases: []string{"dump-rev"},
	Short:   "Dump a REVISION file with the current sha1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dump-revision")
	},
}
