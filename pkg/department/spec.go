package department

import (
	"fmt"
	"net/http"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
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
	WithRole      bool
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

// JoinDepartmentRequest todo
type JoinDepartmentRequest struct {
	Account      string `bson:"account" json:"account" validate:"required"`             // 申请人
	DepartmentID string `bson:"department_id" json:"department_id" validate:"required"` // 申请加入的部门
	Message      string `bson:"message" json:"message"`                                 // 留言

	*token.Session
}

// Validate todo
func (req *JoinDepartmentRequest) Validate() error {
	return validate.Struct(req)
}

// NewApplicationForm todo
func NewApplicationForm(req *JoinDepartmentRequest) (*ApplicationForm, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk := req.GetToken()

	ins := &ApplicationForm{
		ID:                    xid.New().String(),
		CreateAt:              ftime.Now(),
		UpdateAt:              ftime.Now(),
		Creater:               tk.Account,
		JoinDepartmentRequest: req,
		Status:                Pending,
	}

	return ins, nil
}

// ApplicationForm todo
type ApplicationForm struct {
	ID       string                `bson:"_id" json:"id"`              // 部门加入申请单ID
	Creater  string                `bson:"creater" json:"creater"`     // 申请人
	CreateAt ftime.Time            `bson:"create_at" json:"create_at"` // 创建时间
	UpdateAt ftime.Time            `bson:"update_at" json:"update_at"` // 更新时间
	Status   ApplicationFormStatus `bson:"status" json:"status"`       // 状态
	*JoinDepartmentRequest
}

// DealApplicationFormRequest todo
type DealApplicationFormRequest struct {
	*token.Session
	ID     string                `json:"id"`
	Status ApplicationFormStatus `json:"status"` // 状态
}
