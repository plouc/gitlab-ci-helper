package cmd

import (
	"testing"
)

func TestHipchatCmd(t *testing.T) {
	ts := createMockServer(t, []*httpMock{})
	defer ts.Close()

	runCommandTestCases(t, []*cmdTestCase{
		{
			[]string{"hipchat", "--help"},
			nil,
			"hipchat_help",
			false,
		},
	})
}
