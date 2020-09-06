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
	req := NewQueryDepartmentRequest()
	req.PageRequest = request.NewPageRequestFromHTTP(r)

	qs := r.URL.Query()
	pid := qs.Get("parent_id")
	if pid != "*" {
		req.ParentID = &pid
	}
	req.Keywords = qs.Get("keywords")
	req.WithSubCount = qs.Get("with_sub_count") == "true"
	req.WithUserCount = qs.Get("with_user_count") == "true"
	return req
}

// NewQueryDepartmentRequest todo
func NewQueryDepartmentRequest() *QueryDepartmentRequest {
	return &QueryDepartmentRequest{
		Session:       token.NewSession(),
		PageRequest:   request.NewPageRequest(20, 1),
		SkipItems:     false,
		WithSubCount:  false,
		WithUserCount: false,
	}
}

// QueryDepartmentRequest todo
type QueryDepartmentRequest struct {
	*token.Session
	*request.PageRequest
	ParentID      *string
	Keywords      string
	SkipItems     bool
	WithSubCount  bool
	WithUserCount bool
}

// Validate todo
func (req *QueryDepartmentRequest) Validate() error {
	if req.GetToken() == nil {
		return fmt.Errorf("token required")
	}

	return nil
}

// NewDescriptDepartmentRequest new实例
func NewDescriptDepartmentRequest() *DescribeDeparmentRequest {
	return &DescribeDeparmentRequest{
		Session: token.NewSession(),
	}
}

// NewDescriptDepartmentRequestWithID new实例
func NewDescriptDepartmentRequestWithID(id string) *DescribeDeparmentRequest {
	req := NewDescriptDepartmentRequest()
	req.ID = id
	return req
}

// DescribeDeparmentRequest 详情查询
type DescribeDeparmentRequest struct {
	*token.Session
	ID            string
	Name          string
	WithSubCount  bool
	WithUserCount bool
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
