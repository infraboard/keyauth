package tag

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/infraboard/keyauth/common/tools"
	"github.com/infraboard/mcube/http/request"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

func (req *CreateTagRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateTagRequest) GenUUID() string {
	switch req.ScopeType {
	case ScopeType_GLOBAL:
		return "G." + req.KeyName
	case ScopeType_DOMAIN:
		return "D." + req.KeyName
	case ScopeType_NAMESPACE:
		return "N." + tools.GenHashID(req.Namespace+req.KeyName)
	default:
		return "O." + req.KeyName
	}
}

func (req *QueryTagKeyRequest) Validate() error {
	return validate.Struct(req)
}

func (req *QueryTagValueRequest) Validate() error {
	return validate.Struct(req)
}

func (req *DescribeTagRequest) Validate() error {
	return validate.Struct(req)
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

// NewDescribeTagRequestWithID todo
func NewDescribeTagRequestWithID(id string) *DescribeTagRequest {
	return &DescribeTagRequest{TagId: id}
}

// NewDeleteTagRequestWithID todo
func NewDeleteTagRequestWithID(id string) *DeleteTagRequest {
	return &DeleteTagRequest{TagId: id}
}
