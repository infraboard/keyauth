package http

import (
	"net/http"

	"github.com/infraboard/keyauth/apps/otp"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
)

// enable  OTP
func (h *handler) CreateOTP(w http.ResponseWriter, r *http.Request) {
	req := otp.NewCreateOTPAuthRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	ins, err := h.service.CreateOTPAuth(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}
func (h *handler) GetOTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	account := ctx.PS.ByName("account")
	req := otp.NewDescribeOTPAuthRequestWithName(account)
	ins, err := h.service.DescribeOTPAuth(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) DeleteOTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	account := ctx.PS.ByName("account")
	req := otp.NewDeleteOTPAuthRequestWithName(account)
	ins, err := h.service.DeleteOTPAuth(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) UpdateOTP(w http.ResponseWriter, r *http.Request) {
	req := otp.NewUpdateOTPStatusRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	d, err := h.service.UpdateOTPAuthStatus(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, d)
}
