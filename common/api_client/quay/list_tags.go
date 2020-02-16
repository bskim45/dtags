package quay

import (
	"fmt"
	"path"
	"sort"
	"strconv"
)

const ListTagsPath = "api/v1/repository/%s/tag/"

type listTagsResponseQuay struct {
	HasAdditional bool            `json:"has_additional"`
	Page          int             `json:"page"`
	Tags          []tagResultQuay `json:"tags"`
}

type tagResultQuay struct {
	Name           string   `json:"name"`
	Reversion      bool     `json:"reversion"`
	StartTs        uint64   `json:"start_ts"`
	ImageId        string   `json:"image_id"`
	LastModified   QuayTime `json:"last_modified"`
	ManifestDigest string   `json:"manifest_digest"`
	DockerImageId  string   `json:"docker_image_id"`
	IsManifestList bool     `json:"is_manifest_list"`
	Size           int64    `json:"full_size"`
}

func (c *QuayApiClient) ListTags(name string, size uint) (*[]string, error) {
	p := c.BaseURL

	p.Path = path.Join(p.Path, fmt.Sprintf(ListTagsPath, name)) + "/"

	res, err := c.R().
		SetQueryParams(map[string]string{
			"limit":          strconv.FormatUint(uint64(size), 10),
			"page":           "1",
			"onlyActiveTags": "true",
		}).
		SetHeader("Accept", "application/json").
		SetResult(&listTagsResponseQuay{}).
		Get(p.String())

	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to fetch %s : %s", p.String(), res.Status())
	}

	tagList := res.Result().(*listTagsResponseQuay).Tags

	sort.Slice(tagList, func(i, j int) bool {
		// reverse sort on last-updated
		return tagList[i].LastModified.Time.After(tagList[j].LastModified.Time)
	})

	var tagResult []string

	for _, tagQuay := range tagList {
		tagResult = append(tagResult, tagQuay.Name)
	}
	return &tagResult, nil
}
