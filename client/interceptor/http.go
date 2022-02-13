package interceptor

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/infraboard/keyauth/apps/micro"
	"github.com/infraboard/keyauth/apps/permission"
	"github.com/infraboard/keyauth/apps/token"
	"github.com/infraboard/keyauth/apps/user/types"
	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/common/header"
	"github.com/infraboard/mcube/exception"
	httpctx "github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	httpb "github.com/infraboard/mcube/pb/http"
	"github.com/rs/xid"
)

type PermissionCheckMode int

const (
	// PRBAC_MODE 基于策略的权限校验
	PRBAC_MODE PermissionCheckMode = 1
	// ACL_MODE 基于用户类型的权限校验
	ACL_MODE = 2
)

// NewInternalAuther 内部使用的auther
func NewHTTPAuther(c *client.Client) *HTTPAuther {
	return &HTTPAuther{
		keyauth: c,
		l:       zap.L().Named("Http Interceptor"),
		mode:    PRBAC_MODE,
	}
}

// internal todo
type HTTPAuther struct {
	l       logger.Logger
	keyauth *client.Client
	mode    PermissionCheckMode
	svr     *micro.Micro
	lock    sync.Mutex
}

func (a *HTTPAuther) SetPermissionCheckMode(m PermissionCheckMode) {
	a.mode = m
}

func (a *HTTPAuther) Auth(r *http.Request, entry httpb.Entry) (
	authInfo interface{}, err error) {
	var tk *token.Token

	// 从请求中获取access token
	acessToken := r.Header.Get(header.OAuthTokenHeader)

	if entry.AuthEnable {
		ctx := r.Context()

		// 校验身份
		tk, err = a.ValidateIdentity(ctx, acessToken)
		if err != nil {
			return nil, err
		}

		// namesapce检查
		if entry.RequiredNamespace && tk.NamespaceId == "" {
			return nil, exception.NewBadRequest("namespace required!")
		}

		// 权限检查
		if entry.PermissionEnable {
			err = a.CheckPermission(ctx, a.mode, tk, entry)
			if err != nil {
				return nil, err
			}
		}
	}

	// 设置RequestID
	if r.Header.Get(header.RequestIdHeader) == "" {
		r.Header.Set(header.RequestIdHeader, xid.New().String())
	}

	return tk, nil
}

func (a *HTTPAuther) ValidateIdentity(ctx context.Context, accessToken string) (*token.Token, error) {
	a.l.Debug("start token identity check ...")

	if accessToken == "" {
		return nil, exception.NewBadRequest("token required")
	}

	req := token.NewValidateTokenRequest()
	req.AccessToken = accessToken
	tk, err := a.keyauth.Token().ValidateToken(ctx, req)
	if err != nil {
		return nil, err
	}

	a.l.Debugf("token check ok, username: %s", tk.Account)
	return tk, nil
}

func (a *HTTPAuther) CheckPermission(ctx context.Context, mod PermissionCheckMode, tk *token.Token, e httpb.Entry) error {
	if tk == nil {
		return exception.NewUnauthorized("validate permission need token")
	}

	// 如果是超级管理员不做权限校验, 直接放行
	if tk.UserType.IsIn(types.UserType_SUPPER) {
		a.l.Debugf("[%s] supper admin skip permission check!", tk.Account)
		return nil
	}

	switch a.mode {
	case ACL_MODE:
		return a.ValidatePermissionByACL(ctx, tk, e)
	case PRBAC_MODE:
		return a.ValidatePermissionByPRBAC(ctx, tk, e)
	default:
		return fmt.Errorf("only support acl and prbac")
	}

}

func (a *HTTPAuther) ValidatePermissionByACL(ctx context.Context, tk *token.Token, e httpb.Entry) error {
	// 检查是否是允许的类型
	if len(e.Allow) > 0 {
		a.l.Debugf("[%s] start check permission to keyauth ...", tk.Account)
		if !e.IsAllow(tk.UserType) {
			return exception.NewPermissionDeny("no permission, allow: %s, but current: %s", e.Allow, tk.UserType)
		}
		a.l.Debugf("[%s] permission check passed", tk.Account)
	}

	return nil
}

func (a *HTTPAuther) ValidatePermissionByPRBAC(ctx context.Context, tk *token.Token, e httpb.Entry) error {
	svr, err := a.GetClientService(ctx)
	if err != nil {
		return err
	}

	req := permission.NewCheckPermissionRequest()
	req.Account = tk.Account
	req.NamespaceId = tk.NamespaceId
	req.ServiceId = svr.Id
	req.Path = e.UniquePath()
	_, err = a.keyauth.Permission().CheckPermission(ctx, req)
	if err != nil {
		return exception.NewPermissionDeny(err.Error())
	}
	a.l.Debugf("[%s] permission check passed", tk.Account)
	return nil
}

func (a *HTTPAuther) GetClientService(ctx context.Context) (*micro.Micro, error) {
	if a.svr != nil {
		return a.svr, nil
	}
	a.lock.Lock()
	defer a.lock.Unlock()

	req := micro.NewDescribeServiceRequestWithClientID(a.keyauth.GetClientID())
	ins, err := a.keyauth.Micro().DescribeService(ctx, req)
	if err != nil {
		return nil, err
	}
	a.svr = ins
	return ins, nil
}

func (a *HTTPAuther) ResponseHook(w http.ResponseWriter, r *http.Request, entry httpb.Entry) {
	ctx := httpctx.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	// 审计日志
	od := newOperateEventData(&entry, tk)
	hd := newEventHeaderFromHTTP(r)
	if entry.AuditLog {
		if err := SendOperateEvent(r.URL.String(), nil, hd, od); err != nil {
			a.l.Errorf("send operate event error, %s", err)
		}
	}
}

// SetLogger todo
func (a *HTTPAuther) SetLogger(l logger.Logger) {
	a.l = l
}
