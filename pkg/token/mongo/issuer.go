package mongo

import (
	"fmt"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func (s *service) newTokenIssuer(req *token.IssueTokenRequest) (*TokenIssuer, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	issuer := &TokenIssuer{
		IssueTokenRequest: req,
		clientChecker:     newClientChecker(s.app),
		user:              s.user,
		domain:            s.domain,
		token:             s,
	}
	return issuer, nil
}

// TokenIssuer 基于该数据进行扩展
type TokenIssuer struct {
	*token.IssueTokenRequest
	*clientChecker
	token  *service
	user   user.Service
	domain domain.Service
}

func (i *TokenIssuer) checkUser() (*user.User, error) {
	u, err := i.getUser(i.Username)
	if err != nil {
		return nil, err
	}
	if err := u.HashedPassword.CheckPassword(i.Password); err != nil {
		return nil, err
	}
	return u, nil
}

func (i *TokenIssuer) getUser(name string) (*user.User, error) {
	req := user.NewDescriptAccountRequest()
	req.Account = name
	return i.user.DescribeAccount(req)
}

func (i *TokenIssuer) setTokenDomain(tk *token.Token) error {
	// 获取最近1个
	req := domain.NewQueryDomainRequest(request.NewPageRequest(1, 1))
	req.WithToken(tk)

	domains, err := i.domain.QueryDomain(req)
	if err != nil {
		return err
	}

	if domains.Length() > 0 {
		tk.DomainID = domains.Items[0].ID
	}

	return nil
}

// IssueToken 颁发token
func (i *TokenIssuer) IssueToken() (*token.Token, error) {
	app, err := i.CheckClient(i.ClientID, i.ClientSecret)
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	switch i.GrantType {
	case token.PASSWORD:
		u, checkErr := i.checkUser()
		if checkErr != nil {
			return nil, exception.NewUnauthorized("user or password not connrect")
		}

		tk := i.issueUserToken(app, u)
		switch u.Type {
		case types.SupperAccount, types.PrimaryAccount:
			err := i.setTokenDomain(tk)
			if err != nil {
				return nil, fmt.Errorf("set token domain error, %s", err)
			}
		case types.ServiceAccount, types.SubAccount:
			tk.DomainID = u.DomainID
		}

		return tk, nil
	case token.REFRESH:
		descReq := newDescribeTokenRequestWithRefresh(i.RefreshToken)
		tk, err := i.token.describeToken(descReq)
		if err != nil {
			err = exception.NewUnauthorized(err.Error())
			return nil, err
		}
		if tk.CheckRefreshIsExpired() {
			return nil, exception.NewRefreshTokenExpired("refresh token is expoired")
		}
		u, err := i.getUser(tk.Account)
		if err != nil {
			return nil, err
		}
		tk = i.issueUserToken(app, u)
		if err := i.token.destoryToken(descReq); err != nil {
			return nil, err
		}
		return tk, nil
	case token.Access:
		descReq := newDescribeTokenRequestWithAccess(i.AccessToken)
		tk, err := i.token.describeToken(descReq)
		if err != nil {
			return nil, exception.NewUnauthorized(err.Error())
		}
		if tk.CheckRefreshIsExpired() {
			return nil, exception.NewRefreshTokenExpired("access token is expoired")
		}
		u, err := i.getUser(tk.Account)
		if err != nil {
			return nil, err
		}
		tk = i.issueUserToken(app, u)
		return tk, nil
	case token.CLIENT:
		return nil, exception.NewInternalServerError("not impl")
	case token.AUTHCODE:
		return nil, exception.NewInternalServerError("not impl")
	default:
		return nil, exception.NewInternalServerError("unknown grant type %s", i.GrantType)
	}
}

func (i *TokenIssuer) issueUserToken(app *application.Application, u *user.User) *token.Token {
	tk := i.newBearToken(app)
	tk.Account = u.Account
	tk.UserID = u.ID
	tk.UserType = u.Type
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
	tk := &token.Token{
		Type:          token.Bearer,
		AccessToken:   token.MakeBearer(24),
		RefreshToken:  token.MakeBearer(32),
		CreatedAt:     ftime.T(now),
		ClientID:      i.ClientID,
		GrantType:     i.GrantType,
		ApplicationID: app.ID,
	}

	if app.AccessTokenExpireSecond != 0 {
		accessExpire := now.Add(time.Duration(app.AccessTokenExpireSecond) * time.Second)
		tk.AccessExpiredAt = ftime.T(accessExpire)
	}

	if app.RefreshTokenExpiredSecond != 0 {
		refreshExpir := now.Add(time.Duration(app.RefreshTokenExpiredSecond) * time.Second)
		tk.RefreshExpiredAt = ftime.T(refreshExpir)
	}

	return tk
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
