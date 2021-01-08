package micro

import (
	"errors"
	"fmt"

	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/token/session"
)

// Service token管理服务
type Service interface {
	CreateService(req *CreateMicroRequest) (*Micro, error)
	QueryService(req *QueryMicroRequest) (*Set, error)
	DescribeService(req *DescribeMicroRequest) (*Micro, error)
	DeleteService(req *DeleteMicroRequest) error
	RefreshServiceToken(req *DescribeMicroRequest) (*token.Token, error)
}

// NewQueryMicroRequest 列表查询请求
func NewQueryMicroRequest(pageReq *request.PageRequest) *QueryMicroRequest {
	return &QueryMicroRequest{
		PageRequest: pageReq,
	}
}

// QueryMicroRequest 查询应用列表
type QueryMicroRequest struct {
	*request.PageRequest
}

// NewDescribeServiceRequestWithAccount new实例
func NewDescribeServiceRequestWithAccount(account string) *DescribeMicroRequest {
	req := NewDescribeServiceRequest()
	req.Account = account
	return req
}

// NewDescribeServiceRequest new实例
func NewDescribeServiceRequest() *DescribeMicroRequest {
	return &DescribeMicroRequest{}
}

// DescribeMicroRequest 查询应用详情
type DescribeMicroRequest struct {
	Account string
	Name    string
	ID      string
}

// Validate 校验详情查询请求
func (req *DescribeMicroRequest) Validate() error {
	if req.ID == "" && req.Name == "" && req.Account == "" {
		return errors.New("id, name or account is required")
	}

	return nil
}

// NewDeleteMicroRequestWithID todo
func NewDeleteMicroRequestWithID(id string) *DeleteMicroRequest {
	return &DeleteMicroRequest{
		Session: session.NewSession(),
		ID:      id,
	}
}

// Validate todo
func (req *DeleteMicroRequest) Validate() error {
	tk := req.GetToken()
	if tk == nil {
		return fmt.Errorf("token required")
	}

	if req.ID == "" {
		return fmt.Errorf("micro service id required")
	}

	return nil
}

// DeleteMicroRequest todo
type DeleteMicroRequest struct {
	*session.Session
	ID string
}
