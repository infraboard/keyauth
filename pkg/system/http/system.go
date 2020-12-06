package http

import (
	"net/http"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/system/notify"
	"github.com/infraboard/keyauth/pkg/system/notify/mail"
	"github.com/infraboard/keyauth/pkg/system/notify/sms"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func (h *handler) GetSystemConfig(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	if !tk.UserType.Is(types.SupperAccount) {
		response.Failed(w, exception.NewPermissionDeny("only system admin can operate"))
		return
	}

	conf, err := h.service.GetConfig()
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, conf)
	return
}

func (h *handler) TestEmailSend(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	if !tk.UserType.Is(types.SupperAccount) {
		response.Failed(w, exception.NewPermissionDeny("only system admin can operate"))
		return
	}

	req := notify.NewSendMailRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	conf, err := h.service.GetConfig()
	if err != nil {
		response.Failed(w, err)
		return
	}

	sd, err := mail.NewSender(conf.Email)
	if err := sd.Send(req); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "ok")
	return
}

func (h *handler) EmailSetting(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	if !tk.UserType.Is(types.SupperAccount) {
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

	response.Success(w, "ok")
	return
}

func (h *handler) TestSMSSend(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	if !tk.UserType.Is(types.SupperAccount) {
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

func (h *handler) SMSSetting(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	if !tk.UserType.Is(types.SupperAccount) {
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

	response.Success(w, "ok")
	return
}
