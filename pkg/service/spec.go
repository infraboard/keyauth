package service

import "github.com/infraboard/mcube/http/router"

// Service token管理服务
type Service interface {
	CreateService(req *CreateServiceRequest) (*MicroService, error)
	Registry(req *RegistryRequest) error
}

// RegistryRequest 服务注册请求
type RegistryRequest struct {
	ServiceToken string
	EntrySet     router.EntrySet
}
