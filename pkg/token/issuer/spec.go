package issuer

import (
	"context"

	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/token"
)

// Issuer todo
type Issuer interface {
	CheckClient(ctx context.Context, clientID, clientSecret string) (*application.Application, error)
	IssueToken(context.Context, *token.IssueTokenRequest) (*token.Token, error)
}
