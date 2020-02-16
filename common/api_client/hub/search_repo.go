package hub

import (
	"fmt"
	"path"
	"strconv"
)

const SearchRepoPath = "v1/search/"

type searchRepoResponseHub struct {
	NumPages   int                   `json:"num_pages"`
	NumResults int                   `json:"num_results"`
	PageSize   int                   `json:"page_size"`
	Page       int                   `json:"page"`
	Query      string                `json:"query"`
	Results    []searchRepoResultHub `json:"results"`
}

type searchRepoResultHub struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StarCount   int64  `json:"star_count"`
	Istrusted   bool   `json:"is_trusted"`
	Isautomated bool   `json:"is_automated"`
	Isofficial  bool   `json:"is_official"`
}

func (c *HubApiClient) SearchRepo(query string, size uint) (*[]string, error) {
	p := c.BaseURL

	p.Path = path.Join(p.Path, SearchRepoPath)

	res, err := c.Resty.R().
		SetQueryParams(map[string]string{
			"q":         query,
			"page_size": strconv.FormatUint(uint64(size), 10),
		}).
		SetHeader("Accept", "application/json").
		SetResult(&searchRepoResponseHub{}).
		Get(p.String())

	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to fetch %s : %s", p.String(), res.Status())
	}

	searchResult := res.Result().(*searchRepoResponseHub).Results

	var imageList []string

	for _, imageHub := range searchResult {
		imageList = append(imageList, imageHub.Name)
	}
	return &imageList, nil
}
