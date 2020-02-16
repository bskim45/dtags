package hub

import "github.com/bskim45/dtags/common/api_client/base_client"

type HubApiClient struct {
	*base_client.BaseClient
}

func (c *HubApiClient) GetBaseClient() *base_client.BaseClient {
	return c.BaseClient
}
