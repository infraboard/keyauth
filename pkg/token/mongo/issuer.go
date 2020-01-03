package mongo

import (
	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
)

// TokenIssuer 基于该数据进行扩展
type TokenIssuer struct {
	*token.IssueTokenRequest
	app  application.Service
	user user.Service
}

func (i *TokenIssuer) checkClient() error {
	req := application.NewDescriptApplicationRequest()
	req.ClientID = i.ClientID
	app, err := i.app.DescriptionApplication(req)
	if err != nil {
		return err
	}

	return app.CheckClientSecret(i.ClientSecret)
}

func (i *TokenIssuer) checkUser() error {
	req := user.NewDescriptAccountRequest()
	req.Account = i.Username
	u, err := i.user.DescribeAccount(req)
	if err != nil {
		return err
	}

	return u.HashedPassword.CheckPassword(i.Password)
}

// IssueToken 颁发token
func (i *TokenIssuer) IssueToken() (*token.Token, error) {
	return nil, nil
}
