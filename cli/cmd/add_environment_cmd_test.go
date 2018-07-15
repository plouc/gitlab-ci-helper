package cmd

import (
	"testing"
)

func TestAddEnvironmentCmd(t *testing.T) {
	ts := createMockServer(t, []*httpMock{})
	defer ts.Close()

	runCommandTestCases(t, []*cmdTestCase{
		{
			[]string{"environments", "add", "--help"},
			nil,
			"environments_add_help",
			false,
		},
		{
			[]string{"envs", "add", "--help"},
			nil,
			"environments_add_help",
			false,
		},
	})
}
