package common

import (
	"fmt"
	"net/url"
	"strings"
)

type ImageName struct {
	Endpoint string
	Repo     string
	Name     string
}

func (i ImageName) GetImageName() string {
	return fmt.Sprintf("%s/%s", i.Repo, i.Name)
}

func ParseName(name string) *ImageName {
	sp := strings.Split(name, "/")

	switch len(sp) {
	case 1:
		return &ImageName{
			Endpoint: "",
			Repo:     "",
			Name:     sp[0],
		}
	case 2:
		return &ImageName{
			Endpoint: "",
			Repo:     sp[0],
			Name:     sp[1],
		}
	case 3:
		return &ImageName{
			Endpoint: sp[0],
			Repo:     sp[1],
			Name:     sp[2],
		}
	default:
		return &ImageName{
			Endpoint: "",
			Repo:     "",
			Name:     sp[0],
		}
	}
}

func IsDockerHub(endpoint string) bool {
	return strings.Contains(endpoint, "docker.com") || strings.Contains(endpoint, "docker.io")
}

func IsQuay(endpoint string) bool {
	return strings.Contains(endpoint, "quay.io")
}

func CheckDockerHubLibrary(endpoint string, q *ImageName) {
	if IsDockerHub(endpoint) && q.Repo == "" {
		q.Repo = "library"
	}
}

func BuildUrl(endpoint string) (*url.URL, error) {
	newUrl, err := url.Parse(endpoint)

	if err != nil {
		return nil, err
	}

	if newUrl.Host == "" {
		newUrl = &url.URL{
			Scheme: "https",
			Host:   endpoint,
		}
	}

	if newUrl.Scheme == "" {
		newUrl.Scheme = "https"
	}

	return newUrl, nil
}
