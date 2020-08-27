package issuer

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/provider"
	"github.com/infraboard/keyauth/pkg/provider/ldap"
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
	if pkg.LDAP == nil {
		return nil, fmt.Errorf("dependence ldap application is nil")
	}

	issuer := &issuer{
		user:    pkg.User,
		domain:  pkg.Domain,
		token:   pkg.Token,
		ldap:    pkg.LDAP,
		app:     pkg.Application,
		emailRE: regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9\.]+)\.([a-zA-Z0-9]+)`),
		log:     zap.L().Named("Token Issuer"),
	}
	return issuer, nil
}

// TokenIssuer 基于该数据进行扩展
type issuer struct {
	app     application.Service
	token   token.Service
	user    user.Service
	domain  domain.Service
	ldap    provider.LDAP
	emailRE *regexp.Regexp
	log     logger.Logger
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
		tk.Domain = domains.Items[0].Name
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
			i.log.Debugf("issue password token error, %s", checkErr)
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
			tk.Domain = u.Domain
		}

		return tk, nil
	case token.REFRESH:
		validateReq := token.NewValidateTokenRequest()
		validateReq.RefreshToken = req.RefreshToken
		tk, err := i.token.ValidateToken(validateReq)
		if err != nil {
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
		newTK.Domain = tk.Domain
		newTK.StartGrantType = tk.GrantType

		revolkReq := token.NewRevolkTokenRequest(app.ClientID, app.ClientSecret)
		revolkReq.AccessToken = req.AccessToken
		if err := i.token.RevolkToken(revolkReq); err != nil {
			return nil, err
		}
		return newTK, nil
	case token.ACCESS:
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
		newTK := i.issueUserToken(app, u, token.ACCESS)
		newTK.Domain = tk.Domain
		return newTK, nil
	case token.LDAP:
		userName, dn, err := i.genBaseDN(req.Username)
		if err != nil {
			return nil, err
		}

		descReq := provider.NewDescribeLDAPConfigWithBaseDN(dn)
		ldapConf, err := i.ldap.DescribeConfig(descReq)
		if err != nil {
			return nil, err
		}
		pv := ldap.NewProvider(ldapConf.Config)
		ok, err := pv.CheckUserPassword(userName, req.Password)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, exception.NewUnauthorized("用户名或者密码不对")
		}
		mockPrimary := i.mockBuildInToken(app, userName, ldapConf.Domain)
		u, err := i.syncLDAPUser(mockPrimary, userName)
		if err != nil {
			return nil, err
		}
		newTK := i.issueUserToken(app, u, token.LDAP)
		newTK.Domain = ldapConf.Domain
		return newTK, nil
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

	dns := []string{}
	for _, dn := range sub[2:] {
		dns = append(dns, "dc="+dn)
	}

	return sub[1], strings.Join(dns, ","), nil
}

func (i *issuer) syncLDAPUser(tk *token.Token, userName string) (*user.User, error) {
	descUser := user.NewDescriptAccountRequestWithAccount(userName)
	u, err := i.user.DescribeAccount(descUser)
	if u.Type.Is(types.PrimaryAccount, types.SupperAccount) {
		return nil, exception.NewBadRequest("用户名和主账号用户名冲突, 请修改")
	}
	if err != nil {
		if exception.IsNotFoundError(err) {
			req := user.NewCreateUserRequest()
			req.WithToken(tk)
			u, err = i.user.CreateAccount(types.SubAccount, req)
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	return u, nil
}

func (i *issuer) mockBuildInToken(app *application.Application, userName, domainID string) *token.Token {
	tk := i.newBearToken(app, token.LDAP)
	tk.Account = userName
	tk.UserType = types.PrimaryAccount
	tk.Domain = domainID
	return tk
}

func (i *issuer) issueUserToken(app *application.Application, u *user.User, gt token.GrantType) *token.Token {
	tk := i.newBearToken(app, gt)
	tk.Account = u.Account
	tk.UserType = u.Type
	return tk
}

func (i *issuer) newBearToken(app *application.Application, gt token.GrantType) *token.Token {
	now := time.Now()
	tk := &token.Token{
		Type:            token.Bearer,
		AccessToken:     token.MakeBearer(24),
		RefreshToken:    token.MakeBearer(32),
		CreatedAt:       ftime.T(now),
		ClientID:        app.ClientID,
		GrantType:       gt,
		ApplicationID:   app.ID,
		ApplicationName: app.Name,
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
