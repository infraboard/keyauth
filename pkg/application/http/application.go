package http

import (
	"net/http"

	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
)

// CreateApplication 创建主账号
func (h *handler) CreateUserApplication(w http.ResponseWriter, r *http.Request) {
	req := application.NewCreateApplicatonRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.CreateUserApplication("xxx", req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

// DestroyPrimaryAccount 注销账号
func (h *handler) DestroyApplication(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	if err := h.service.DeleteApplication(rctx.PS.ByName("id")); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}
