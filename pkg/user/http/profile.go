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
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := user.NewDescriptAccountRequest()
	req.ID = tk.UserID

	ins, err := h.service.DescribeAccount(req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
	return
}

func (h *handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := user.NewDescriptAccountRequest()
	req.ID = tk.UserID

	ins, err := h.service.DescribeAccount(req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.WithToken(tk)

	if err := request.GetDataFromRequest(r, ins.CreateUserRequest); err != nil {
		response.Failed(w, err)
		return
	}

	if err := h.service.UpdateAccountProfile(ins); err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
	return
}

func (h *handler) QueryDomain(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := domain.NewDescriptDomainRequest()
	req.ID = tk.DomainID

	ins, err := h.domain.DescriptionDomain(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) UpdateDomain(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 查找出原来的domain
	req := domain.NewDescriptDomainRequest()
	req.ID = tk.DomainID
	d, err := h.domain.DescriptionDomain(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, d.CreateDomainRequst); err != nil {
		response.Failed(w, err)
		return
	}

	if err := h.domain.UpdateDomain(d); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}
