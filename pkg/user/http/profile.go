package http

import (
	"context"
	"net/http"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) QueryProfile(w http.ResponseWriter, r *http.Request) {
	ctx, err := session.GetTokenCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := user.NewDescriptAccountRequest()
	ins, err := h.service.DescribeAccount(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
	return
}

func (h *handler) PutProfile(w http.ResponseWriter, r *http.Request) {
	tk, err := session.GetTokenFromHTTPRequest(r)
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

	ctx := session.WithTokenContext(context.Background(), tk)
	ins, err := h.service.UpdateAccountProfile(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
	return
}

func (h *handler) PatchProfile(w http.ResponseWriter, r *http.Request) {
	tk, err := session.GetTokenFromHTTPRequest(r)
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

	ctx := session.WithTokenContext(context.Background(), tk)
	ins, err := h.service.UpdateAccountProfile(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
	return
}

func (h *handler) QueryDomain(w http.ResponseWriter, r *http.Request) {
	tk, err := session.GetTokenFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ctx := session.WithTokenContext(context.Background(), tk)

	req := domain.NewDescribeDomainRequest()
	req.Name = tk.Domain

	ins, err := h.domain.DescribeDomain(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) UpdateDomainInfo(w http.ResponseWriter, r *http.Request) {
	tk, err := session.GetTokenFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ctx := session.WithTokenContext(context.Background(), tk)

	// 查找出原来的domain
	req := domain.NewPatchDomainRequest()
	req.Name = tk.Domain

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.domain.UpdateDomain(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}

func (h *handler) UpdateDomainSecurity(w http.ResponseWriter, r *http.Request) {
	tk, err := session.GetTokenFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ctx := session.WithTokenContext(context.Background(), tk)

	// 查找出原来的domain
	req := domain.NewPutDomainSecurityRequest()
	req.Name = tk.Domain

	// 解析需要更新的数据
	if err := request.GetDataFromRequest(r, req.SecuritySetting); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.domain.UpdateDomainSecurity(ctx, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
	return
}
