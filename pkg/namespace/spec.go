package namespace

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg/token/session"
)

// Service todo
type Service interface {
	CreateNamespace(req *CreateNamespaceRequest) (*Namespace, error)
	QueryNamespace(req *QueryNamespaceRequest) (*Set, error)
	DescribeNamespace(req *DescriptNamespaceRequest) (*Namespace, error)
	DeleteNamespace(req *DeleteNamespaceRequest) error
}

// NewQueryNamespaceRequestFromHTTP 列表查询请求
func NewQueryNamespaceRequestFromHTTP(r *http.Request) *QueryNamespaceRequest {
	qs := r.URL.Query()
	return &QueryNamespaceRequest{
		Session:           session.NewSession(),
		PageRequest:       request.NewPageRequestFromHTTP(r),
		DepartmentID:      qs.Get("department_id"),
		WithDepartment:    qs.Get("with_department") == "true",
		WithSubDepartment: qs.Get("with_sub_department") == "true",
	}
}

// NewQueryNamespaceRequest 列表查询请求
func NewQueryNamespaceRequest(pageReq *request.PageRequest) *QueryNamespaceRequest {
	return &QueryNamespaceRequest{
		Session:        session.NewSession(),
		PageRequest:    pageReq,
		WithDepartment: false,
	}
}

// QueryNamespaceRequest 查询应用列表
type QueryNamespaceRequest struct {
	*session.Session
	*request.PageRequest
	DepartmentID      string
	WithSubDepartment bool
	WithDepartment    bool
}

// NewNewDescriptNamespaceRequestWithID todo
func NewNewDescriptNamespaceRequestWithID(id string) *DescriptNamespaceRequest {
	req := NewDescriptNamespaceRequest()
	req.ID = id
	return req
}

// NewDescriptNamespaceRequest new实例
func NewDescriptNamespaceRequest() *DescriptNamespaceRequest {
	return &DescriptNamespaceRequest{
		Session: session.NewSession(),
	}
}

// DescriptNamespaceRequest 查询应用详情
type DescriptNamespaceRequest struct {
	*session.Session
	ID             string `json:"id,omitempty"`
	WithDepartment bool
}

func (req *DescriptNamespaceRequest) String() string {
	return req.ID
}

// Validate 校验详情查询请求
func (req *DescriptNamespaceRequest) Validate() error {
	if req.ID == "" {
		return errors.New("id  is required")
	}

	return nil
}

// NewDeleteNamespaceRequestWithID todo
func NewDeleteNamespaceRequestWithID(id string) *DeleteNamespaceRequest {
	return &DeleteNamespaceRequest{
		Session: session.NewSession(),
		ID:      id,
	}
}

// DeleteNamespaceRequest todo
type DeleteNamespaceRequest struct {
	*session.Session
	ID string
}

// Validate todo
func (req *DeleteNamespaceRequest) Validate() error {
	tk := req.GetToken()
	if tk == nil {
		return fmt.Errorf("token required")
	}

	if req.ID == "" {
		return fmt.Errorf("id required")
	}

	return nil
}
