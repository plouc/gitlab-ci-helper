package cmd

import (
	"testing"
)

func TestDumpMetaCmd(t *testing.T) {
	runCommandTestCases(t, []*cmdTestCase{
		{
			[]string{"dump-meta", "--help"},
			nil,
			nil,
			"dump_meta_help",
			false,
		},
		{
			[]string{"dump-meta", "-v"},
			map[string]string{
				"CI_JOB_ID":          "CI_JOB_ID",
				"CI_COMMIT_SHA":      "CI_COMMIT_SHA",
				"CI_COMMIT_REF_NAME": "CI_COMMIT_REF_NAME",
				"CI_COMMIT_TAG":      "CI_COMMIT_TAG",
				"CI_JOB_STAGE":       "CI_JOB_STAGE",
				"CI_JOB_NAME":        "CI_JOB_NAME",
				"CI_JOB_URL":         "CI_JOB_URL",
				"CI_PROJECT_ID":      "CI_PROJECT_ID",
				"CI_PROJECT_DIR":     "CI_PROJECT_DIR",
				"CI_SERVER_NAME":     "CI_SERVER_NAME",
				"CI_SERVER_REVISION": "CI_SERVER_REVISION",
				"CI_SERVER_VERSION":  "CI_SERVER_VERSION",
			},
			nil,
			"dump_meta",
			false,
		},
	})
}
