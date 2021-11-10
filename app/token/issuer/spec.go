package issuer

import (
	"context"

	"github.com/infraboard/keyauth/app/application"
	"github.com/infraboard/keyauth/app/token"
)

// Issuer todo
type Issuer interface {
	CheckClient(ctx context.Context, clientID, clientSecret string) (*application.Application, error)
	IssueToken(context.Context, *token.IssueTokenRequest) (*token.Token, error)
}
