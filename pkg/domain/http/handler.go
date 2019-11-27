package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg/domain"
)

type handler struct {
	service domain.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	router.AddProtected("GET", "/", h.ListDomains)
}

func (h *handler) ListDomains(w http.ResponseWriter, r *http.Request) {

}
