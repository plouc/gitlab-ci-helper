package cmd

import (
	"testing"
)

func TestEnvironmentsCmd(t *testing.T) {
	runCommandTestCases(t, []*cmdTestCase{
		{
			[]string{"environments", "--help"},
			nil,
			"environments_help",
			false,
		},
		{
			[]string{"envs", "--help"},
			map[string]string{
				"WHATEVER": "whatever",
			},
			"environments_help",
			false,
		},
	})
}
