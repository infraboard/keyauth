package namespace

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg/token"
)

// NewQueryNamespaceRequestFromHTTP 列表查询请求
func NewQueryNamespaceRequestFromHTTP(r *http.Request) *QueryNamespaceRequest {
	qs := r.URL.Query()
	return &QueryNamespaceRequest{
		Page:              &request.NewPageRequestFromHTTP(r).PageRequest,
		DepartmentId:      qs.Get("department_id"),
		Name:              qs.Get("name"),
		WithDepartment:    qs.Get("with_department") == "true",
		WithSubDepartment: qs.Get("with_sub_department") == "true",
	}
}

// NewQueryNamespaceRequest 列表查询请求
func NewQueryNamespaceRequest(pageReq *request.PageRequest) *QueryNamespaceRequest {
	return &QueryNamespaceRequest{
		Page:           &pageReq.PageRequest,
		WithDepartment: false,
	}
}

func (req *QueryNamespaceRequest) UpdateOwner(tk *token.Token) {
	req.Domain = tk.Domain
	req.Account = tk.Account
}

// NewNewDescriptNamespaceRequestWithID todo
func NewNewDescriptNamespaceRequestWithID(id string) *DescriptNamespaceRequest {
	req := NewDescriptNamespaceRequest()
	req.Id = id
	return req
}

// NewDescriptNamespaceRequest new实例
func NewDescriptNamespaceRequest() *DescriptNamespaceRequest {
	return &DescriptNamespaceRequest{}
}

// Validate 校验详情查询请求
func (req *DescriptNamespaceRequest) Validate() error {
	if req.Id == "" {
		return errors.New("id  is required")
	}

	return nil
}

// NewDeleteNamespaceRequestWithID todo
func NewDeleteNamespaceRequestWithID(id string) *DeleteNamespaceRequest {
	return &DeleteNamespaceRequest{
		Id: id,
	}
}

// Validate todo
func (req *DeleteNamespaceRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("id required")
	}

	return nil
}
