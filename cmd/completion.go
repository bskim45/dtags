package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

func newCompletionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion [SHELL]",
		Short: "Generate an autocompletion script for the specified shell (bash or zsh)",
		Long: `Generate an autocompletion script for the specified shell (bash or zsh)

Generate autocompletion:
    $ dtags completion bash

Sourcing into the shell:
    $ source <(dtags completion bash)
`,
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: []string{"bash", "zsh"},
		Run: func(cmd *cobra.Command, args []string) {
			arg := strings.Join(args, " ")

			switch arg {
			case "bash":
				_ = cmd.Root().GenBashCompletion(cmd.OutOrStdout())
			case "zsh":
				_ = cmd.Root().GenZshCompletion(cmd.OutOrStdout())
			}
		},
	}

	return cmd
}
