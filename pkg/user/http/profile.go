package http

import (
	"net/http"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) QueryProfile(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := user.NewDescriptAccountRequest()
	ins, err := h.service.DescribeAccount(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
	return
}

func (h *handler) PutProfile(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	tk, err := ctx.GetToken()
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := user.NewPutAccountRequest()
	req.Account = tk.Account

	if err := request.GetDataFromRequest(r, req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateAccountProfile(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
	return
}

func (h *handler) PatchProfile(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	tk, err := ctx.GetToken()
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := user.NewPatchAccountRequest()
	req.Account = tk.Account

	if err := request.GetDataFromRequest(r, req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateAccountProfile(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
	return
}

func (h *handler) QueryDomain(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	tk, err := ctx.GetToken()
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := domain.NewDescribeDomainRequest()
	req.Name = tk.Domain

	ins, err := h.domain.DescribeDomain(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) UpdateDomainInfo(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	tk, err := ctx.GetToken()
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 查找出原来的domain
	req := domain.NewPatchDomainRequest()
	req.Name = tk.Domain

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.domain.UpdateDomain(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) UpdateDomainSecurity(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	tk, err := ctx.GetToken()
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 查找出原来的domain
	req := domain.NewPutDomainSecurityRequest()
	req.Name = tk.Domain

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.SecuritySetting); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.domain.UpdateDomainSecurity(ctx.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}
