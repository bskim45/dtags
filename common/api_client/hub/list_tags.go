package hub

import (
	"fmt"
	"path"
	"sort"
	"strconv"
	"time"
)

const ListTagsPath = "v2/repositories/%s/tags"

type listTagsResponseHub struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []tagResultHub `json:"results"`
}

type tagResultHub struct {
	Name        string           `json:"name"`
	Size        int64            `json:"full_size"`
	Images      []imageResultHub `json:"images"`
	Id          int64            `json:"id"`
	Repository  int64            `json:"repository"`
	Creator     int64            `json:"creator"`
	LastUpdater int64            `json:"last_updater"`
	LastUpdated time.Time        `json:"last_updated"`
	ImageId     int64            `json:"image_id"`
	IsV2        bool             `json:"v2"`
}

type imageResultHub struct {
	Size         int64  `json:"size"`
	Architecture string `json:"architecture"`
	Variant      string `json:"variant"`
	Features     string `json:"features"`
	Os           string `json:"os"`
	OsVersion    string `json:"os_version"`
	OsFeatures   string `json:"os_features"`
}

func (c *HubApiClient) ListTags(name string, size uint) (*[]string, error) {
	p := c.BaseURL

	p.Path = path.Join(p.Path, fmt.Sprintf(ListTagsPath, name)) + "/"

	res, err := c.R().
		SetQueryParams(map[string]string{
			"page_size": strconv.FormatUint(uint64(size), 10),
			"page":      "1",
		}).
		SetHeader("Accept", "application/json").
		SetResult(&listTagsResponseHub{}).
		Get(p.String())

	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to fetch %s : %s", p.String(), res.Status())
	}

	tagList := res.Result().(*listTagsResponseHub).Results

	sort.Slice(tagList, func(i, j int) bool {
		// reverse sort on last-updated
		return tagList[i].LastUpdated.After(tagList[j].LastUpdated)
	})

	var tagResult []string

	for _, tagHub := range tagList {
		tagResult = append(tagResult, tagHub.Name)
	}
	return &tagResult, nil
}
