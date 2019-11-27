package pkg

import (
	"github.com/infraboard/keyauth/pkg/domain"
)

var (
	// Domain 服务
	Domain domain.Service
)

var (
	servers       []Service
	successLoaded []string
)

// LoadedService 查询加载成功的服务
func LoadedService() []string {
	return successLoaded
}

func addService(name string, svr Service) {
	servers = append(servers, svr)
	successLoaded = append(successLoaded, name)
}

// Service 注册上的服务必须实现的方法
type Service interface {
	Config() error
}

// RegistryService 服务实例注册
func RegistryService(name string, svr Service) {
	switch value := svr.(type) {
	case domain.Service:
		if Domain != nil {
			panic("service " + name + " has registried")
		}
		Domain = value
		addService(name, svr)
	default:
		return
	}
}

// InitService 初始化所有服务
func InitService() error {
	for _, s := range servers {
		if err := s.Config(); err != nil {
			return err
		}
	}

	return nil
}
