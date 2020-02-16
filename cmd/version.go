package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/bskim45/dtags/common"
)

type versionOptions struct {
	short bool
}

func newVersionCmd() *cobra.Command {
	o := &versionOptions{}

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of dtags",
		Long:  `Show the version number of dtags.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.run(cmd.OutOrStdout())
		},
	}

	f := cmd.Flags()
	f.BoolVar(&o.short, "short", false, "print the version number")

	return cmd
}

func (o *versionOptions) run(out io.Writer) error {
	_, err := fmt.Fprintln(out, common.BuildCurrentVersionString(o.short))

	return err
}
