package cmd

import (
	"testing"
)

func TestRemoveEnvironmentCmd(t *testing.T) {
	ts := createMockServer(t, []*httpMock{
		{
			"DELETE",
			"/api/v4/projects/1/environments/3",
			"gitlab/delete.txt",
		},
	})
	defer ts.Close()

	runCommandTestCases(t, []*cmdTestCase{
		{
			[]string{"environments", "remove", "--help"},
			nil,
			"environments_rm_help",
			false,
		},
		{
			[]string{"envs", "rm", "--help"},
			nil,
			"environments_rm_help",
			false,
		},
		{
			[]string{"envs", "rm"},
			nil,
			"environments_rm_missing_arg",
			true,
		},
		{
			[]string{"envs", "rm", "3"},
			map[string]string{
				"GITLAB_HOST":   ts.URL,
				"CI_PROJECT_ID": "1",
			},
			"environments_rm",
			false,
		},
	})
}
