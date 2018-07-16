package cmd

import (
	"testing"
)

func TestEnvironmentsCmd(t *testing.T) {
	runCommandTestCases(t, []*cmdTestCase{
		{
			[]string{"environments", "--help"},
			nil,
			nil,
			"environments_help",
			false,
		},
		{
			[]string{"envs", "--help"},
			map[string]string{
				"WHATEVER": "whatever",
			},
			nil,
			"environments_help",
			false,
		},
	})
}
