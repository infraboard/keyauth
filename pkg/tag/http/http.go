package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/tag"
)

var (
	api = &handler{}
)

type handler struct {
	service tag.TagServiceClient
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
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Tag()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("tag", api)
}
