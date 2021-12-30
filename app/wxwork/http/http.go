package http

import (
	"github.com/infraboard/keyauth/app/wxwork"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

)

var (
	api = &handler{}
)

type handler struct {
	service wxwork.WechatWork
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	wx := router.ResourceRouter("wechat")
	wx.BasePath("settings/wechat")
	wx.Permission(true)
	wx.Handle("GET", "/", h.GetConf).AddLabel(label.Get)
	wx.Handle("GET", "/list", h.ListConf).AddLabel(label.List)
	wx.Handle("POST", "/", h.CreateConf).AddLabel(label.Create)
	wx.Handle("DELETE", "/", h.DestroyConfig).AddLabel(label.Create)
	wx.Handle("GET", "/callback", h.WechatWorkCheck).AddLabel(label.Create).DisableAuth()
}

func (h *handler) Config() error {
	h.service = app.GetInternalApp(wxwork.AppName).(wxwork.WechatWork)
	return nil
}

func (h *handler) Name() string {
	return wxwork.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
