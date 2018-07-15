package cmd

import (
	"testing"
)

func TestDumpMetaCmd(t *testing.T) {
	runCommandTestCases(t, []*cmdTestCase{
		{
			[]string{"dump-meta", "--help"},
			nil,
			"dump_meta_help",
			false,
		},
	})
}
