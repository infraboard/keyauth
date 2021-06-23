package client

import (
	"encoding/json"
	"os"

	"github.com/infraboard/mcube/bus"
	"github.com/infraboard/mcube/bus/event"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/grpc/gcontext"
	"github.com/infraboard/mcube/logger"
	httpb "github.com/infraboard/mcube/pb/http"
	"github.com/infraboard/mcube/types/ftime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/keyauth/version"
)

func newEntryEngine(client *Client, entry *httpb.Entry, log logger.Logger) *entryEngine {
	return &entryEngine{
		client: client,
		Entry:  entry,
		log:    log,
	}
}

type entryEngine struct {
	*httpb.Entry
	client    *Client
	serviceId string
	log       logger.Logger
}

func (e *entryEngine) ValidatePermission(ctx *gcontext.GrpcInCtx) (*token.Token, error) {
	var (
		tk  *token.Token
		err error
	)

	outCtx := gcontext.NewGrpcOutCtx()
	outCtx.SetAccessToken(ctx.GetAccessToKen())

	// 获取需要校验的access token(用户的身份凭证)
	accessToken := ctx.GetAccessToKen()
	if accessToken != "" {
		req := token.NewValidateTokenRequest()
		if accessToken == "" {
			return nil, status.Errorf(codes.Unauthenticated, "access_token meta required")
		}
		req.AccessToken = accessToken

		tk, err = e.client.Token().ValidateToken(outCtx.Context(), req)
		if err != nil {
			return nil, err
		}
	}

	if e.RequiredNamespace && tk != nil {
		if tk.Namespace == "" {
			return nil, exception.NewBadRequest("namespace required!")
		}
	}

	if e.PermissionEnable && tk != nil {
		// 如果是超级管理员不做权限校验, 直接放行
		if tk.UserType.IsIn(types.UserType_SUPPER) {
			return tk, nil
		}

		eid, err := e.endpointHashID(outCtx, e.Path)
		if err != nil {
			return nil, err
		}

		// 权限检测
		req := permission.NewCheckPermissionRequest()
		req.EndpointId = eid
		req.NamespaceId = tk.Namespace
		perm, err := e.client.Permission().CheckPermission(outCtx.Context(), req)
		if err != nil {
			return nil, exception.NewPermissionDeny("no permission, %s", err)
		}
		tk.Scope = perm.Scope
	}

	return tk, nil
}

func (e *entryEngine) endpointHashID(ctx *gcontext.GrpcOutCtx, path string) (string, error) {
	if e.serviceId == "" {
		descReq := micro.NewDescribeServiceRequestWithClientID(e.client.GetClientID())
		svr, err := e.client.Micro().DescribeService(ctx.Context(), descReq)
		if err != nil {
			return "", err
		}
		e.serviceId = svr.Id
	}

	return endpoint.GenHashID(e.serviceId, path), nil
}

func (e *entryEngine) SendOperateEvent(req, resp interface{}, hd *event.Header, od *event.OperateEventData) {
	if od == nil {
		return
	}

	reqd, err := json.Marshal(req)
	if err != nil {
		e.log.Warnf("marshal req for event error, %s", err)
	}

	respd, err := json.Marshal(resp)
	if err != nil {
		e.log.Warnf("marshal resp for event error, %s", err)
	}

	od.Request = string(reqd)
	od.Response = string(respd)
	od.Cost = ftime.Now().Timestamp() - hd.Time
	oe, err := event.NewOperateEvent(od)
	if err != nil {
		e.log.Errorf("new operate event error, %s", err)
	}
	oe.Header = hd

	if err := bus.Pub(oe); err != nil {
		e.log.Warnf("pub audit log error, %s", err)
	}
}

func newOperateEventData(e *httpb.Entry, tk *token.Token) *event.OperateEventData {
	od := &event.OperateEventData{
		Action:       e.GetLableValue("action"),
		FeaturePath:  e.Path,
		ResourceType: e.Resource,
		ServiceName:  version.ServiceName,
	}

	if tk != nil {
		// 补充审计的用户信息
		od.Account = tk.Account
		od.UserDomain = tk.Domain
		od.Session = tk.SessionId
		od.UserType = tk.UserType.String()
	}
	return od
}

func newEventHeaderFromCtx(ctx *gcontext.GrpcInCtx) *event.Header {
	hd := event.NewHeader()
	// hd.IpAddress = ctx.GetRemoteIP()
	// hd.UserAgent = ctx.GetUserAgent()
	hd.RequestId = ctx.GetRequestID()
	hd.Source = version.ServiceName
	hd.Meta["host"], _ = os.Hostname()
	return hd
}
