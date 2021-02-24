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

// NewDescribeServiceRequestWithClientID new实例
func NewDescribeServiceRequestWithClientID(account string) *DescribeMicroRequest {
	req := NewDescribeServiceRequest()
	req.ClientId = account
	return req
}

// NewDescribeServiceRequest new实例
func NewDescribeServiceRequest() *DescribeMicroRequest {
	return &DescribeMicroRequest{}
}

// Validate 校验详情查询请求
func (req *DescribeMicroRequest) Validate() error {
	if req.Id == "" && req.Name == "" && req.ClientId == "" {
		return errors.New("id, name or client_id is required")
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
