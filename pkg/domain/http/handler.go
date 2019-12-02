package http

import (
	"strconv"
	"errors"
	"fmt"
	"net/http"

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
}

func (h *handler) Config() error {
	if pkg.Domain == nil {
		return errors.New("denpence Domain service is nil")
	}

	h.service = pkg.Domain
	return nil
}

func (h *handler) ListDomains(w http.ResponseWriter, r *http.Request) {
	ps := r.URL.Query().Get("page_size")
	pn := r.URL.Query().Get("page_number")
	fmt.Println(ps, pn)
	req := domain.NewRequest()

	resp := new(response.PageData)
	resp.PageSize, _ := strconv.ParseUint(ps, 10, 16)
	h.service.ListDomain(req)
	response.Success(w, "ok")
	return
}

func init() {
	pkg.RegistryHTTP("domains", api)
}
