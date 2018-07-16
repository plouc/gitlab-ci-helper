package cmd

import (
	"testing"
)

func TestSlackCmd(t *testing.T) {
	ts := createMockServer(t, []*httpMock{})
	defer ts.Close()

	runCommandTestCases(t, []*cmdTestCase{
		{
			[]string{"slack", "--help"},
			nil,
			nil,
			"slack_help",
			false,
		},
	})
}
