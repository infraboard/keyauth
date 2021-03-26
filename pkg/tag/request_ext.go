package tag

import (
	"net/http"

	"github.com/infraboard/mcube/http/request"
)

func (req *CreateTagRequest) Validate() error {
	return nil
}

func (req *QueryTagKeyRequest) Validate() error {
	return nil
}

func (req *QueryTagValueRequest) Validate() error {
	return nil
}

// NewQueryTageKeyRequestFromHTTP 列表查询请求
func NewQueryTageKeyRequestFromHTTP(r *http.Request) *QueryTagKeyRequest {
	page := request.NewPageRequestFromHTTP(r)
	req := NewQueryTagKeyRequest(page)
	return req
}

// NewQueryRoleRequest 列表查询请求
func NewQueryTagKeyRequest(pageReq *request.PageRequest) *QueryTagKeyRequest {
	return &QueryTagKeyRequest{
		Page: &pageReq.PageRequest,
	}
}

// NewQueryTageValueRequestFromHTTP 列表查询请求
func NewQueryTageValueRequestFromHTTP(r *http.Request) *QueryTagValueRequest {
	page := request.NewPageRequestFromHTTP(r)
	req := NewQueryTagValueRequest(page)
	return req
}

// NewQueryTagValueRequest 列表查询请求
func NewQueryTagValueRequest(pageReq *request.PageRequest) *QueryTagValueRequest {
	return &QueryTagValueRequest{
		Page: &pageReq.PageRequest,
	}
}
