package api_client

import (
	"github.com/bskim45/dtags/common"
	"github.com/bskim45/dtags/common/api_client/base_client"
	"github.com/bskim45/dtags/common/api_client/general"
	"github.com/bskim45/dtags/common/api_client/hub"
	"github.com/bskim45/dtags/common/api_client/quay"
)

type ApiClient interface {
	GetBaseClient() *base_client.BaseClient
	ListTags(name string, size uint) (*[]string, error)
	SearchRepo(query string, size uint) (*[]string, error)
}

func New(url string) (ApiClient, error) {
	baseClient, err := base_client.New(url)

	if err != nil {
		return nil, err
	}

	u := baseClient.BaseURL.String()

	switch {
	case common.IsDockerHub(u):
		return &hub.HubApiClient{
			BaseClient: baseClient,
		}, nil
	case common.IsQuay(u):
		return &quay.QuayApiClient{
			BaseClient: baseClient,
		}, nil
	default:
		return &general.GeneralApiClient{
			BaseClient: baseClient,
		}, nil
	}
}
