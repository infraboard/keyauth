package http

import (
	"net/http"

	"github.com/infraboard/keyauth/app/domain"
	"github.com/infraboard/keyauth/app/token"
	"github.com/infraboard/keyauth/app/user"
	"github.com/infraboard/keyauth/app/user/types"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) QueryProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := user.NewDescriptAccountRequest()
	req.Account = tk.Account

	ins, err := h.service.DescribeAccount(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
	return
}

func (h *handler) PutProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := user.NewPutAccountRequest()
	req.Account = tk.Account

	if err := request.GetDataFromRequest(r, req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateAccountProfile(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
}

func (h *handler) PatchProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := user.NewPatchAccountRequest()
	req.Account = tk.Account

	if err := request.GetDataFromRequest(r, req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	// 更新部门
	if req.DepartmentId != "" {
		if !tk.UserType.IsIn(types.UserType_SUPPER, types.UserType_INTERNAL, types.UserType_DOMAIN_ADMIN, types.UserType_ORG_ADMIN) {
			response.Failed(w, exception.NewBadRequest("组织管理员才能直接修改用户部门"))
			return
		}
	}

	ins, err := h.service.UpdateAccountProfile(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
}

func (h *handler) QueryDomain(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	req := domain.NewDescribeDomainRequest()
	req.Name = tk.Domain

	ins, err := h.domain.DescribeDomain(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) UpdateDomainInfo(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	// 查找出原来的domain
	req := domain.NewPatchDomainRequest()
	req.Name = tk.Domain

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.domain.UpdateDomain(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) UpdateDomainSecurity(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	tk := ctx.AuthInfo.(*token.Token)

	// 查找出原来的domain
	req := domain.NewPutDomainSecurityRequest()
	req.Name = tk.Domain

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.SecuritySetting); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.domain.UpdateDomainSecurity(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
