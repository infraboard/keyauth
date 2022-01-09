package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/apps/micro"
)

func (h *handler) QueryService(w http.ResponseWriter, r *http.Request) {
	page := request.NewPageRequestFromHTTP(r)
	req := micro.NewQueryMicroRequest(page)
	req.Type = micro.Type_CUSTOM

	apps, err := h.service.QueryService(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
	return
}

func (h *handler) CreateService(w http.ResponseWriter, r *http.Request) {
	req := micro.NewCreateMicroRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.CreateService(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) GetService(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := micro.NewDescribeServiceRequest()
	req.Id = rctx.PS.ByName("id")

	d, err := h.service.DescribeService(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

// DestroyService 销毁服务
func (h *handler) DestroyService(w http.ResponseWriter, r *http.Request) {

	rctx := context.GetContext(r)
	req := micro.NewDeleteMicroRequestWithID(rctx.PS.ByName("id"))

	_, err := h.service.DeleteService(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}

func (h *handler) RefreshServiceClientSecret(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := micro.NewDescribeServiceRequest()
	req.Id = rctx.PS.ByName("id")

	d, err := h.service.RefreshServiceClientSecret(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}
