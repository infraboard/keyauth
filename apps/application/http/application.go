package http

import (
	"net/http"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/apps/application"
	"github.com/infraboard/keyauth/apps/token"
)

func (h *handler) QueryApplication(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	page := request.NewPageRequestFromHTTP(r)
	req := application.NewQueryApplicationRequest(page)
	req.Account = tk.Account

	apps, err := h.service.QueryApplication(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
}

// CreateApplication 创建主账号
func (h *handler) CreateApplication(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := application.NewCreateApplicatonRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateOwner(tk)

	d, err := h.service.CreateApplication(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) GetApplication(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := application.NewDescriptApplicationRequest()
	req.Id = ctx.PS.ByName("id")

	ins, err := h.service.DescribeApplication(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	if !ins.IsOwner(tk.Account) {
		response.Failed(w, exception.NewPermissionDeny("this application is not yours"))
		return
	}

	response.Success(w, ins)
}

// DestroyPrimaryAccount 注销账号
func (h *handler) DestroyApplication(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	descReq := application.NewDescriptApplicationRequest()
	descReq.Id = ctx.PS.ByName("id")

	ins, err := h.service.DescribeApplication(r.Context(), descReq)

	if err != nil {
		response.Failed(w, err)
		return
	}

	if !ins.IsOwner(tk.Account) {
		response.Failed(w, exception.NewPermissionDeny("this application is not yours"))
		return
	}

	req := application.NewDeleteApplicationRequestWithID(ctx.PS.ByName("id"))

	_, err = h.service.DeleteApplication(r.Context(), req)

	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
}
