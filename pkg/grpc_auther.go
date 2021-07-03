package pkg

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/infraboard/mcube/bus"
	"github.com/infraboard/mcube/bus/event"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"
	"github.com/infraboard/mcube/types/ftime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/keyauth/version"
)

var (
	interceptor = newGrpcAuther()
)

// 检测是不是owner请求
type OwnerChecker interface {
	CheckOwner(account string) bool
}

// AuthUnaryServerInterceptor returns a new unary server interceptor for auth.
func AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return interceptor.Auth
}

func newGrpcAuther() *grpcAuther {
	return &grpcAuther{}
}

func newEventHeaderFromCtx(ctx *GrpcInCtx) *event.Header {
	hd := event.NewHeader()
	hd.IpAddress = ctx.GetRemoteIP()
	hd.UserAgent = ctx.GetUserAgent()
	hd.RequestId = ctx.GetRequestID()
	hd.Source = version.ServiceName
	hd.Meta["host"], _ = os.Hostname()
	return hd
}

// internal todo
type grpcAuther struct {
	l logger.Logger
}

func (a *grpcAuther) Auth(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// 转换异常类型
	defer func() {
		if err != nil {
			switch t := err.(type) {
			case exception.APIException:
				err = status.Errorf(codes.Code(t.ErrorCode()), t.Error())
				trailer := metadata.Pairs(
					ResponseCodeHeader, strconv.Itoa(t.ErrorCode()),
					ResponseReasonHeader, t.Reason(),
					ResponseDescHeader, t.Error(),
				)
				if err := grpc.SetTrailer(ctx, trailer); err != nil {
					a.log().Errorf("send grpc trailer error, %s", err)
				}
			}
		}
	}()

	// 重上下文中获取认证信息
	rctx, err := GetGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 校验调用的客户端凭证是否有效
	if err := a.validateServiceCredential(rctx); err != nil {
		return nil, err
	}

	entry := GetPathEntry(info.FullMethod)
	if entry == nil {
		return nil, grpc.Errorf(codes.Internal, "entry gprc path: %s not found, check is registry", info.FullMethod)
	}

	// 审计日志
	od := newOperateEventData(entry)
	hd := newEventHeaderFromCtx(rctx)
	if entry.AuditLog {
		defer a.sendOperateEvent(req, resp, hd, od)
	}

	// 权限校验
	if entry.AuthEnable {
		tk, err := a.checkToken(rctx)
		if err != nil {
			return nil, err
		}

		// 补充审计的用户信息
		od.Account = tk.Account
		od.UserDomain = tk.Domain
		od.Session = tk.SessionId
		od.UserType = tk.UserType.String()

		// 权限校验
		if err := a.validatePermission(tk, entry, req); err != nil {
			return nil, err
		}
	}

	rctx.ClearInternl()
	resp, err = handler(rctx.ClearInternl().Context(), req)
	return resp, err
}

func (a *grpcAuther) sendOperateEvent(req, resp interface{}, hd *event.Header, od *event.OperateEventData) {
	if od == nil {
		return
	}

	reqd, err := json.Marshal(req)
	if err != nil {
		a.log().Warnf("marshal req for event error, %s", err)
	}

	respd, err := json.Marshal(resp)
	if err != nil {
		a.log().Warnf("marshal resp for event error, %s", err)
	}

	od.Request = string(reqd)
	od.Response = string(respd)
	od.Cost = ftime.Now().Timestamp() - hd.Time
	oe, err := event.NewProtoOperateEvent(od)
	if err != nil {
		a.log().Errorf("new operate event error, %s", err)
	}
	oe.Header = hd

	if err := bus.Pub(oe); err != nil {
		a.log().Warnf("pub audit log error, %s", err)
	}
}

func (a *grpcAuther) validateServiceCredential(ctx *GrpcInCtx) error {
	clientID := ctx.GetClientID()
	clientSecret := ctx.GetClientSecret()

	if clientID == "" && clientSecret == "" {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret is \"\"")
	}

	if Micro == nil {
		return status.Errorf(codes.Internal, "micro service is not initial")
	}

	vsReq := micro.NewValidateClientCredentialRequest(clientID, clientSecret)
	_, err := Micro.ValidateClientCredential(context.Background(), vsReq)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "service auth error, %s", err)
	}

	return nil
}

// 获取需要校验的access token(用户的身份凭证)
func (a *grpcAuther) checkToken(ctx *GrpcInCtx) (*token.Token, error) {
	accessToken := ctx.GetAccessToKen()
	req := token.NewValidateTokenRequest()
	if accessToken == "" {
		return nil, status.Errorf(codes.Unauthenticated, "access_token meta required")
	}
	req.AccessToken = accessToken
	return Token.ValidateToken(context.Background(), req)
}

func (a *grpcAuther) validatePermission(tk *token.Token, entry *http.Entry, req interface{}) error {
	// 如果是超级管理员不做权限校验, 直接放行
	if tk.UserType.IsIn(types.UserType_SUPPER, types.UserType_INTERNAL) {
		return nil
	}

	// 如果是owner 直接放行
	if v, ok := req.(OwnerChecker); ok {
		if v.CheckOwner(tk.Account) {
			return nil
		}
	}

	// 如果是允许的用户类型 直接放行
	if v, ok := entry.Labels["allow"]; ok {
		if v != "*" {
			a.log().Debugf("allows: %s", v)
			if !tk.UserType.IsIn(transferToUserType(v)...) {
				return status.Errorf(codes.PermissionDenied, "access grpc path %s permission deny, allow types: %s", entry.Path, v)
			}
		}
		return nil
	}

	return nil
}

func (a *grpcAuther) endpointHashID(entry *http.Entry) string {
	return endpoint.GenHashID(version.ServiceName, entry.Path)
}

func (a *grpcAuther) log() logger.Logger {
	if a.l == nil {
		a.l = zap.L().Named("GRPC Auther")
	}

	return a.l
}

func transferToUserType(allows string) []types.UserType {
	set := []types.UserType{}
	for _, t := range strings.Split(allows, ",") {
		var ut types.UserType
		switch t {
		case "sub":
			ut = types.UserType_SUB
		case "primary":
			ut = types.UserType_PRIMARY
		case "super":
			ut = types.UserType_SUPPER
		case "internal":
			ut = types.UserType_INTERNAL
		case "domain_admin":
			ut = types.UserType_DOMAIN_ADMIN
		case "org_admin":
			ut = types.UserType_ORG_ADMIN
		case "perm_admin":
			ut = types.UserType_PERM_ADMIN
		case "audit_admin":
			ut = types.UserType_AUDIT_ADMIN
		}
		set = append(set, ut)
	}

	return set
}

func newOperateEventData(entry *http.Entry) *event.OperateEventData {
	return &event.OperateEventData{
		Action:       entry.GetLableValue("action"),
		FeaturePath:  entry.Path,
		ResourceType: entry.Resource,
		ServiceName:  version.ServiceName,
	}
}
