package department

import (
	"fmt"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/http/request"
)

// Service 服务
type Service interface {
	QueryDepartment(*QueryDepartmentRequest) (*Set, error)
	DescribeDepartment(*DescribeDeparmentRequest) (*Department, error)
	CreateDepartment(*CreateDepartmentRequest) (*Department, error)
	DeleteDepartment(id string) error
}

// NewQueryDepartmentRequest 列表查询请求
func NewQueryDepartmentRequest(req *request.PageRequest) *QueryDepartmentRequest {
	return &QueryDepartmentRequest{
		PageRequest: req,
		Session:     token.NewSession(),
	}
}

// QueryDepartmentRequest todo
type QueryDepartmentRequest struct {
	*token.Session
	*request.PageRequest
	ParentID *string
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
