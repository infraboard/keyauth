package client

import (
	"context"
	"strconv"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/grpc/gcontext"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	httpb "github.com/infraboard/mcube/pb/http"
	"github.com/rs/xid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/infraboard/keyauth/client/session"
	"github.com/infraboard/keyauth/pkg/micro"
)

type PathEntryHandleFunc func(path string) *httpb.Entry

// NewGrpcKeyauthAuther todo
func NewGrpcKeyauthAuther(hf PathEntryHandleFunc, c *Client, store session.Store) *GrpcAuther {
	return &GrpcAuther{
		hf:    hf,
		c:     c,
		store: store,
		l:     zap.L().Named("Grpc Interceptor"),
	}
}

// GrpcAuther todo
type GrpcAuther struct {
	hf    PathEntryHandleFunc
	l     logger.Logger
	c     *Client
	store session.Store
}

// AuthUnaryServerInterceptor returns a new unary server interceptor for auth.
func (a *GrpcAuther) AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return a.auth
}

// SetLogger todo
func (a *GrpcAuther) SetLogger(l logger.Logger) {
	a.l = l
}

// Auth impl interface
func (a *GrpcAuther) auth(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if err != nil {
			switch t := err.(type) {
			case exception.APIException:
				err = status.Errorf(codes.Code(t.ErrorCode()), t.Error())
				trailer := metadata.Pairs(
					gcontext.ResponseCodeHeader, strconv.Itoa(t.ErrorCode()),
					gcontext.ResponseReasonHeader, t.Reason(),
					gcontext.ResponseDescHeader, t.Error(),
				)
				if err := grpc.SetTrailer(ctx, trailer); err != nil {
					a.log().Errorf("send grpc trailer error, %s", err)
				}
			}
		}
	}()
	// 重上下文中获取认证信息
	rctx, err := gcontext.GetGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 校验调用的客户端凭证是否有效
	if err := a.validateServiceCredential(rctx); err != nil {
		return nil, err
	}

	entry := a.hf(info.FullMethod)
	if entry == nil {
		return nil, status.Errorf(codes.Internal, "entry not found, check is registry")
	}
	engine := newEntryEngine(a.c, entry, a.log())

	// 校验身份
	tk := a.store.LeaseToken(rctx.GetAccessToKen())
	if tk == nil {
		tk, err = engine.ValidateIdentity(rctx)
		if err != nil {
			return nil, err
		}
		a.store.SetToken(tk)
	}
	defer a.store.ReturnToken(tk)

	// 校验权限
	if err := engine.ValidatePermission(tk, rctx); err != nil {
		return nil, err
	}

	// 保存会话
	rid := rctx.GetRequestID()
	if rid == "" {
		rid = xid.New().String()
		rctx.SetRequestID(rid)
	}

	// 审计日志
	if entry.AuditLog {
		od := newOperateEventData(entry, tk)
		hd := newEventHeaderFromCtx(rctx)
		defer func() {
			a.log().Debugf("[%s] start send operate event ...")
			err := SendOperateEvent(req, resp, hd, od)
			if err != nil {
				a.log().Warnf("[%s] send operate event failed, %s", err)
				return
			}
			a.log().Warnf("[%s] send operate event ok", err)
		}()
	}

	return handler(rctx.Context(), req)
}

func (a *GrpcAuther) validateServiceCredential(ctx *gcontext.GrpcInCtx) error {
	clientID := ctx.GetClientID()
	clientSecret := ctx.GetClientSecret()
	a.log().Debugf("start check client[%s] credential ...", clientID)

	if clientID == "" && clientSecret == "" {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	vsReq := micro.NewValidateClientCredentialRequest(clientID, clientSecret)
	_, err := a.c.Micro().ValidateClientCredential(context.Background(), vsReq)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "service auth error, %s", err)
	}

	a.log().Debugf("check client[%s] credential ok", clientID)
	return nil
}

func (a *GrpcAuther) log() logger.Logger {
	if a == nil {
		a.l = zap.L().Named("GRPC Auther")
	}

	return a.l
}
