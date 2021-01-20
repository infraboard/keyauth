package http

import (
	"net/http"

	httpcontext "github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
)

func (h *handler) ListDomains(w http.ResponseWriter, r *http.Request) {
	page := request.NewPageRequestFromHTTP(r)
	req := domain.NewQueryDomainRequest(page)
	ctx, err := pkg.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	dommains, err := h.service.QueryDomain(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, dommains)
	return
}

func (h *handler) GetDomain(w http.ResponseWriter, r *http.Request) {
	rctx := httpcontext.GetContext(r)

	req := domain.NewDescribeDomainRequest()
	req.Name = rctx.PS.ByName("name")

	ctx, err := pkg.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.DescribeDomain(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) CreateDomain(w http.ResponseWriter, r *http.Request) {
	req := domain.NewCreateDomainRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	ctx, err := pkg.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.CreateDomain(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) PutDomain(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}
	rctx := httpcontext.GetContext(r)

	// 查找出原来的domain
	req := domain.NewPutDomainRequest()
	req.Data.Name = rctx.PS.ByName("name")

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateDomain(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) PatchDomain(w http.ResponseWriter, r *http.Request) {
	rctx := httpcontext.GetContext(r)
	ctx, err := pkg.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 查找出原来的domain
	req := domain.NewPatchDomainRequest()
	req.Data.Name = rctx.PS.ByName("name")

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateDomain(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) DeleteDomain(w http.ResponseWriter, r *http.Request) {
	rctx := httpcontext.GetContext(r)
	ctx, err := pkg.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := domain.NewDeleteDomainRequestByName(rctx.PS.ByName("name"))
	if _, err := h.service.DeleteDomain(ctx, req); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}

func (h *handler) UpdateDomainSecurity(w http.ResponseWriter, r *http.Request) {
	rctx := httpcontext.GetContext(r)
	ctx, err := pkg.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 查找出原来的domain
	req := domain.NewPutDomainSecurityRequest()
	req.Name = rctx.PS.ByName("name")
	req.SecuritySetting = domain.NewDefaultSecuritySetting()

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.SecuritySetting); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateDomainSecurity(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}
