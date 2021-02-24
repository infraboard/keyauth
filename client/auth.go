package client

import (
	"context"
)

// Authentication todo
type Authentication struct {
	clientID     string
	clientSecret string
}

// SetPasswordCredentials todo
func (a *Authentication) SetPasswordCredentials(clientID, clientSecret string) {
	a.clientID = clientID
	a.clientSecret = clientID
}

// GetRequestMetadata todo
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{"client_id": a.clientID, "client_secret": a.clientSecret}, nil
}

// RequireTransportSecurity todo
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
