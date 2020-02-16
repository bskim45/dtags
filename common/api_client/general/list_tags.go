package general

import (
	"fmt"
	"path"
	"sort"

	"github.com/bskim45/dtags/common/api_client/base_client"
	"github.com/go-resty/resty/v2"
)

const ListTagsPath = "v2/%s/tags/list"

type listTagsResponse struct {
	ImageName string   `json:"name"`
	Tags      []string `json:"tags"`
}

func listTags(c *GeneralApiClient, name string) (*resty.Response, error) {
	res, err := c.R().
		SetQueryParams(map[string]string{
			// size is ignored
			//"n": "100",
		}).
		SetHeader("Accept", "application/json").
		SetResult(&listTagsResponse{}).
		Get(name)

	return res, err
}

// Performs a listing tag against the docker registry API
func (c *GeneralApiClient) ListTags(name string, _ uint) (*[]string, error) {
	p := c.BaseURL

	p.Path = path.Join(p.Path, fmt.Sprintf(ListTagsPath, name))

	res, err := listTags(c, p.String())

	if err != nil {
		return nil, err
	}

	// auth required
	if res.StatusCode() == 401 {
		challenge := base_client.GetAuthChallenge(res.Header().Get("Www-Authenticate"))

		if err = c.Authenticate(challenge); err != nil {
			return nil, err
		}

		res, err = listTags(c, p.String())

		if err != nil {
			return nil, err
		}

		if res.StatusCode() != 200 {
			return nil, fmt.Errorf("failed to fetch %s : %s", p.String(), res.Status())
		}
	} else if res.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to fetch %s : %s", p.String(), res.Status())
	}

	tagList := res.Result().(*listTagsResponse).Tags
	sort.Sort(sort.Reverse(sort.StringSlice(tagList)))

	var tagResult []string

	tagResult = append(tagResult, tagList...)

	return &tagResult, nil
}
