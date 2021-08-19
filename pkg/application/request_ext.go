package application

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/http/request"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewDescriptApplicationRequest new实例
func NewDescriptApplicationRequest() *DescribeApplicationRequest {
	return &DescribeApplicationRequest{}
}

// Validate 校验详情查询请求
func (req *DescribeApplicationRequest) Validate() error {
	if req.Id == "" && req.ClientId == "" {
		return errors.New("application id or client_id is required")
	}

	return nil
}

// NewQueryApplicationRequest 列表查询请求
func NewQueryApplicationRequest(pageReq *request.PageRequest) *QueryApplicationRequest {
	return &QueryApplicationRequest{
		Page: &pageReq.PageRequest,
	}
}

// NewCreateApplicatonRequest 请求
func NewCreateApplicatonRequest() *CreateApplicatonRequest {
	return &CreateApplicatonRequest{
		AccessTokenExpireSecond:  DefaultAccessTokenExpireSecond,
		RefreshTokenExpireSecond: DefaultRefreshTokenExpiredSecond,
		ClientType:               ClientType_PUBLIC,
	}
}

// Validate 请求校验
func (req *CreateApplicatonRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateApplicatonRequest) UpdateOwner(tk *token.Token) {
	req.CreateBy = tk.Account
	req.Domain = tk.Domain
}

// NewDeleteApplicationRequestWithID todo
func NewDeleteApplicationRequestWithID(id string) *DeleteApplicationRequest {
	return &DeleteApplicationRequest{Id: id}
}
