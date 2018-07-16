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
			nil,
			"environments_rm_help",
			false,
		},
		{
			[]string{"envs", "rm", "--help"},
			nil,
			nil,
			"environments_rm_help",
			false,
		},
		{
			[]string{"envs", "rm"},
			nil,
			nil,
			"environments_rm_missing_arg",
			true,
		},
		{
			[]string{"envs", "rm", "5"},
			map[string]string{
				"CI_PROJECT_ID": "",
			},
			nil,
			"environments_rm_no_project_id",
			true,
		},
		{
			[]string{"envs", "rm", "3"},
			map[string]string{
				"GITLAB_HOST":   ts.URL,
				"CI_PROJECT_ID": "1",
			},
			nil,
			"environments_rm",
			false,
		},
	})
}
