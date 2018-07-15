package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	RootCmd.AddCommand(dumpMetaCmd)
}

var dumpMetaCmd = &cobra.Command{
	Use:     "dump-meta",
	Aliases: []string{"dm"},
	Short:   "Dump meta information about ci into a ci.json file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dump-meta")
	},
}
