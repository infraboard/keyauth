package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/token"
)

func (h *handler) QueryService(w http.ResponseWriter, r *http.Request) {
	page := request.NewPageRequestFromHTTP(r)
	req := micro.NewQueryMicroRequest(page)

	apps, err := h.service.QueryService(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
	return
}

func (h *handler) CreateService(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := micro.NewCreateMicroRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	req.WithToken(tk)
	d, err := h.service.CreateService(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) GetService(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := micro.NewDescribeServiceRequest()
	req.ID = rctx.PS.ByName("id")

	d, err := h.service.DescribeService(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

// DestroyService 销毁服务
func (h *handler) DestroyService(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)
	req := micro.NewDeleteMicroRequestWithID(rctx.PS.ByName("id"))
	req.WithToken(tk)
	if err := h.service.DeleteService(req); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}

func (h *handler) GetServiceToken(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := micro.NewDescribeServiceRequest()
	req.ID = rctx.PS.ByName("id")

	d, err := h.service.DescribeService(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	tk, err := h.token.DescribeToken(token.NewDescribeTokenRequestWithAccessToken(d.AccessToken))
	tk.Desensitize()
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, tk)
}

func (h *handler) RefreshServiceToken(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := micro.NewDescribeServiceRequest()
	req.ID = rctx.PS.ByName("id")

	d, err := h.service.RefreshServiceToken(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}
