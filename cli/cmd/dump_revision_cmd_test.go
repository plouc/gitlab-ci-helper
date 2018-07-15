package cmd

import (
	"testing"
)

func TestDumpRevisionCmd(t *testing.T) {
	runCommandTestCases(t, []*cmdTestCase{
		{
			[]string{"dump-revision", "--help"},
			nil,
			"dump_revision_help",
			false,
		},
		{
			[]string{"dump-rev", "--help"},
			nil,
			"dump_revision_help",
			false,
		},
	})
}
