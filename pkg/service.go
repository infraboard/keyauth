package pkg

import (
	"fmt"

	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/service"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	// Domain 服务
	Domain domain.Service
	// User 用户服务
	User user.Service
	// Application 应用
	Application application.Service
	// Token 令牌服务
	Token token.Service
	// MicroService 服务
	MicroService service.Service
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
			registryError(name)
		}
		Domain = value
		addService(name, svr)
	case user.Service:
		if User != nil {
			registryError(name)
		}
		User = value
		addService(name, svr)
	case application.Service:
		if Application != nil {
			registryError(name)
		}
		Application = value
		addService(name, svr)
	case token.Service:
		if Token != nil {
			registryError(name)
		}
		Token = value
		addService(name, svr)
	case service.Service:
		if MicroService != nil {
			registryError(name)
		}
		MicroService = value
		addService(name, svr)
	default:
		panic(fmt.Sprintf("unknown service type %s", name))
	}
}

func registryError(name string) {
	panic("service " + name + " has registried")
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
