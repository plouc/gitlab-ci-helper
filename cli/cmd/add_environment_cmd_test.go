package cmd

import (
	"testing"
)

func TestAddEnvironmentCmd(t *testing.T) {
	ts := createMockServer(t, []*httpMock{
		{
			"POST",
			"/api/v4/projects/1/environments",
			"gitlab/environment.json",
		},
	})
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
		{
			[]string{"envs", "add"},
			nil,
			"environments_add_missing_arg",
			true,
		},
		{
			[]string{"envs", "add", "my env"},
			nil,
			"environments_add_missing_arg",
			true,
		},
		{
			[]string{"envs", "add", "my env", "http://fake.io"},
			map[string]string{
				"GITLAB_HOST":   ts.URL,
				"CI_PROJECT_ID": "1",
			},
			"environments_add",
			false,
		},
		{
			[]string{"envs", "add", "my env", "http://fake.io", "--verbose"},
			map[string]string{
				"GITLAB_HOST":   ts.URL,
				"CI_PROJECT_ID": "1",
			},
			"environments_add_verbose",
			false,
		},
	})
}
