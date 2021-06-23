package client

import (
	"net/http"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/grpc/gcontext"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	httpb "github.com/infraboard/mcube/pb/http"
)

// NewInternalAuther 内部使用的auther
func NewHTTPAuther(c *Client) router.Auther {
	return &HTTPAuther{
		keyauth: c,
	}
}

// internal todo
type HTTPAuther struct {
	l       logger.Logger
	keyauth *Client
}

func (a *HTTPAuther) Auth(r *http.Request, entry httpb.Entry) (
	authInfo interface{}, err error) {
	ctx, err := gcontext.NewGrpcInCtxFromHTTPRequest(r)
	if err != nil {
		return nil, err
	}

	// 获取需要校验的access token(用户的身份凭证)
	accessToken := r.Header.Get(pkg.OauthTokenHeader)
	if accessToken == "" {
		return nil, exception.NewUnauthorized("auth header: %s required", pkg.OauthTokenHeader)
	}

	engine := newEntryEngine(a.keyauth, &entry, a.log())
	engine.UseUniPath()

	// 校验用户权限是否合法
	ctx.Context()
	tk, err := engine.ValidatePermission(ctx)
	if err != nil {
		return nil, err
	}

	// 审计日志
	od := newOperateEventData(&entry, tk)
	hd := newEventHeaderFromCtx(ctx)
	if entry.AuditLog {
		defer engine.SendOperateEvent(r.URL, nil, hd, od)
	}

	return tk, nil
}

func (a *HTTPAuther) log() logger.Logger {
	if a == nil {
		a.l = zap.L().Named("GRPC Auther")
	}

	return a.l
}
