package cmd

import (
	"testing"
)

func TestDumpRevisionCmd(t *testing.T) {
	runCommandTestCases(t, []*cmdTestCase{
		{
			[]string{"dump-revision", "--help"},
			nil,
			nil,
			"dump_revision_help",
			false,
		},
		{
			[]string{"dump-rev", "--help"},
			nil,
			nil,
			"dump_revision_help",
			false,
		},
		{
			[]string{"dump-rev"},
			map[string]string{
				"CI_COMMIT_SHA": "e0e2f21877d4dd826bc21af3bc8bf0f0b273846d",
			},
			nil,
			"dump_revision",
			false,
		},
	})
}
