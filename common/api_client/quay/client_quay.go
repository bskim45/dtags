package quay

import "github.com/bskim45/dtags/common/api_client/base_client"

type QuayApiClient struct {
	*base_client.BaseClient
}

func (c *QuayApiClient) GetBaseClient() *base_client.BaseClient {
	return c.BaseClient
}
