package http

import (
	"net/http"

	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/mcube/http/router"
)

// Registry 注册HTTP服务路由
func Registry(router router.SubRouter) {
	h := new(handler)
	router.AddProtected("GET", "/", h.ListDomains)
}

type handler struct {
	service domain.Service
}

func (h *handler) ListDomains(w http.ResponseWriter, r *http.Request) {

}
