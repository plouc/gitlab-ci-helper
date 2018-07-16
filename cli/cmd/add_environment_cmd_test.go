package cmd

import (
	"github.com/stretchr/testify/assert"
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
			nil,
			"environments_add_help",
			false,
		},
		{
			[]string{"envs", "add", "--help"},
			nil,
			nil,
			"environments_add_help",
			false,
		},
		{
			[]string{"envs", "add"},
			nil,
			nil,
			"environments_add_missing_arg",
			true,
		},
		{
			[]string{"envs", "add", "my env"},
			nil,
			nil,
			"environments_add_missing_arg",
			true,
		},
		{
			[]string{"envs", "add", "no project id", "http://fake.io"},
			map[string]string{
				"CI_PROJECT_ID": "",
			},
			nil,
			"environments_add_no_project_id",
			true,
		},
		{
			[]string{"envs", "add", "my env", "http://fake.io"},
			map[string]string{
				"GITLAB_HOST":   ts.URL,
				"CI_PROJECT_ID": "1",
			},
			func(t *testing.T, output string) {
				assert.Contains(t, output, "successfully created env my env")
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
			nil,
			"environments_add_verbose",
			false,
		},
	})
}
