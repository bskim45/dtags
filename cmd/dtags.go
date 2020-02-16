package cmd

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

func Execute(args []string) error {
	return newRootCmd(args).Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func newRootCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dtags",
		Short: "Retrieves Docker repositories and image tags",
		Long: `dtags is a small binary retrieves a list of Docker repositories and Docker Image tags
from various Docker registries.`,
		Args: cobra.NoArgs,
	}

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Path to config file (default $HOME/.dtagsrc)")
	cmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	_ = viper.BindPFlag("verbose", cmd.PersistentFlags().Lookup("verbose"))

	flags := cmd.PersistentFlags()
	_ = flags.Parse(args)

	cmd.AddCommand(
		newListTagsCmd(),
		newSearchCmd(),
		newVersionCmd(),
		newCompletionCmd(),
	)

	return cmd
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, _ := homedir.Dir()

		viper.AddConfigPath(home)
		viper.SetConfigName(".dtagsrc")
		viper.SetConfigType("toml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore
		} else {
			panic(fmt.Errorf("Fatal error on reading config file: %s \n", err))
		}
	}

	//fmt.Println("Using config file:", viper.ConfigFileUsed())
}
