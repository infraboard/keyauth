package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
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
	if pkg.Storage == nil {
		return errors.New("denpence Storage service is nil")
	}

	h.service = pkg.Storage
	return nil
}

func init() {
	pkg.RegistryHTTPV1("buckets", api)
}
