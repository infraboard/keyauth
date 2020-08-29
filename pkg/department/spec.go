package department

import (
	"fmt"
	"net/http"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/http/request"
)

const (
	// DefaultDepartmentName 默认部门名称
	DefaultDepartmentName = "default"
)

// Service 服务
type Service interface {
	QueryDepartment(*QueryDepartmentRequest) (*Set, error)
	DescribeDepartment(*DescribeDeparmentRequest) (*Department, error)
	CreateDepartment(*CreateDepartmentRequest) (*Department, error)
	UpdateDepartment(*UpdateDepartmentRequest) (*Department, error)
	DeleteDepartment(*DeleteDepartmentRequest) error
}

// NewQueryDepartmentRequestFromHTTP 列表查询请求
func NewQueryDepartmentRequestFromHTTP(r *http.Request) *QueryDepartmentRequest {
	req := &QueryDepartmentRequest{
		PageRequest: request.NewPageRequestFromHTTP(r),
		Session:     token.NewSession(),
	}

	qs := r.URL.Query()
	pid := qs.Get("parent_id")
	if pid != "" {
		req.ParentID = &pid
	}
	req.Keywords = qs.Get("keywords")

	return req
}

// QueryDepartmentRequest todo
type QueryDepartmentRequest struct {
	*token.Session
	*request.PageRequest
	ParentID *string
	Keywords string
}

// NewDescriptDepartmentRequest new实例
func NewDescriptDepartmentRequest() *DescribeDeparmentRequest {
	return &DescribeDeparmentRequest{}
}

// NewDescriptDepartmentRequestWithID new实例
func NewDescriptDepartmentRequestWithID(id string) *DescribeDeparmentRequest {
	return &DescribeDeparmentRequest{ID: id}
}

// DescribeDeparmentRequest 详情查询
type DescribeDeparmentRequest struct {
	ID   string
	Name string
}

// Validate 参数校验
func (req *DescribeDeparmentRequest) Validate() error {
	if req.ID == "" && req.Name == "" {
		return fmt.Errorf("id or name required")
	}

	return nil
}

// NewDeleteDepartmentRequestWithID todo
func NewDeleteDepartmentRequestWithID(id string) *DeleteDepartmentRequest {
	return &DeleteDepartmentRequest{
		Session: token.NewSession(),
		ID:      id,
	}
}

// DeleteDepartmentRequest todo
type DeleteDepartmentRequest struct {
	*token.Session
	ID string
}

// Validate todo
func (req *DeleteDepartmentRequest) Validate() error {
	if req.ID == "" {
		return fmt.Errorf("department id required")
	}

	if req.GetToken() == nil {
		return fmt.Errorf("token required")
	}

	return nil
}
