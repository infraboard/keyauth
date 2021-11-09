package http

import (
	"fmt"
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg/tag"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
)

// CreateApplication 创建自定义角色
func (h *handler) CreateTag(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := tag.NewCreateTagRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	switch req.ScopeType {
	case tag.ScopeType_GLOBAL:
		if !tk.UserType.IsIn(types.UserType_SUPPER) {
			response.Failed(w, fmt.Errorf("only supper account can create global tag"))
			return
		}
	case tag.ScopeType_DOMAIN:
		if !tk.UserType.IsIn(types.UserType_SUPPER, types.UserType_DOMAIN_ADMIN) {
			response.Failed(w, fmt.Errorf("only domain account can create domain tag"))
			return
		}
	}

	d, err := h.service.CreateTag(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) QueryTagKey(w http.ResponseWriter, r *http.Request) {
	req := tag.NewQueryTageKeyRequestFromHTTP(r)

	apps, err := h.service.QueryTagKey(
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

func (h *handler) QueryTagValue(w http.ResponseWriter, r *http.Request) {

	rctx := context.GetContext(r)
	req := tag.NewQueryTageValueRequestFromHTTP(r)
	req.TagId = rctx.PS.ByName("id")

	apps, err := h.service.QueryTagValue(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
	return
}

func (h *handler) DescribeTag(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	pid := rctx.PS.ByName("id")

	req := tag.NewDescribeTagRequestWithID(pid)

	ins, err := h.service.DescribeTag(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) DeleteTag(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	req := tag.NewDeleteTagRequestWithID(rctx.PS.ByName("id"))

	_, err := h.service.DeleteTag(
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
