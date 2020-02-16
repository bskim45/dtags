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

type listTagsHubOptions struct {
	searchEndpoint string
	fetchSize      uint
}

func newListTagsCmd() *cobra.Command {
	o := &listTagsHubOptions{}

	cmd := &cobra.Command{
		Use:   "tags [image]",
		Short: "Get the list of tags of the docker image in the docker registry",
		Long:  "",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.run(cmd.OutOrStdout(), args)
		},
	}

	f := cmd.Flags()
	f.StringVar(&o.searchEndpoint, "endpoint", "https://registry.hub.docker.com/", "docker registry instance to query for image")
	f.UintVarP(&o.fetchSize, "fetch-size", "n", 100, "maximum result size for output table")

	return cmd
}

func (o *listTagsHubOptions) run(out io.Writer, args []string) error {
	q := common.ParseName(strings.Join(args, " "))

	endpoint := o.searchEndpoint

	if q.Endpoint != "" {
		hostUrl, err := common.BuildUrl(q.Endpoint)

		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("invalid endpoint: %q", o.searchEndpoint))
		}
		endpoint = hostUrl.String()
	}

	c, err := api_client.New(endpoint)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to create connection to %q", o.searchEndpoint))
	}

	if common.IsDockerHub(endpoint) {
		common.CheckDockerHubLibrary(endpoint, q)
	}

	tags, err := c.ListTags(q.GetImageName(), o.fetchSize)

	if err != nil {
		fmt.Printf("%s", err)
		return fmt.Errorf("unable to list tags against %q", endpoint)
	}

	var e error
	for i, t := range *tags {
		if uint(i) > o.fetchSize {
			break
		}
		_, e = fmt.Fprintln(out, t)
	}

	return e
}
