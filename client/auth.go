package client

import (
	"context"

	"github.com/infraboard/keyauth/common/header"
)

// Authentication todo
type Authentication struct {
	clientID     string
	clientSecret string
}

// SetClientCredentials todo
func (a *Authentication) SetClientCredentials(clientID, clientSecret string) {
	a.clientID = clientID
	a.clientSecret = clientSecret
}

// GetRequestMetadata todo
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{
		header.ClientHeaderKey: a.clientID,
		header.ClientSecretKey: a.clientSecret,
	}, nil
}

// RequireTransportSecurity todo
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
