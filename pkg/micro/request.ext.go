package micro

import (
	"errors"
	"fmt"

	"github.com/infraboard/mcube/http/request"
)

// NewQueryMicroRequest 列表查询请求
func NewQueryMicroRequest(pageReq *request.PageRequest) *QueryMicroRequest {
	return &QueryMicroRequest{
		Page: &pageReq.PageRequest,
	}
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

// Validate 校验详情查询请求
func (req *DescribeMicroRequest) Validate() error {
	if req.Id == "" && req.Name == "" && req.Account == "" {
		return errors.New("id, name or account is required")
	}

	return nil
}

// NewDeleteMicroRequestWithID todo
func NewDeleteMicroRequestWithID(id string) *DeleteMicroRequest {
	return &DeleteMicroRequest{
		Id: id,
	}
}

// Validate todo
func (req *DeleteMicroRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("micro service id required")
	}

	return nil
}
