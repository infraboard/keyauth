package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg/tag"
)

var (
	api = &handler{}
)

type handler struct {
	service tag.TagServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("tag")
	r.BasePath("tags")
	r.Handle("POST", "/", h.CreateTag)
	r.Handle("GET", "/:id", h.DescribeTag)
	r.Handle("DELETE", "/:id", h.DeleteTag)
	r.Handle("GET", "/", h.QueryTagKey)
	r.Handle("GET", "/:id/values", h.QueryTagValue)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(tag.AppName).(tag.TagServiceServer)
	return nil
}

func (h *handler) Name() string {
	return tag.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
