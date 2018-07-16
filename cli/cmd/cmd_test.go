package cmd

import (
	"fmt"
	"github.com/plouc/gosnap"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

const binaryName = "gitlab-ci-helper"

var snapshotsDir string

type cmdTestCase struct {
	args     []string
	env      map[string]string
	manual   func(t *testing.T, output string)
	snapshot string
	wantErr  bool
}

type httpMock struct {
	method   string
	path     string
	bodyFile string
}

func runCommandTestCase(t *testing.T, ctx *gosnap.Context, tc *cmdTestCase) {
	t.Run(strings.Join(tc.args, "_"), func(t *testing.T) {
		var s *gosnap.Snapshot
		if !ctx.Has(tc.snapshot) {
			s = ctx.NewSnapshot(tc.snapshot)
		} else {
			s = ctx.Get(tc.snapshot)
		}

		c := exec.Command(fmt.Sprintf("./%s", binaryName), tc.args...)
		env := os.Environ()
		for k, v := range tc.env {
			env = append(env, fmt.Sprintf("%s=%s", k, v))
		}
		c.Env = env

		output, err := c.CombinedOutput()
		if tc.wantErr {
			assert.Error(t, err, "command '%s' was expected to throw an error", strings.Join(tc.args, " "))
		} else {
			assert.NoError(t, err, "command '%s' wasn't expected to throw an error", strings.Join(tc.args, " "))
		}

		if !t.Failed() && tc.manual != nil {
			tc.manual(t, string(output))
		}

		if !t.Failed() {
			s.AssertString(string(output))
		}
	})
}

func runCommandTestCases(t *testing.T, testCases []*cmdTestCase) {
	ctx := gosnap.NewContext(t, snapshotsDir)
	for _, testCase := range testCases {
		runCommandTestCase(t, ctx, testCase)
	}
}

func createMockServer(t *testing.T, mocks []*httpMock) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, m := range mocks {
			if m.method == r.Method && m.path == r.URL.Path {
				f := filepath.Join("..", "mocks", m.bodyFile)
				content, err := ioutil.ReadFile(f)
				if err != nil {
					t.Errorf("unable to read mock file: %s\n%v", f, err)
					return
				}

				w.Write(content)

				return
			}
		}

		t.Errorf("no request found matching: %s %s\nevery request has to be mocked!", r.Method, r.URL.Path)
	}))

	return ts
}

func TestMain(m *testing.M) {
	err := os.Chdir("..")
	if err != nil {
		fmt.Printf("could not change dir: %v\n", err)
		os.Exit(1)
	}

	cmd := exec.Command("go", "build", "-o", binaryName, "main.go")
	cmd.Start()
	if err := cmd.Wait(); err != nil {
		output, _ := cmd.CombinedOutput()
		fmt.Printf("could not make '%s' binary:\n%v\n%s", binaryName, err, output)
		os.Exit(1)
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("problems recovering caller information")
		os.Exit(1)
	}
	snapshotsDir = filepath.Join(filepath.Dir(filename), "snapshots")

	os.Exit(m.Run())
}
