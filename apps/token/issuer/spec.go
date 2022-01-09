package issuer

import (
	"context"

	"github.com/infraboard/keyauth/apps/application"
	"github.com/infraboard/keyauth/apps/token"
)

// Issuer todo
type Issuer interface {
	CheckClient(ctx context.Context, clientID, clientSecret string) (*application.Application, error)
	IssueToken(context.Context, *token.IssueTokenRequest) (*token.Token, error)
}
