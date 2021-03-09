package http

import (
	"context"
	"net/http"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (h *handler) IssueCodeByPass(w http.ResponseWriter, r *http.Request) {
	req := verifycode.NewIssueCodeRequestByPass()
	// 从Header中获取client凭证, 如果有
	req.ClientId, req.ClientSecret, _ = r.BasicAuth()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	req.IssueType = verifycode.IssueType_PASS

	var header, trailer metadata.MD
	code, err := h.service.IssueCode(
		context.Background(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	response.Success(w, code)
	return
}

func (h *handler) IssueCodeByToken(w http.ResponseWriter, r *http.Request) {
	ctx, err := pkg.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := verifycode.NewIssueCodeRequestByToken()

	var header, trailer metadata.MD
	code, err := h.service.IssueCode(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, pkg.NewExceptionFromTrailer(trailer, err))
		return
	}

	response.Success(w, code)
	return
}
