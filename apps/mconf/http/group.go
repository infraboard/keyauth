package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/apps/mconf"
)

func (h *handler) QueryGroup(w http.ResponseWriter, r *http.Request) {
	page := request.NewPageRequestFromHTTP(r)
	req := mconf.NewQueryGroupRequest(page)

	apps, err := h.service.QueryGroup(
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

func (h *handler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	req := mconf.NewCreateGroupRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.CreateGroup(
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
func (h *handler) DestroyGroup(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	req := mconf.NewDeleteGroupRequestWithName(rctx.PS.ByName("name"))

	_, err := h.service.DeleteGroup(
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
