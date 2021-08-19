package auther

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/infraboard/mcube/bus"
	"github.com/infraboard/mcube/bus/event"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/grpc/gcontext"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	httpb "github.com/infraboard/mcube/pb/http"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/keyauth/version"
)

func newEntryEngine(entry *httpb.Entry) *entryEngine {
	if pkg.Token == nil || pkg.Permission == nil || pkg.Micro == nil {
		panic("token, permission, micro required")
	}

	return &entryEngine{
		Entry:      entry,
		log:        zap.L().Named("Auther Entry"),
		uniPath:    false,
		token:      pkg.Token,
		permission: pkg.Permission,
		micro:      pkg.Micro,
	}
}

type entryEngine struct {
	*httpb.Entry

	serviceId  string
	log        logger.Logger
	uniPath    bool
	token      token.TokenServiceServer
	permission permission.PermissionServiceServer
	micro      micro.MicroServiceServer
}

func (e *entryEngine) UseUniPath() {
	e.uniPath = true
}

func (e *entryEngine) ValidateIdentity(ctx *gcontext.GrpcInCtx) (*token.Token, error) {
	e.log.Debug("start token identity check ...")

	if !e.AuthEnable {
		e.log.Debug("auth disabled skip")
		return nil, nil
	}

	// 获取需要校验的access token(用户的身份凭证)
	accessToken := ctx.GetAccessToKen()

	if accessToken == "" {
		e.log.Debugf("[%s] auth enabled, but not get access token", e.Path)
		return nil, exception.NewBadRequest("token required")
	}

	req := token.NewValidateTokenRequest()
	req.AccessToken = accessToken

	outCtx := gcontext.NewGrpcOutCtx()
	outCtx.SetAccessToken(ctx.GetAccessToKen())
	tk, err := e.token.ValidateToken(outCtx.Context(), req)
	if err != nil {
		return nil, err
	}

	e.log.Debugf("token check ok, username: %s", tk.Account)
	return tk, nil
}

func (e *entryEngine) ValidatePermission(tk *token.Token, ctx *gcontext.GrpcInCtx) error {
	if !e.AuthEnable {
		return nil
	}

	if tk == nil {
		return exception.NewUnauthorized("validate permission need token")
	}

	if !e.PermissionEnable {
		e.log.Debugf("[%s] permission disabled skip check!", tk.Account)
		return nil
	}

	if e.RequiredNamespace && tk.Namespace == "" {
		return exception.NewBadRequest("namespace required!")
	}

	// 如果是超级管理员不做权限校验, 直接放行
	if tk.UserType.IsIn(types.UserType_SUPPER) {
		e.log.Debugf("[%s] supper admin skip permission check!", tk.Account)
		return nil
	}

	// 检查是否是允许的类型
	if len(e.Allow) > 0 {
		e.log.Debugf("[%s] start check permission to keyauth ...", tk.Account)
		if !e.IsAllow(tk.UserType) {
			return exception.NewPermissionDeny("no permission, allow: %s, but current: %s", e.Allow, tk.UserType)
		}
		e.log.Debugf("[%s] permission check passed", tk.Account)
	}

	return nil
}

func SendOperateEvent(req, resp interface{}, hd *event.Header, od *event.OperateEventData) error {
	if od == nil {
		return nil
	}

	reqd, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal req for event error, %s", err)
	}

	respd, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("marshal resp for event error, %s", err)
	}

	od.Request = string(reqd)
	od.Response = string(respd)
	od.Cost = ftime.Now().Timestamp() - hd.Time
	oe, err := event.NewProtoOperateEvent(od)
	if err != nil {
		return fmt.Errorf("new operate event error, %s", err)
	}
	oe.Header = hd

	if err := bus.Pub(oe); err != nil {
		return fmt.Errorf("pub audit log error, %s", err)
	}

	return nil
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
