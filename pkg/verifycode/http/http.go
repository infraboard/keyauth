package http

import (
	"errors"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

var (
	api = &handler{}
)

type handler struct {
	service verifycode.VerifyCodeServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("verify_code")

	r.BasePath("verify_code")
	r.Handle("POST", "/pass", h.IssueCodeByPass)
	r.Handle("POST", "/token", h.IssueCodeByToken)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Verifycode()
	return nil
}

func (h *handler) Name() string {
	return verifycode.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
