package issuer

import (
	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/token"
)

// Issuer todo
type Issuer interface {
	CheckClient(clientID, clientSecret string) (*application.Application, error)
	IssueToken(req *token.IssueTokenRequest) (*token.Token, error)
}
