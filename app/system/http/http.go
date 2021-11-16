package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/app/system"
	"github.com/infraboard/keyauth/app/user/types"
)

var (
	api = &handler{}
)

type handler struct {
	service system.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("system_config")
	r.Allow(types.UserType_SUPPER)
	r.BasePath("system_config")
	r.Handle("GET", "/", h.GetSystemConfig)
	r.Handle("POST", "/email/", h.SettingEmail)
	r.Handle("POST", "/email/test", h.TestEmailSend)
	r.Handle("POST", "/sms/", h.SettingSMS)
	r.Handle("POST", "/sms/test", h.TestSMSSend)
	r.Handle("POST", "/verify_code/", h.SettingVerifyCode)
}

func (h *handler) Config() error {
	h.service = app.GetInternalApp(system.AppName).(system.Service)
	return nil
}

func (h *handler) Name() string {
	return system.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
