package general

import (
	"fmt"
)

func (c *GeneralApiClient) SearchRepo(query string, size uint) (*[]string, error) {
	return nil, fmt.Errorf("search is unsupported: %s", c.BaseURL.String())
}
