package issuer

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/rs/xid"

	wechatWork "github.com/infraboard/keyauth/apps/wxwork"

	"github.com/infraboard/keyauth/apps/application"
	"github.com/infraboard/keyauth/apps/domain"
	"github.com/infraboard/keyauth/apps/provider"
	"github.com/infraboard/keyauth/apps/provider/auth/ldap"
	"github.com/infraboard/keyauth/apps/provider/auth/wxwork"
	"github.com/infraboard/keyauth/apps/token"
	"github.com/infraboard/keyauth/apps/user"
	"github.com/infraboard/keyauth/apps/user/types"
	"github.com/infraboard/keyauth/common/password"
)

// NewTokenIssuer todo
func NewTokenIssuer() (Issuer, error) {
	issuer := &issuer{
		user:       app.GetGrpcApp(user.AppName).(user.ServiceServer),
		domain:     app.GetGrpcApp(domain.AppName).(domain.ServiceServer),
		token:      app.GetGrpcApp(token.AppName).(token.ServiceServer),
		ldap:       app.GetInternalApp(provider.AppName).(provider.LDAP),
		wechatWork: app.GetInternalApp(wechatWork.AppName).(wechatWork.WechatWork),
		app:        app.GetGrpcApp(application.AppName).(application.ServiceServer),
		emailRE:    regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9\.]+)\.([a-zA-Z0-9]+)`),
		log:        zap.L().Named("Token Issuer"),
	}
	return issuer, nil
}

// TokenIssuer 基于该数据进行扩展
type issuer struct {
	app        application.ServiceServer
	token      token.ServiceServer
	user       user.ServiceServer
	domain     domain.ServiceServer
	ldap       provider.LDAP
	wechatWork wechatWork.WechatWork
	emailRE    *regexp.Regexp
	log        logger.Logger
}

func (i *issuer) checkUserPass(ctx context.Context, user, pass string) (*user.User, error) {
	u, err := i.getUser(ctx, user)
	if err != nil {
		return nil, err
	}

	if err := u.HashedPassword.CheckPassword(pass); err != nil {
		return nil, err
	}
	return u, nil
}

func (i *issuer) checkUserPassExpired(ctx context.Context, u *user.User) error {
	d, err := i.getDomain(ctx, u)
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

func (i *issuer) getUser(ctx context.Context, name string) (*user.User, error) {
	req := user.NewDescriptAccountRequest()
	req.Account = name
	return i.user.DescribeAccount(ctx, req)
}

func (i *issuer) getDomain(ctx context.Context, u *user.User) (*domain.Domain, error) {
	req := domain.NewDescribeDomainRequestWithName(u.Domain)
	return i.domain.DescribeDomain(ctx, req)
}

func (i *issuer) setTokenDomain(ctx context.Context, tk *token.Token) error {
	// 获取最近1个
	req := domain.NewQueryDomainRequest(request.NewPageRequest(1, 1))
	domains, err := i.domain.QueryDomain(ctx, req)
	if err != nil {
		return err
	}

	if domains.Length() > 0 {
		tk.Domain = domains.Items[0].Name
	}

	return nil
}

// IssueToken 颁发token
func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// 校验应用端的合法性
	app, err := i.CheckClient(ctx, req.ClientId, req.ClientSecret)
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	// 校验用户的身份
	switch req.GrantType {
	case token.GrantType_PASSWORD:
		// 判断用户的密码是否正确
		u, checkErr := i.checkUserPass(ctx, req.Username, req.Password)
		if checkErr != nil {
			i.log.Debugf("issue password token error, %s", checkErr)
			return nil, exception.NewUnauthorized("user or password not connrect")
		}

		// 校验用的密码是否过期
		if err := i.checkUserPassExpired(ctx, u); err != nil {
			i.log.Debugf("issue password token error, %s", err)
			if v, ok := err.(exception.APIException); ok {
				v.WithData(u.Account)
			}
			return nil, err
		}

		// 颁发给用户一个token
		tk := i.issueUserToken(app, u, token.GrantType_PASSWORD)
		switch u.Type {
		case types.UserType_SUPPER, types.UserType_PRIMARY:
			err := i.setTokenDomain(ctx, tk)
			if err != nil {
				return nil, fmt.Errorf("set token domain error, %s", err)
			}
			tk.Domain = u.Domain
		default:
			tk.Domain = u.Domain
		}

		return tk, nil
	case token.GrantType_REFRESH:
		validateReq := token.NewValidateTokenRequest()
		validateReq.RefreshToken = req.RefreshToken
		tk, err := i.token.ValidateToken(context.Background(), validateReq)
		if err != nil {
			return nil, err
		}
		if tk.AccessToken != req.AccessToken {
			return nil, exception.NewPermissionDeny("refresh_token's access_tken not connrect")
		}

		u, err := i.getUser(ctx, tk.Account)
		if err != nil {
			return nil, err
		}
		newTK := i.issueUserToken(app, u, token.GrantType_REFRESH)
		newTK.Domain = tk.Domain
		newTK.StartGrantType = tk.GetStartGrantType()
		newTK.SessionId = tk.SessionId
		newTK.NamespaceId = tk.NamespaceId

		revolkReq := token.NewRevolkTokenRequest(app.ClientId, app.ClientSecret)
		revolkReq.AccessToken = req.AccessToken
		revolkReq.LogoutSession = false

		if _, err := i.token.RevolkToken(ctx, revolkReq); err != nil {
			return nil, err
		}
		return newTK, nil
	case token.GrantType_ACCESS:
		validateReq := token.NewValidateTokenRequest()
		validateReq.AccessToken = req.AccessToken
		tk, err := i.token.ValidateToken(context.Background(), validateReq)
		if err != nil {
			return nil, exception.NewUnauthorized(err.Error())
		}

		u, err := i.getUser(ctx, tk.Account)
		if err != nil {
			return nil, err
		}
		newTK := i.issueUserToken(app, u, token.GrantType_ACCESS)
		newTK.Domain = tk.Domain
		newTK.AccessExpiredAt = req.AccessExpiredAt
		newTK.RefreshExpiredAt = 4 * req.AccessExpiredAt
		newTK.Description = req.Description
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
		u, err := i.syncLDAPUser(ctx, req.Username)
		if err != nil {
			return nil, err
		}
		newTK := i.issueUserToken(app, u, token.GrantType_LDAP)
		newTK.Domain = ldapConf.Domain
		return newTK, nil
	case token.GrantType_WECHAT_WORK:
		ww, err := i.wechatWork.DescribeConfig(&wechatWork.DescribeWechatWorkConf{Domain: req.Username})
		if err != nil {
			return nil, err
		}
		if req.State != ww.State {
			return nil, errors.New("State is error! ")
		}
		np := wxwork.NewAuth(ww.CorpID, ww.CorpSecret, ww.AgentID)
		userID, err := np.CheckCallBack(&wxwork.ScanCodeRequest{
			Code:    req.AuthCode,
			State:   req.State,
			AppID:   req.UserAgent,
			Service: req.Service,
		})
		if err != nil {
			return nil, err
		}
		wxUser := np.GetUserInfo(userID)
		u, err := i.syncWXWORKUser(ctx, &user.User{
			Account: wxUser.Name,
			Domain:  ww.Domain,
			Profile: &user.Profile{
				RealName: wxUser.UserID,
				NickName: wxUser.Name,
				Avatar:   wxUser.Avatar,
				Email:    wxUser.Email,
				Phone:    wxUser.Mobile,
			},
		})
		if err != nil {
			return nil, err
		}
		newTK := i.issueUserToken(app, u, token.GrantType_WECHAT_WORK)
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

func (i *issuer) syncLDAPUser(ctx context.Context, userName string) (*user.User, error) {
	descUser := user.NewDescriptAccountRequestWithAccount(userName)
	u, err := i.user.DescribeAccount(ctx, descUser)

	if u != nil && u.Type.IsIn(types.UserType_PRIMARY, types.UserType_SUPPER) {
		return nil, exception.NewBadRequest("用户名和主账号用户名冲突, 请修改")
	}

	if err != nil {
		if exception.IsNotFoundError(err) {
			req := user.NewCreateUserRequestWithLDAPSync(userName, i.randomPass())
			req.UserType = types.UserType_SUB
			u, err = i.user.CreateAccount(ctx, req)
			if err != nil {
				return nil, err
			}
		}
		return u, err
	}

	return u, nil
}

func (i *issuer) syncWXWORKUser(ctx context.Context, reqUser *user.User) (*user.User, error) {
	descUser := user.NewDescriptAccountRequestWithAccount(reqUser.Account)
	u, err := i.user.DescribeAccount(ctx, descUser)

	if u != nil && u.Type.IsIn(types.UserType_PRIMARY, types.UserType_SUPPER) {
		return nil, exception.NewBadRequest("用户名和主账号用户名冲突, 请修改")
	}

	if err != nil {
		if exception.IsNotFoundError(err) {
			req := user.NewCreateUserRequestWithWXWORKSync(reqUser.Account, i.randomPass())
			req.UserType = types.UserType_SUB
			req.Profile = reqUser.Profile
			req.Domain = reqUser.Domain
			u, err = i.user.CreateAccount(ctx, req)
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
