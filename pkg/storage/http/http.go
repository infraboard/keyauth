package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg/storage"
)

var (
	api = &handler{}
)

type handler struct {
	service storage.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	geoipRouter := router.ResourceRouter("buckets")
	geoipRouter.BasePath("buckets")
	geoipRouter.Permission(true)
	geoipRouter.Handle("POST", "/:name/objects", h.UploadGEOIPDBFile)
}

func (h *handler) Config() error {
	h.service = app.GetInternalApp(storage.AppName).(storage.Service)
	return nil
}

func (h *handler) Name() string {
	return storage.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
