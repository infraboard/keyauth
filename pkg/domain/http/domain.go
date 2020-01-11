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

	dommains, total, err := h.service.QueryDomain(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	data := response.PageData{
		PageRequest: page,
		TotalCount:  uint(total),
		List:        dommains,
	}
	response.Success(w, data)
	return
}

func (h *handler) GetDomain(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := domain.NewDescriptDomainRequest()
	req.ID = rctx.PS.ByName("id")
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

func (h *handler) UpdateDomain(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	// 查找出原来的domain
	req := domain.NewDescriptDomainRequest()
	req.ID = rctx.PS.ByName("id")
	d, err := h.service.DescriptionDomain(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, d.CreateDomainRequst); err != nil {
		response.Failed(w, err)
		return
	}

	if err := h.service.UpdateDomain(d); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) DeleteDomain(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	if err := h.service.DeleteDomain(rctx.PS.ByName("id")); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}
