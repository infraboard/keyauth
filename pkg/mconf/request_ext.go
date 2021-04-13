package mconf

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewQueryGroupRequest 列表查询请求
func NewQueryGroupRequest(pageReq *request.PageRequest) *QueryGroupRequest {
	return &QueryGroupRequest{
		Page: &pageReq.PageRequest,
	}
}

// NewCreateGroupRequest todo
func NewCreateGroupRequest() *CreateGroupRequest {
	return &CreateGroupRequest{
		Type: Type_GLOBAL,
	}
}

func (req *CreateGroupRequest) Validate() error {
	return validate.Struct(req)
}
