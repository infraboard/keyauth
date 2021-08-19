package http

import (
	"fmt"
	"net/http"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (h *handler) QueryProfile(w http.ResponseWriter, r *http.Request) {
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

	req := user.NewDescriptAccountRequest()
	req.Account = tk.Account

	var header, trailer metadata.MD
	ins, err := h.service.DescribeAccount(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
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

	var header, trailer metadata.MD
	ins, err := h.service.UpdateAccountProfile(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
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

	var header, trailer metadata.MD
	ins, err := h.service.UpdateAccountProfile(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
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

	fmt.Println(tk)

	req := domain.NewDescribeDomainRequest()
	req.Name = tk.Domain

	var header, trailer metadata.MD
	ins, err := h.domain.DescribeDomain(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
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

	var header, trailer metadata.MD
	ins, err := h.domain.UpdateDomain(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
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

	var header, trailer metadata.MD
	ins, err := h.domain.UpdateDomainSecurity(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	response.Success(w, ins)
	return
}
