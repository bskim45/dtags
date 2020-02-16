package base_client

import (
	"fmt"
	"time"
)

type Token struct {
	// https://docs.docker.com/registry/spec/auth/token/

	// An opaque Bearer token that clients should supply to subsequent requests in the Authorization header.
	Token string `json:"token"`

	// For compatibility with OAuth 2.0, we will also accept token under the name access_token.
	// At least one of these fields must be specified, but both may also appear (for compatibility with older clients).
	// When both are specified, they should be equivalent; if they differ the client's choice is undefined.
	AccessToken string `json:"access_token"`

	// (Optional) The duration in seconds since the token was issued that it will remain valid.
	// When omitted, this defaults to 60 seconds. For compatibility with older clients,
	// a token should never be returned with less than 60 seconds to live.
	ExpiresIn int `json:"expires_in,omitempty"`

	// (Optional) The RFC3339-serialized UTC standard time at which a given token was issued.
	// If issued_at is omitted, the expiration is from when the token exchange completed.
	IssuedAt time.Time `json:"issued_at,omitempty"`

	// (Optional) Token which can be used to get additional access tokens for the same subject with different scopes.
	// This token should be kept secure by the client and only sent to the authorization server which issues bearer
	// tokens. This field will only be set when `offline_token=true` is provided in the request.
	RefreshToken time.Time `json:"refresh_token,omitempty"`
}

func (c *BaseClient) GetToken(request *AuthChallenge) (*Token, error) {
	authUrl, err := validateUrl(request.Endpoint)

	if err != nil {
		return nil, err
	}

	res, err := c.R().
		SetQueryParams(map[string]string{
			"service": request.Service,
			"scope":   request.Scope,
		}).
		SetHeader(HEADER_ACCEPT, "application/json").
		SetResult(&Token{}).
		Get(authUrl.String())

	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to fetch %s : %s", request.Endpoint, res.Status())
	}

	return res.Result().(*Token), nil
}
