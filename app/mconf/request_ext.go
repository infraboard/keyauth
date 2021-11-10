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

func (req *DescribeGroupRequest) Validate() error {
	return validate.Struct(req)
}

func (req *DeleteGroupRequest) Validate() error {
	return validate.Struct(req)
}

func NewDescribeGroupRequestWithName(name string) *DescribeGroupRequest {
	return &DescribeGroupRequest{Name: name}
}

func NewDeleteGroupRequestWithName(name string) *DeleteGroupRequest {
	return &DeleteGroupRequest{Name: name}
}

// NewQueryItemRequest 列表查询请求
func NewQueryItemRequest(pageReq *request.PageRequest) *QueryItemRequest {
	return &QueryItemRequest{
		Page: &pageReq.PageRequest,
	}
}

// NewAddItemToGroupRequest todo
func NewAddItemToGroupRequest() *AddItemToGroupRequest {
	return &AddItemToGroupRequest{}
}
