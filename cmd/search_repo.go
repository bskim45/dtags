package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/bskim45/dtags/common"
	"github.com/bskim45/dtags/common/api_client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type searchHubOptions struct {
	searchEndpoint string
	fetchSize      uint
}

func newSearchCmd() *cobra.Command {
	o := &searchHubOptions{}

	cmd := &cobra.Command{
		Use:   "search [keyword]",
		Short: "Search for docker images in the docker registry",
		Long:  "",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.run(cmd.OutOrStdout(), args)
		},
	}

	f := cmd.Flags()
	f.StringVar(&o.searchEndpoint, "endpoint", "https://registry.hub.docker.com/", "monocular instance to query for charts")
	f.UintVarP(&o.fetchSize, "fetch-size", "n", 50, "maximum result size for output table")

	return cmd
}

func (o *searchHubOptions) run(out io.Writer, args []string) error {
	url, err := common.BuildUrl(o.searchEndpoint)

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid endpoint: %q", o.searchEndpoint))
	}

	c, err := api_client.New(url.String())

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to create connection to %q", o.searchEndpoint))
	}

	q := strings.Join(args, " ")

	results, err := c.SearchRepo(q, o.fetchSize)
	if err != nil {
		return fmt.Errorf("unable to perform search against %q", o.searchEndpoint)
	}

	var e error
	for i, t := range *results {
		if uint(i) > o.fetchSize {
			break
		}
		_, e = fmt.Fprintln(out, t)
	}

	return e
}
