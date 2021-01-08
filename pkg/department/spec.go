package department

import (
	"fmt"
	"net/http"

	"github.com/infraboard/keyauth/pkg/token/session"
	"github.com/infraboard/mcube/exception"
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

	QueryApplicationForm(*QueryApplicationFormRequet) (*ApplicationFormSet, error)
	DescribeApplicationForm(*DescribeApplicationFormRequet) (*ApplicationForm, error)
	JoinDepartment(*JoinDepartmentRequest) (*ApplicationForm, error)
	DealApplicationForm(*DealApplicationFormRequest) (*ApplicationForm, error)
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
	req.WithRole = qs.Get("with_role") == "true"
	return req
}

// NewQueryDepartmentRequest todo
func NewQueryDepartmentRequest() *QueryDepartmentRequest {
	return &QueryDepartmentRequest{
		Session:       session.NewSession(),
		PageRequest:   request.NewPageRequest(20, 1),
		SkipItems:     false,
		WithSubCount:  false,
		WithUserCount: false,
	}
}

// QueryDepartmentRequest todo
type QueryDepartmentRequest struct {
	*session.Session
	*request.PageRequest
	ParentID      *string
	Keywords      string
	SkipItems     bool
	WithSubCount  bool
	WithUserCount bool
	WithRole      bool
}

// Validate todo
func (req *QueryDepartmentRequest) Validate() error {
	if req.GetToken() == nil {
		return fmt.Errorf("token required")
	}

	return nil
}

// NewDescribeDepartmentRequest new实例
func NewDescribeDepartmentRequest() *DescribeDeparmentRequest {
	return &DescribeDeparmentRequest{
		Session: session.NewSession(),
	}
}

// NewDescribeDepartmentRequestWithID new实例
func NewDescribeDepartmentRequestWithID(id string) *DescribeDeparmentRequest {
	req := NewDescribeDepartmentRequest()
	req.ID = id
	return req
}

// DescribeDeparmentRequest 详情查询
type DescribeDeparmentRequest struct {
	*session.Session
	ID            string
	Name          string
	WithSubCount  bool
	WithUserCount bool
	WithRole      bool
}

func (req *DescribeDeparmentRequest) String() string {
	if req.ID != "" {
		return req.ID
	}

	return req.Name
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
		Session: session.NewSession(),
		ID:      id,
	}
}

// DeleteDepartmentRequest todo
type DeleteDepartmentRequest struct {
	*session.Session
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

// NewJoinDepartmentRequest todo
func NewJoinDepartmentRequest() *JoinDepartmentRequest {
	return &JoinDepartmentRequest{
		Session: session.NewSession(),
	}
}

// JoinDepartmentRequest todo
type JoinDepartmentRequest struct {
	Account      string `bson:"account" json:"account" validate:"required"`             // 申请人
	DepartmentID string `bson:"department_id" json:"department_id" validate:"required"` // 申请加入的部门
	Message      string `bson:"message" json:"message"`                                 // 留言

	*session.Session `bson:"-"`
}

// Validate todo
func (req *JoinDepartmentRequest) Validate() error {
	if req.GetToken() == nil {
		return fmt.Errorf("token required")
	}

	return validate.Struct(req)
}

// NewDefaultDealApplicationFormRequest todo
func NewDefaultDealApplicationFormRequest() *DealApplicationFormRequest {
	return &DealApplicationFormRequest{
		Session: session.NewSession(),
	}
}

// DealApplicationFormRequest todo
type DealApplicationFormRequest struct {
	*session.Session
	ID      string                `json:"id"`      // 用户
	Status  ApplicationFormStatus `json:"status"`  // 状态
	Message string                `json:"message"` // 备注
}

// Validate todo
func (req *DealApplicationFormRequest) Validate() error {
	if req.ID == "" {
		return fmt.Errorf("account required one")
	}

	if req.Status.Is(Pending) {
		return fmt.Errorf("status must be passed or deny")
	}

	return nil
}

// NewQueryApplicationFormRequestFromHTTP todo
func NewQueryApplicationFormRequestFromHTTP(r *http.Request) (*QueryApplicationFormRequet, error) {
	req := NewQueryApplicationFormRequet()
	req.PageRequest = request.NewPageRequestFromHTTP(r)

	qs := r.URL.Query()
	req.Account = qs.Get("account")
	req.DepartmentID = qs.Get("department_id")
	req.SkipItems = qs.Get("skip_items") == "true"

	status, err := ParseApplicationFormStatus(qs.Get("status"))
	if err != nil {
		return nil, exception.NewBadRequest("parse status error, %s", err)
	}
	req.Status = &status
	return req, nil
}

// NewQueryApplicationFormRequet todo
func NewQueryApplicationFormRequet() *QueryApplicationFormRequet {
	return &QueryApplicationFormRequet{
		Session:     session.NewSession(),
		PageRequest: request.NewPageRequest(20, 1),
		SkipItems:   false,
	}
}

// QueryApplicationFormRequet todo
type QueryApplicationFormRequet struct {
	*request.PageRequest
	*session.Session
	Account      string
	DepartmentID string
	Status       *ApplicationFormStatus
	SkipItems    bool
}

// Validate 请求参数校验
func (req *QueryApplicationFormRequet) Validate() error {
	if req.GetToken() == nil {
		return fmt.Errorf("token required")
	}

	if req.Account == "" && req.DepartmentID == "" {
		return fmt.Errorf("account and department_id required one")
	}

	return nil
}

// NewDescribeApplicationFormRequetWithID todo
func NewDescribeApplicationFormRequetWithID(id string) *DescribeApplicationFormRequet {
	return &DescribeApplicationFormRequet{
		Session: session.NewSession(),
		ID:      id,
	}
}

// DescribeApplicationFormRequet todo
type DescribeApplicationFormRequet struct {
	*session.Session
	ID string
}
