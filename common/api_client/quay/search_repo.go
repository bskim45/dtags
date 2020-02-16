package quay

import (
	"fmt"
	"path"
)

const SearchRepoPath = "api/v1/find/repositories"

type searchRepoResponseQuay struct {
	HasAdditional bool                   `json:"has_additional"`
	StartIndex    int                    `json:"start_index"`
	PageSize      int                    `json:"page_size"`
	Page          int                    `json:"page"`
	Results       []searchRepoResultQuay `json:"results"`
}

type searchRepoResultQuay struct {
	Name         string                  `json:"name"`
	Description  string                  `json:"description"`
	LastModified int64                   `json:"last_modified"`
	IsPublic     bool                    `json:"is_public"`
	Popularity   float32                 `json:"popularity"`
	Namespace    searchRepoNamespaceQuay `json:"namespace"`
	Score        int                     `json:"score"`
	Stars        int64                   `json:"stars"`
}

type searchRepoNamespaceQuay struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

func (c *QuayApiClient) SearchRepo(query string, _ uint) (*[]string, error) {
	p := c.BaseURL

	p.Path = path.Join(p.Path, SearchRepoPath)

	res, err := c.Resty.R().
		SetQueryParams(map[string]string{
			"query": query,
			//"page": 1,
		}).
		SetHeader("Accept", "application/json").
		SetResult(&searchRepoResponseQuay{}).
		Get(p.String())

	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to fetch %s : %s", p.String(), res.Status())
	}

	searchResult := res.Result().(*searchRepoResponseQuay).Results

	var repoList []string

	for _, imageHub := range searchResult {
		repoList = append(repoList, fmt.Sprintf("%s/%s", imageHub.Namespace.Name, imageHub.Name))
	}
	return &repoList, nil
}
