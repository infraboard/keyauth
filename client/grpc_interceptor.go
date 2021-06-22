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

	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/token"
)

type PathEntryHandleFunc func(path string) *httpb.Entry

// NewGrpcKeyauthAuther todo
func NewGrpcKeyauthAuther(hf PathEntryHandleFunc, c *Client) *GrpcAuther {
	return &GrpcAuther{
		hf:       hf,
		c:        c,
		sessions: map[string]*token.Token{},
	}
}

// GrpcAuther todo
type GrpcAuther struct {
	hf PathEntryHandleFunc
	l  logger.Logger
	c  *Client

	sessions map[string]*token.Token
}

// AuthUnaryServerInterceptor returns a new unary server interceptor for auth.
func (a *GrpcAuther) AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return a.auth
}

// SetLogger todo
func (a *GrpcAuther) SetLogger(l logger.Logger) {
	a.l = l
}

func (a *GrpcAuther) GetToken(requestID string) *token.Token {
	if v, ok := a.sessions[requestID]; ok {
		return v
	}

	return nil
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

	// 校验用户权限是否合法
	tk, err := engine.ValidatePermission(rctx)
	if err != nil {
		return nil, err
	}

	// 审计日志
	od := newOperateEventData(entry, tk)
	hd := newEventHeaderFromCtx(rctx)
	if entry.AuditLog {
		defer engine.SendOperateEvent(req, resp, hd, od)
	}

	// 保存会话
	rid := xid.New().String()
	rctx.SetRequestID(rid)
	a.addSession(rid, tk)
	defer a.delSession(rid)

	return handler(rctx.Context(), req)
}

func (a *GrpcAuther) validateServiceCredential(ctx *gcontext.GrpcInCtx) error {
	clientID := ctx.GetClientID()
	clientSecret := ctx.GetClientSecret()

	if clientID == "" && clientSecret == "" {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	vsReq := micro.NewValidateClientCredentialRequest(clientID, clientSecret)
	_, err := a.c.Micro().ValidateClientCredential(context.Background(), vsReq)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "service auth error, %s", err)
	}

	return nil
}

func (a *GrpcAuther) addSession(requestID string, tk *token.Token) {
	a.sessions[requestID] = tk
}

func (a GrpcAuther) delSession(requestID string) {
	delete(a.sessions, requestID)
}

func (a *GrpcAuther) log() logger.Logger {
	if a == nil {
		a.l = zap.L().Named("GRPC Auther")
	}

	return a.l
}
