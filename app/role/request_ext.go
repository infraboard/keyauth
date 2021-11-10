package role

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewQueryRoleRequestFromHTTP 列表查询请求
func NewQueryRoleRequestFromHTTP(r *http.Request) *QueryRoleRequest {
	page := request.NewPageRequestFromHTTP(r)

	req := NewQueryRoleRequest(page)
	return req
}

// NewQueryRoleRequest 列表查询请求
func NewQueryRoleRequest(pageReq *request.PageRequest) *QueryRoleRequest {
	return &QueryRoleRequest{
		Page: &pageReq.PageRequest,
	}
}

// Validate todo
func (req *QueryRoleRequest) Validate() error {
	return nil
}

// NewDescribeRoleRequestWithID todo
func NewDescribeRoleRequestWithID(id string) *DescribeRoleRequest {
	return &DescribeRoleRequest{
		Id:              id,
		WithPermissions: false,
	}
}

// NewDescribeRoleRequestWithName todo
func NewDescribeRoleRequestWithName(name string) *DescribeRoleRequest {
	return &DescribeRoleRequest{
		Name:            name,
		WithPermissions: false,
	}
}

// Validate todo
func (req *DescribeRoleRequest) Validate() error {
	if req.Id == "" && req.Name == "" {
		return fmt.Errorf("id or name required")
	}

	return nil
}

// NewDeleteRoleWithID todo
func NewDeleteRoleWithID(id string) *DeleteRoleRequest {
	return &DeleteRoleRequest{
		Id: id,
	}
}

// NewAddPermissionToRoleRequest todo
func NewAddPermissionToRoleRequest() *AddPermissionToRoleRequest {
	return &AddPermissionToRoleRequest{
		Permissions: []*CreatePermssionRequest{},
	}
}

func (req *AddPermissionToRoleRequest) Validate() error {
	return validate.Struct(req)
}

func (req *AddPermissionToRoleRequest) Length() int {
	return len(req.Permissions)
}

// NewRemovePermissionFromRoleRequest todo
func NewRemovePermissionFromRoleRequest() *RemovePermissionFromRoleRequest {
	return &RemovePermissionFromRoleRequest{
		PermissionId: []string{},
	}
}

func (req *RemovePermissionFromRoleRequest) Validate() error {
	return validate.Struct(req)
}

// NewQueryPermissionRequest todo
func NewQueryPermissionRequest(pageReq *request.PageRequest) *QueryPermissionRequest {
	return &QueryPermissionRequest{
		Page: &pageReq.PageRequest,
	}
}

// NewQueryPermissionRequestFromHTTP 列表查询请求
func NewQueryPermissionRequestFromHTTP(r *http.Request) *QueryPermissionRequest {
	page := request.NewPageRequestFromHTTP(r)
	req := NewQueryPermissionRequest(page)

	return req
}

func NewDescribePermissionRequestWithID(id string) *DescribePermissionRequest {
	return &DescribePermissionRequest{Id: id}
}

func (req *DescribePermissionRequest) Validate() error {
	if req.Id == "" {
		return exception.NewBadRequest("id required")
	}
	return nil
}

func NewUpdatePermissionRequest() *UpdatePermissionRequest {
	return &UpdatePermissionRequest{}
}

func (req *UpdatePermissionRequest) Validate() error {
	if req.Id == "" {
		return exception.NewBadRequest("id required")
	}

	return nil
}
