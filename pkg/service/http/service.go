package http

import (
	"fmt"
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/service"
)

func (h *handler) QueryService(w http.ResponseWriter, r *http.Request) {
	page := request.NewPageRequestFromHTTP(r)
	req := service.NewQueryServiceRequest(page)

	apps, err := h.service.QueryService(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
	return
}

func (h *handler) CreateService(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := service.NewCreateServiceRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.UserID = tk.UserID

	d, err := h.service.CreateService(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

func (h *handler) GetService(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := service.NewDescriptServiceRequest()
	req.Name = rctx.PS.ByName("name")
	fmt.Println(req.Name)

	d, err := h.service.DescribeService(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}

// DestroyService 销毁服务
func (h *handler) DestroyService(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	if err := h.service.DeleteService(rctx.PS.ByName("name")); err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, "delete ok")
	return
}
