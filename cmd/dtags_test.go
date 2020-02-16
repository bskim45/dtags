package cmd

import (
	"bytes"
	"testing"

	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
)

type CmdTestCase struct {
	name     string
	cmd      string
	expected string
}

func runTestCmds(t *testing.T, tests []CmdTestCase) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("running cmd: ", tt.cmd)
			_, out, _ := executeActionCommandC(tt.cmd)
			if tt.expected != "" {
				expected := tt.expected + "\n"
				if expected != out {
					t.Fatalf("WANT:\n'%s'\n\nGOT:\n'%s'\n", expected, out)
				}
			}
		})
	}
}

func executeActionCommandC(cmd string) (*cobra.Command, string, error) {
	args, err := shellwords.Parse(cmd)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)

	root := newRootCmd(args)
	root.SetOut(buf)
	root.SetArgs(args)

	c, err := root.ExecuteC()

	return c, buf.String(), err
}
