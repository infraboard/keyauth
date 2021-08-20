package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/department"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
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

	req := department.NewQueryDepartmentRequestFromHTTP(r)
	req.Domain = tk.Domain

	var header, trailer metadata.MD
	apps, err := h.service.QueryDepartment(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	response.Success(w, apps)
}

// Create 创建主账号
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
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

	req := department.NewCreateDepartmentRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateOwner(tk)

	var header, trailer metadata.MD
	ins, err := h.service.CreateDepartment(
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
}

// Create 创建主账号
func (h *handler) Put(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := department.NewPutUpdateDepartmentRequest(rctx.PS.ByName("id"))
	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}

	var header, trailer metadata.MD
	ins, err := h.service.UpdateDepartment(
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
}

// Create 创建主账号
func (h *handler) Patch(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := department.NewPatchUpdateDepartmentRequest(rctx.PS.ByName("id"))
	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}

	var header, trailer metadata.MD
	ins, err := h.service.UpdateDepartment(
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
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
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

	rctx := context.GetContext(r)
	qs := r.URL.Query()

	req := department.NewDescribeDepartmentRequest()
	req.Domain = tk.Domain

	req.Id = rctx.PS.ByName("id")
	req.WithSubCount = qs.Get("with_sub_count") == "true"
	req.WithUserCount = qs.Get("with_user_count") == "true"
	req.WithRole = qs.Get("with_role") == "true"

	var header, trailer metadata.MD
	ins, err := h.service.DescribeDepartment(
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
}

func (h *handler) GetSub(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)
	pid := rctx.PS.ByName("id")

	req := department.NewQueryDepartmentRequestFromHTTP(r)
	req.ParentId = pid

	var header, trailer metadata.MD
	ins, err := h.service.QueryDepartment(
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
}

// DestroyPrimaryAccount 注销账号
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	rctx := context.GetContext(r)

	req := department.NewDeleteDepartmentRequestWithID(rctx.PS.ByName("id"))

	var header, trailer metadata.MD
	_, err = h.service.DeleteDepartment(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	response.Success(w, "delete ok")
}
