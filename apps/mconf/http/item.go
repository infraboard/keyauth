package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/apps/mconf"
)

func (h *handler) QueryItem(w http.ResponseWriter, r *http.Request) {
	page := request.NewPageRequestFromHTTP(r)
	req := mconf.NewQueryItemRequest(page)

	rctx := context.GetContext(r)
	req.GroupName = rctx.PS.ByName("name")

	apps, err := h.service.QueryItem(
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

func (h *handler) AddItemToGroup(w http.ResponseWriter, r *http.Request) {
	req := mconf.NewAddItemToGroupRequest()
	rctx := context.GetContext(r)
	req.GroupName = rctx.PS.ByName("name")
	if err := request.GetDataFromRequest(r, &req.Items); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.AddItemToGroup(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}
