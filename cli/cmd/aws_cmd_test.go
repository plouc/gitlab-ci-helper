package cmd

import (
	"testing"
)

func TestAwsCmd(t *testing.T) {
	ts := createMockServer(t, []*httpMock{})
	defer ts.Close()

	runCommandTestCases(t, []*cmdTestCase{
		{
			[]string{"aws", "--help"},
			nil,
			nil,
			"aws_help",
			false,
		},
	})
}
