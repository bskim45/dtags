package general

import "github.com/bskim45/dtags/common/api_client/base_client"

type GeneralApiClient struct {
	*base_client.BaseClient
}

func (c *GeneralApiClient) GetBaseClient() *base_client.BaseClient {
	return c.BaseClient
}
