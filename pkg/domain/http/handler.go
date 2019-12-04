package http

import (
	"errors"
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
)

var (
	api = &handler{}
)

type handler struct {
	service domain.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	router.AddProtected("GET", "/", h.ListDomains)
	router.AddProtected("GET", "/:id", h.GetDomain)
}

func (h *handler) Config() error {
	if pkg.Domain == nil {
		return errors.New("denpence Domain service is nil")
	}

	h.service = pkg.Domain
	return nil
}

func (h *handler) ListDomains(w http.ResponseWriter, r *http.Request) {
	page := request.LoadPagginFromReq(r)
	req := domain.NewPageRequest(page)

	dommains, total, err := h.service.ListDomain(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	data := response.PageData{
		PageRequest: page,
		TotalCount:  uint(total),
		List:        dommains,
	}
	response.Success(w, data)
	return
}

func (h *handler) GetDomain(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	d, err := h.service.GetDomainByID(rctx.PS.ByName("id"))
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func init() {
	pkg.RegistryHTTP("domains", api)
}
