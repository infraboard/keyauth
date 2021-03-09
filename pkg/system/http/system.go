package http

import (
	"net/http"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/system/notify/mail"
	"github.com/infraboard/keyauth/pkg/system/notify/sms"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (h *handler) GetSystemConfig(w http.ResponseWriter, r *http.Request) {
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

	if !tk.UserType.IsIn(types.UserType_SUPPER) {
		response.Failed(w, exception.NewPermissionDeny("only system admin can operate"))
		return
	}

	conf, err := h.service.GetConfig()
	if err != nil {
		response.Failed(w, err)
		return
	}

	conf.Desensitize()
	response.Success(w, conf)
	return
}

func (h *handler) TestEmailSend(w http.ResponseWriter, r *http.Request) {
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

	if !tk.UserType.IsIn(types.UserType_SUPPER) {
		response.Failed(w, exception.NewPermissionDeny("only system admin can operate"))
		return
	}

	req := mail.NewDeaultTestSendRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	if err := req.Send(); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "ok")
	return
}

func (h *handler) SettingEmail(w http.ResponseWriter, r *http.Request) {
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

	if !tk.UserType.IsIn(types.UserType_SUPPER) {
		response.Failed(w, exception.NewPermissionDeny("only system admin can operate"))
		return
	}

	req := mail.NewDefaultConfig()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	err = h.service.UpdateEmail(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, req)
	return
}

func (h *handler) TestSMSSend(w http.ResponseWriter, r *http.Request) {
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

	if !tk.UserType.IsIn(types.UserType_SUPPER) {
		response.Failed(w, exception.NewPermissionDeny("only system admin can operate"))
		return
	}

	req := sms.NewDeaultTestSendRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	if err := req.Send(); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "ok")
	return
}

func (h *handler) SettingSMS(w http.ResponseWriter, r *http.Request) {
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

	if !tk.UserType.IsIn(types.UserType_SUPPER) {
		response.Failed(w, exception.NewPermissionDeny("only system admin can operate"))
		return
	}

	req := sms.NewDefaultConfig()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	err = h.service.UpdateSMS(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, req)
	return
}

func (h *handler) SettingVerifyCode(w http.ResponseWriter, r *http.Request) {
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

	if !tk.UserType.IsIn(types.UserType_SUPPER) {
		response.Failed(w, exception.NewPermissionDeny("only system admin can operate"))
		return
	}

	req := verifycode.NewDefaultConfig()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	err = h.service.UpdateVerifyCode(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, req)
	return
}
