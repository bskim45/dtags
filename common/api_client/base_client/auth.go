package base_client

import (
	"strings"

	auth "github.com/abbot/go-http-auth"
)

type AuthChallenge struct {
	Endpoint, Service, Scope string
}

func GetAuthChallenge(header string) *AuthChallenge {
	parts := strings.SplitN(header, " ", 2)

	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil
	}

	authPairs := auth.ParsePairs(parts[1])

	if authPairs == nil {
		return nil
	}

	return &AuthChallenge{
		authPairs["realm"], authPairs["service"], authPairs["scope"],
	}
}

func (c *BaseClient) Authenticate(challenge *AuthChallenge) error {
	token, err := c.GetToken(challenge)

	if err != nil {
		return err
	}

	c.Resty.SetAuthToken(token.Token)

	return nil
}
