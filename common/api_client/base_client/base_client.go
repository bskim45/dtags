package base_client

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/bskim45/dtags/common"
	"github.com/go-resty/resty/v2"
)

//noinspection ALL
var (
	HEADER_USER_AGENT       = http.CanonicalHeaderKey("User-Agent")
	HEADER_ACCEPT           = http.CanonicalHeaderKey("Accept")
	HEADER_CONTENT_TYPE     = http.CanonicalHeaderKey("Content-Type")
	HEADER_WWW_AUTHENTICATE = http.CanonicalHeaderKey("Www-Authenticate")
	HEADER_AUTHORIZATION    = http.CanonicalHeaderKey("Authorization")
)

type BaseClient struct {
	// The base URL for requests
	BaseURL *url.URL

	// Resty client
	Resty *resty.Client
}

func New(u string) (*BaseClient, error) {
	baseUrl, err := validateUrl(u)

	if err != nil {
		return nil, err
	}

	if baseUrl.Scheme == "" {
		baseUrl.Scheme = "https"
	}

	c := resty.New().
		SetHeader(HEADER_USER_AGENT, common.GetUserAgent()).
		SetTimeout(30 * time.Second)

	//c.SetDebug(true)

	return &BaseClient{
		BaseURL: baseUrl,
		Resty:   c,
	}, nil
}

func (c *BaseClient) R() *resty.Request {
	return c.Resty.R()
}

func validateUrl(u string) (*url.URL, error) {
	// Check if it is parsable
	p, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	// Check that a host is attached
	if p.Hostname() == "" {
		return nil, errors.New("hostname not provided")
	}

	return p, nil
}
