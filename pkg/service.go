package pkg

import (
	"fmt"

	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	// Domain 服务
	Domain domain.Service
	// User 用户服务
	User user.Service
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
	case user.Service:
		if User != nil {
			panic("service " + name + " has registried")
		}
		User = value
		addService(name, svr)
	default:
		panic(fmt.Sprintf("unknown service type %s", name))
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
