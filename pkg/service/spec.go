package service

import (
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/router"
)

// Service token管理服务
type Service interface {
	CreateService(req *CreateServiceRequest) (*MicroService, error)
	QueryService(req *QueryServiceRequest) (*MicroServiceSet, error)
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
