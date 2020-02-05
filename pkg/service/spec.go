package service

import (
	"errors"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/router"
)

// Service token管理服务
type Service interface {
	CreateService(req *CreateServiceRequest) (*MicroService, error)
	QueryService(req *QueryServiceRequest) (*MicroServiceSet, error)
	DescribeService(req *DescriptServiceRequest) (*MicroService, error)
	DeleteService(name string) error
	Registry(req *RegistryRequest) error
}

// RegistryRequest 服务注册请求
type RegistryRequest struct {
	ServiceToken string
	EntrySet     router.EntrySet
}

// NewQueryServiceRequest 列表查询请求
func NewQueryServiceRequest(pageReq *request.PageRequest) *QueryServiceRequest {
	return &QueryServiceRequest{
		PageRequest: pageReq,
	}
}

// QueryServiceRequest 查询应用列表
type QueryServiceRequest struct {
	*request.PageRequest
}

// NewDescriptServiceRequest new实例
func NewDescriptServiceRequest() *DescriptServiceRequest {
	return &DescriptServiceRequest{}
}

// DescriptServiceRequest 查询应用详情
type DescriptServiceRequest struct {
	Name     string
	ClientID string
}

// Validate 校验详情查询请求
func (req *DescriptServiceRequest) Validate() error {
	if req.ClientID == "" && req.Name == "" {
		return errors.New("id, name or client_id is required")
	}

	return nil
}
