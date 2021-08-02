package client

import (
	"net/http"
	"reflect"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/grpc/gcontext"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	httpb "github.com/infraboard/mcube/pb/http"
	"github.com/rs/xid"
)

// NewInternalAuther 内部使用的auther
func NewHTTPAuther(c *Client) *HTTPAuther {
	return &HTTPAuther{
		keyauth: c,
		l:       zap.L().Named("Http Interceptor"),
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

	engine := newEntryEngine(a.keyauth, &entry, a.log())
	engine.UseUniPath()

	// 校验身份
	tk, err := engine.ValidateIdentity(ctx)
	if err != nil {
		return nil, err
	}

	// 校验权限
	if err := engine.ValidatePermission(tk, ctx); err != nil {
		return nil, err
	}

	// 设置RequestID
	if r.Header.Get(gcontext.RequestID) == "" {
		r.Header.Set(gcontext.RequestID, xid.New().String())
	}

	return tk, nil
}

func (a *HTTPAuther) ResponseHook(w http.ResponseWriter, r *http.Request, entry httpb.Entry) {
	ctx, err := gcontext.NewGrpcInCtxFromHTTPRequest(r)
	if err != nil {
		a.log().Errorf("reponse hook NewGrpcInCtxFromHTTPRequest error, %s", err)
		return
	}

	tk, ok := context.GetContext(r).AuthInfo.(*token.Token)
	if !ok {
		a.log().Errorf("context AuthInfo is not *token.Token, is %s", reflect.TypeOf(context.GetContext(r).AuthInfo))
		return
	}

	// 审计日志
	od := newOperateEventData(&entry, tk)
	hd := newEventHeaderFromCtx(ctx)
	if entry.AuditLog {
		if err := SendOperateEvent(r.URL.String(), nil, hd, od); err != nil {
			a.log().Errorf("send operate event error, %s", err)
		}
	}
}

func (a *HTTPAuther) log() logger.Logger {
	if a == nil {
		a.l = zap.L().Named("HTTP Auther")
	}

	return a.l
}

// SetLogger todo
func (a *HTTPAuther) SetLogger(l logger.Logger) {
	a.l = l
}
