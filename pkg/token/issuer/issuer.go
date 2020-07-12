package issuer

import (
	"fmt"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

// NewTokenIssuer todo
func NewTokenIssuer() (Issuer, error) {
	if pkg.Application == nil {
		return nil, fmt.Errorf("dependence service application is nil")
	}
	if pkg.User == nil {
		return nil, fmt.Errorf("dependence user application is nil")
	}
	if pkg.Domain == nil {
		return nil, fmt.Errorf("dependence domain application is nil")
	}

	issuer := &issuer{
		user:   pkg.User,
		domain: pkg.Domain,
		token:  pkg.Token,
	}
	return issuer, nil
}

// TokenIssuer 基于该数据进行扩展
type issuer struct {
	*token.IssueTokenRequest
	app    application.Service
	token  token.Service
	user   user.Service
	domain domain.Service
}

func (i *issuer) checkUser() (*user.User, error) {
	u, err := i.getUser(i.Username)
	if err != nil {
		return nil, err
	}
	if err := u.HashedPassword.CheckPassword(i.Password); err != nil {
		return nil, err
	}
	return u, nil
}

func (i *issuer) getUser(name string) (*user.User, error) {
	req := user.NewDescriptAccountRequest()
	req.Account = name
	return i.user.DescribeAccount(req)
}

func (i *issuer) setTokenDomain(tk *token.Token) error {
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
func (i *issuer) IssueToken(req *token.IssueTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

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
		validateReq := token.NewValidateTokenRequest()
		validateReq.RefreshToken = i.RefreshToken
		tk, err := i.token.ValidateToken(validateReq)
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
		revolkReq := token.NewRevolkTokenRequest(app.ClientID, app.ClientSecret)
		if err := i.token.RevolkToken(revolkReq); err != nil {
			return nil, err
		}
		return tk, nil
	case token.Access:
		validateReq := token.NewValidateTokenRequest()
		validateReq.AccessToken = i.AccessToken
		tk, err := i.token.ValidateToken(validateReq)
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

func (i *issuer) issueUserToken(app *application.Application, u *user.User) *token.Token {
	tk := i.newBearToken(app)
	tk.Account = u.Account
	tk.UserID = u.ID
	tk.UserType = u.Type
	return tk
}

func (i *issuer) refreshToken(tk *token.Token) {
	now := time.Now()
	tk.AccessToken = token.MakeBearer(24)
	tk.RefreshToken = token.MakeBearer(32)
	tk.CreatedAt = ftime.T(now)
}

func (i *issuer) newBearToken(app *application.Application) *token.Token {
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
