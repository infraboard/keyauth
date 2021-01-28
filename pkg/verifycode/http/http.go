package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

var (
	api = &handler{}
)

type handler struct {
	service verifycode.VerifyCodeServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("verify_code")
	r.BasePath("verify_code")
	r.Handle("POST", "/pass", h.IssueCodeByPass).DisableAuth()
	r.Handle("POST", "/token", h.IssueCodeByToken).EnableAuth()
}

func (h *handler) Config() error {
	if pkg.VerifyCode == nil {
		return errors.New("denpence verify code service is nil")
	}

	h.service = pkg.VerifyCode
	return nil
}

func init() {
	pkg.RegistryHTTPV1("verify_code", api)
}
