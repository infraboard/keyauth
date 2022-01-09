package http

import (
	"fmt"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/response"
	"net/http"
)

func (h *handler) WechatCheck(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)

	fmt.Println(ctx)
	response.Success(w, "passed")
}
