package issuer

import (
	"fmt"
	"regexp"
	"strings"
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
	if pkg.Token == nil {
		return nil, fmt.Errorf("dependence token application is nil")
	}

	issuer := &issuer{
		user:    pkg.User,
		domain:  pkg.Domain,
		token:   pkg.Token,
		app:     pkg.Application,
		emailRE: regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9\.]+)\.([a-zA-Z0-9]+)`),
	}
	return issuer, nil
}

// TokenIssuer 基于该数据进行扩展
type issuer struct {
	app     application.Service
	token   token.Service
	user    user.Service
	domain  domain.Service
	emailRE *regexp.Regexp
}

func (i *issuer) checkUser(user, pass string) (*user.User, error) {
	u, err := i.getUser(user)
	if err != nil {
		return nil, err
	}
	if err := u.HashedPassword.CheckPassword(pass); err != nil {
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

	app, err := i.CheckClient(req.ClientID, req.ClientSecret)
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	switch req.GrantType {
	case token.PASSWORD:
		u, checkErr := i.checkUser(req.Username, req.Password)
		if checkErr != nil {
			return nil, exception.NewUnauthorized("user or password not connrect")
		}

		tk := i.issueUserToken(app, u, token.PASSWORD)
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
		validateReq.RefreshToken = req.RefreshToken
		tk, err := i.token.ValidateToken(validateReq)
		if err != nil {
			err = exception.NewUnauthorized(err.Error())
			return nil, err
		}
		if tk.AccessToken != req.AccessToken {
			return nil, exception.NewPermissionDeny("refresh_token's access_tken not connrect")
		}

		u, err := i.getUser(tk.Account)
		if err != nil {
			return nil, err
		}
		newTK := i.issueUserToken(app, u, token.REFRESH)
		newTK.DomainID = tk.DomainID

		revolkReq := token.NewRevolkTokenRequest(app.ClientID, app.ClientSecret)
		revolkReq.AccessToken = req.AccessToken
		if err := i.token.RevolkToken(revolkReq); err != nil {
			return nil, err
		}
		return newTK, nil
	case token.Access:
		validateReq := token.NewValidateTokenRequest()
		validateReq.AccessToken = req.AccessToken
		tk, err := i.token.ValidateToken(validateReq)
		if err != nil {
			return nil, exception.NewUnauthorized(err.Error())
		}
		u, err := i.getUser(tk.Account)
		if err != nil {
			return nil, err
		}
		newTK := i.issueUserToken(app, u, token.Access)
		newTK.DomainID = tk.DomainID
		return newTK, nil
	case token.LDAP:
		cn, dn, err := i.genBaseDN(req.Username)
		if err != nil {
			return nil, err
		}
		fmt.Println(dn, cn)
		return nil, exception.NewInternalServerError("not impl")
	case token.CLIENT:
		return nil, exception.NewInternalServerError("not impl")
	case token.AUTHCODE:
		return nil, exception.NewInternalServerError("not impl")
	default:
		return nil, exception.NewInternalServerError("unknown grant type %s", req.GrantType)
	}
}

func (i *issuer) genBaseDN(username string) (string, string, error) {
	match := i.emailRE.FindAllStringSubmatch(username, -1)
	if len(match) == 0 {
		return "", "", exception.NewBadRequest("ldap user name must like username@company.com")
	}

	sub := match[0]
	if len(sub) < 4 {
		return "", "", exception.NewBadRequest("ldap user name must like username@company.com")
	}

	upn := []string{}
	dns := []string{}
	upn = append(upn, "cn="+sub[1])
	for _, dn := range sub[2:] {
		dns = append(dns, "dc="+dn)
		upn = append(upn, "dc="+dn)
	}

	return strings.Join(upn, ","), strings.Join(dns, ","), nil
}

func (i *issuer) issueUserToken(app *application.Application, u *user.User, gt token.GrantType) *token.Token {
	tk := i.newBearToken(app, gt)
	tk.Account = u.Account
	tk.UserID = u.ID
	tk.UserType = u.Type
	return tk
}

func (i *issuer) newBearToken(app *application.Application, gt token.GrantType) *token.Token {
	now := time.Now()
	tk := &token.Token{
		Type:          token.Bearer,
		AccessToken:   token.MakeBearer(24),
		RefreshToken:  token.MakeBearer(32),
		CreatedAt:     ftime.T(now),
		ClientID:      app.ClientID,
		GrantType:     gt,
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
