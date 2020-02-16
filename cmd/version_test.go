package cmd

import (
	"testing"

	"github.com/bskim45/dtags/common"
)

func TestVersion(t *testing.T) {
	tests := []CmdTestCase{{
		name:     "default",
		cmd:      "version",
		expected: common.BuildCurrentVersionString(false),
	}, {
		name:     "short",
		cmd:      "version --short",
		expected: common.BuildCurrentVersionString(true),
	}}
	runTestCmds(t, tests)
}
