package department

import (
	"fmt"
	"net/http"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
)

const (
	// DefaultDepartmentName 默认部门名称
	DefaultDepartmentName = "default"
)

// NewQueryDepartmentRequestFromHTTP 列表查询请求
func NewQueryDepartmentRequestFromHTTP(r *http.Request) *QueryDepartmentRequest {
	req := NewQueryDepartmentRequest()
	req.Page = &request.NewPageRequestFromHTTP(r).PageRequest

	qs := r.URL.Query()
	pid := qs.Get("parent_id")
	if pid != "*" {
		req.ParentId = pid
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
		Page:          &request.NewPageRequest(20, 1).PageRequest,
		SkipItems:     false,
		WithSubCount:  false,
		WithUserCount: false,
	}
}

// Validate todo
func (req *QueryDepartmentRequest) Validate() error {
	return nil
}

// NewDescribeDepartmentRequest new实例
func NewDescribeDepartmentRequest() *DescribeDeparmentRequest {
	return &DescribeDeparmentRequest{}
}

// NewDescribeDepartmentRequestWithID new实例
func NewDescribeDepartmentRequestWithID(id string) *DescribeDeparmentRequest {
	req := NewDescribeDepartmentRequest()
	req.Id = id
	return req
}

// Validate 参数校验
func (req *DescribeDeparmentRequest) Validate() error {
	if req.Id == "" && req.Name == "" {
		return fmt.Errorf("id or name required")
	}

	return nil
}

// NewDeleteDepartmentRequestWithID todo
func NewDeleteDepartmentRequestWithID(id string) *DeleteDepartmentRequest {
	return &DeleteDepartmentRequest{
		Id: id,
	}
}

// Validate todo
func (req *DeleteDepartmentRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("department id required")
	}

	return nil
}

// NewJoinDepartmentRequest todo
func NewJoinDepartmentRequest() *JoinDepartmentRequest {
	return &JoinDepartmentRequest{}
}

// Validate todo
func (req *JoinDepartmentRequest) Validate() error {
	return validate.Struct(req)
}

func (req *JoinDepartmentRequest) UpdateOwner(tk *token.Token) {
	req.Domain = tk.Domain
	req.Account = tk.Account
}

// NewDefaultDealApplicationFormRequest todo
func NewDefaultDealApplicationFormRequest() *DealApplicationFormRequest {
	return &DealApplicationFormRequest{}
}

// Validate todo
func (req *DealApplicationFormRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("account required one")
	}

	if req.Status.Equal(ApplicationFormStatus_PENDDING) {
		return fmt.Errorf("status must be passed or deny")
	}

	return nil
}

// NewQueryApplicationFormRequestFromHTTP todo
func NewQueryApplicationFormRequestFromHTTP(r *http.Request) (*QueryApplicationFormRequet, error) {
	req := NewQueryApplicationFormRequet()
	req.Page = &request.NewPageRequestFromHTTP(r).PageRequest

	qs := r.URL.Query()
	req.Account = qs.Get("account")
	req.DepartmentId = qs.Get("department_id")
	req.SkipItems = qs.Get("skip_items") == "true"

	status := qs.Get("status")
	if status != "" {
		status, err := ParseApplicationFormStatusFromString(status)
		if err != nil {
			return nil, exception.NewBadRequest("parse status error, %s", err)
		}
		req.Status = status
	}

	return req, nil
}

// NewQueryApplicationFormRequet todo
func NewQueryApplicationFormRequet() *QueryApplicationFormRequet {
	return &QueryApplicationFormRequet{
		Page:      &request.NewPageRequest(20, 1).PageRequest,
		SkipItems: false,
	}
}

// Validate 请求参数校验
func (req *QueryApplicationFormRequet) Validate() error {
	if req.Account == "" && req.DepartmentId == "" {
		return fmt.Errorf("account and department_id required one")
	}

	return nil
}

// NewDescribeApplicationFormRequetWithID todo
func NewDescribeApplicationFormRequetWithID(id string) *DescribeApplicationFormRequet {
	return &DescribeApplicationFormRequet{
		Id: id,
	}
}
