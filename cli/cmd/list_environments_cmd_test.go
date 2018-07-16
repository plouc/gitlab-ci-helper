package cmd

import (
	"testing"
)

func TestEnvironmentsListCmd(t *testing.T) {
	ts := createMockServer(t, []*httpMock{
		{
			"GET",
			"/api/v4/projects/1/environments",
			"gitlab/environments.json",
		},
		{
			"GET",
			"/api/v4/projects/2/environments",
			"gitlab/environments_empty.json",
		},
	})
	defer ts.Close()

	runCommandTestCases(t, []*cmdTestCase{
		{
			[]string{"environments", "list", "--help"},
			nil,
			nil,
			"environments_list_help",
			false,
		},
		{
			[]string{"envs", "ls", "--help"},
			nil,
			nil,
			"environments_list_help",
			false,
		},
		{
			[]string{"envs", "ls", "-p", "1"},
			map[string]string{
				"GITLAB_HOST": ts.URL,
			},
			nil,
			"environments_list",
			false,
		},
		{
			[]string{"envs", "ls"},
			map[string]string{
				"GITLAB_HOST":   ts.URL,
				"CI_PROJECT_ID": "1",
			},
			nil,
			"environments_list",
			false,
		},
		{
			[]string{"envs", "ls", "-p", "2"},
			map[string]string{
				"GITLAB_HOST": ts.URL,
			},
			nil,
			"environments_list_empty",
			false,
		},
	})
}
