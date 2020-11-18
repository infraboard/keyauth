package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg/domain"
)

func (h *handler) ListDomains(w http.ResponseWriter, r *http.Request) {
	page := request.NewPageRequestFromHTTP(r)
	req := domain.NewQueryDomainRequest(page)

	dommains, err := h.service.QueryDomain(req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, dommains)
	return
}

func (h *handler) GetDomain(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := domain.NewDescriptDomainRequest()
	req.Name = rctx.PS.ByName("name")
	d, err := h.service.DescriptionDomain(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) CreateDomain(w http.ResponseWriter, r *http.Request) {
	req := domain.NewCreateDomainRequst()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.CreateDomain("xxx", req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) PutDomain(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	// 查找出原来的domain
	req := domain.NewPutDomainRequest()
	req.Name = rctx.PS.ByName("name")

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.CreateDomainRequst); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateDomain(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) PatchDomain(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	// 查找出原来的domain
	req := domain.NewPatchDomainRequest()
	req.Name = rctx.PS.ByName("name")

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.CreateDomainRequst); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateDomain(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) DeleteDomain(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	if err := h.service.DeleteDomain(rctx.PS.ByName("name")); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}

func (h *handler) UpdateDomainSecurity(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	// 查找出原来的domain
	req := domain.NewPutDomainSecurityRequest()
	req.Name = rctx.PS.ByName("name")
	req.SecuritySetting = domain.NewDefaultSecuritySetting()

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.SecuritySetting); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateDomainSecurity(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}
