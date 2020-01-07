package mongo

import (
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
)

func (s *service) newTokenIssuer(req *token.IssueTokenRequest) (*TokenIssuer, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	issuer := &TokenIssuer{
		IssueTokenRequest: req,
		clientChecker:     newClientChecker(s.app),
		user:              s.user,
		token:             s,
	}
	return issuer, nil
}

// TokenIssuer 基于该数据进行扩展
type TokenIssuer struct {
	*token.IssueTokenRequest
	*clientChecker
	token *service
	user  user.Service
}

func (i *TokenIssuer) checkUser() (*user.User, error) {
	req := user.NewDescriptAccountRequest()
	req.Account = i.Username
	u, err := i.user.DescribeAccount(req)
	if err != nil {
		return nil, err
	}
	if err := u.HashedPassword.CheckPassword(i.Password); err != nil {
		return nil, err
	}

	return u, nil
}

// IssueToken 颁发token
func (i *TokenIssuer) IssueToken() (tk *token.Token, err error) {
	app, err := i.CheckClient(i.ClientID, i.ClientSecret)
	if err != nil {
		err = exception.NewUnauthorized(err.Error())
		return
	}

	switch i.GrantType {
	case token.PASSWORD:
		u, checkErr := i.checkUser()
		if checkErr != nil {
			err = exception.NewUnauthorized("user or password not connrect")
			return
		}

		tk = i.issuePasswordToken(app, u.ID)
		return
	case token.REFRESH:
		tk, err = i.token.queryToken(newQueryTokenRequestWithRefresh(i.RefreshToken))
		if err != nil {
			err = exception.NewUnauthorized(err.Error())
			return
		}
		if tk.CheckRefreshIsExpired() {
			err = exception.NewUnauthorized("")
			return
		}
		tk = i.issuePasswordToken(app, tk.UserID)
	case token.CLIENT:
	case token.AUTHCODE:
	default:
		err = exception.NewInternalServerError("unknown grant type %s", i.GrantType)
		return
	}

	return
}

func (i *TokenIssuer) issuePasswordToken(app *application.Application, userID string) *token.Token {
	tk := i.newBearToken(app)
	tk.UserID = userID
	return tk
}

func (i *TokenIssuer) refreshToken(tk *token.Token) {
	now := time.Now()
	tk.AccessToken = token.MakeBearer(24)
	tk.RefreshToken = token.MakeBearer(32)
	tk.CreatedAt = ftime.T(now)
}

func (i *TokenIssuer) newBearToken(app *application.Application) *token.Token {
	now := time.Now()
	accessExpire := now.Add(time.Duration(app.AccessTokenExpireSecond) * time.Second)
	refreshExpir := now.Add(time.Duration(app.RefreshTokenExpiredSecond) * time.Second)
	return &token.Token{
		Type:             token.Bearer,
		AccessToken:      token.MakeBearer(24),
		RefreshToken:     token.MakeBearer(32),
		CreatedAt:        ftime.T(now),
		ClientID:         i.ClientID,
		GrantType:        i.GrantType,
		AccessExpiredAt:  ftime.T(accessExpire),
		RefreshExpiredAt: ftime.T(refreshExpir),
		ApplicationID:    app.ID,
	}
}

func newClientChecker(app application.Service) *clientChecker {
	return &clientChecker{app}
}

// clientChecker 检测client正确性
type clientChecker struct {
	application.Service
}

func (ck *clientChecker) CheckClient(clientID, clientSecret string) (*application.Application, error) {
	req := application.NewDescriptApplicationRequest()
	req.ClientID = clientID
	app, err := ck.DescriptionApplication(req)
	if err != nil {
		return nil, err
	}

	if err := app.CheckClientSecret(clientSecret); err != nil {
		return nil, err
	}

	return app, nil
}
