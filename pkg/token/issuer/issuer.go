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
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/common/password"
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
	if pkg.ApplicationUser == nil {
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
		app:     pkg.ApplicationUser,
		emailRE: regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9\.]+)\.([a-zA-Z0-9]+)`),
		log:     zap.L().Named("Token Issuer"),
	}
	return issuer, nil
}

// TokenIssuer 基于该数据进行扩展
type issuer struct {
	app     application.UserServiceServer
	token   token.TokenServiceServer
	user    user.Service
	domain  domain.DomainServiceServer
	ldap    provider.LDAP
	emailRE *regexp.Regexp
	log     logger.Logger
}

func (i *issuer) checkUserPass(user, pass string) (*user.User, error) {
	u, err := i.getUser(user)
	if err != nil {
		return nil, err
	}

	if err := u.HashedPassword.CheckPassword(pass); err != nil {
		return nil, err
	}
	return u, nil
}

func (i *issuer) checkUserPassExpired(u *user.User) error {
	d, err := i.getDomain(u)
	if err != nil {
		return err
	}

	// 检测密码是否过期
	err = d.SecuritySetting.PasswordSecurity.IsPasswordExpired(u.HashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (i *issuer) getUser(name string) (*user.User, error) {
	req := user.NewDescriptAccountRequest()
	req.Account = name
	return i.user.DescribeAccount(req)
}

func (i *issuer) getDomain(u *user.User) (*domain.Domain, error) {
	req := domain.NewDescribeDomainRequestWithName(u.Domain)
	return i.domain.DescribeDomain(pkg.GetInternalAdminTokenCtx(u.Account), req)
}

func (i *issuer) setTokenDomain(tk *token.Token) error {
	// 获取最近1个
	req := domain.NewQueryDomainRequest(request.NewPageRequest(1, 1))

	domains, err := i.domain.QueryDomain(pkg.GetInternalAdminTokenCtx(tk.Account), req)
	if err != nil {
		return err
	}

	if domains.Length() > 0 {
		tk.Domain = domains.Items[0].Data.Name
	}

	return nil
}

// IssueToken 颁发token
func (i *issuer) IssueToken(req *token.IssueTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	app, err := i.CheckClient(req.ClientId, req.ClientSecret)
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	switch req.GrantType {
	case token.GrantType_PASSWORD:
		u, checkErr := i.checkUserPass(req.Username, req.Password)
		if checkErr != nil {
			i.log.Debugf("issue password token error, %s", checkErr)
			return nil, exception.NewUnauthorized("user or password not connrect")
		}

		if err := i.checkUserPassExpired(u); err != nil {
			i.log.Debugf("issue password token error, %s", err)
			if v, ok := err.(exception.APIException); ok {
				v.WithData(u.Account)
			}
			return nil, err
		}

		tk := i.issueUserToken(app, u, token.GrantType_PASSWORD)
		switch u.Type {
		case types.UserType_SUPPER, types.UserType_PRIMARY:
			err := i.setTokenDomain(tk)
			if err != nil {
				return nil, fmt.Errorf("set token domain error, %s", err)
			}
		case types.UserType_SERVICE, types.UserType_SUB:
			tk.Domain = u.Domain
		}

		return tk, nil
	case token.GrantType_REFRESH:
		validateReq := token.NewValidateTokenRequest()
		validateReq.RefreshToken = req.RefreshToken
		tk, err := i.token.ValidateToken(nil, validateReq)
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
		newTK := i.issueUserToken(app, u, token.GrantType_REFRESH)
		newTK.Domain = tk.Domain
		newTK.StartGrantType = tk.GetStartGrantType()
		newTK.SessionId = tk.SessionId

		revolkReq := token.NewRevolkTokenRequest(app.ClientId, app.ClientSecret)
		revolkReq.AccessToken = req.AccessToken
		revolkReq.LogoutSession = false
		if _, err := i.token.RevolkToken(nil, revolkReq); err != nil {
			return nil, err
		}
		return newTK, nil
	case token.GrantType_ACCESS:
		validateReq := token.NewValidateTokenRequest()
		validateReq.AccessToken = req.AccessToken
		tk, err := i.token.ValidateToken(nil, validateReq)
		if err != nil {
			return nil, exception.NewUnauthorized(err.Error())
		}
		u, err := i.getUser(tk.Account)
		if err != nil {
			return nil, err
		}
		newTK := i.issueUserToken(app, u, token.GrantType_ACCESS)
		newTK.Domain = tk.Domain
		return newTK, nil
	case token.GrantType_LDAP:
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
		u, err := i.syncLDAPUser(mockPrimary, req.Username)
		if err != nil {
			return nil, err
		}
		newTK := i.issueUserToken(app, u, token.GrantType_LDAP)
		newTK.Domain = ldapConf.Domain
		return newTK, nil
	case token.GrantType_CLIENT:
		return nil, exception.NewInternalServerError("not impl")
	case token.GrantType_AUTH_CODE:
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

	if u != nil && u.Type.IsIn(types.UserType_PRIMARY, types.UserType_SUPPER) {
		return nil, exception.NewBadRequest("用户名和主账号用户名冲突, 请修改")
	}

	if err != nil {
		if exception.IsNotFoundError(err) {
			req := user.NewCreateUserRequestWithLDAPSync(userName, i.randomPass(), tk)
			u, err = i.user.CreateAccount(types.UserType_SUB, req)
			if err != nil {
				return nil, err
			}
		}
		return u, err
	}

	return u, nil
}

func (i *issuer) randomPass() string {
	rpass, err := password.NewWithDefault().Generate()
	if err != nil {
		i.log.Warnf("generate random password error, %s, use uuid for random password", err)
	}
	if rpass != nil {
		return *rpass
	}

	return xid.New().String()
}

func (i *issuer) mockBuildInToken(app *application.Application, userName, domainID string) *token.Token {
	tk := i.newBearToken(app, token.GrantType_LDAP)
	tk.Account = userName
	tk.UserType = types.UserType_PRIMARY
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
		Type:            token.TokenType_BEARER,
		AccessToken:     token.MakeBearer(24),
		RefreshToken:    token.MakeBearer(32),
		CreateAt:        now.UnixNano() / 1000000,
		ClientId:        app.ClientId,
		GrantType:       gt,
		ApplicationId:   app.Id,
		ApplicationName: app.Name,
	}

	if app.AccessTokenExpireSecond != 0 {
		accessExpire := now.Add(time.Duration(app.AccessTokenExpireSecond) * time.Second)
		tk.AccessExpiredAt = accessExpire.UnixNano() / 1000000
	}

	if app.RefreshTokenExpireSecond != 0 {
		refreshExpir := now.Add(time.Duration(app.RefreshTokenExpireSecond) * time.Second)
		tk.RefreshExpiredAt = refreshExpir.UnixNano() / 1000000
	}

	return tk
}
