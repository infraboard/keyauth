package pkg

import (
	"fmt"

	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/counter"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/geoip"
	"github.com/infraboard/keyauth/pkg/ip2region"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/provider"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/keyauth/pkg/storage"
	"github.com/infraboard/keyauth/pkg/system"
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
	// Micro todo
	Micro micro.Service
	// Role 角色服务
	Role role.Service
	// Endpoint 端点服务
	Endpoint endpoint.Service
	// Policy 厕所里
	Policy policy.Service
	// Department 部分服务
	Department department.Service
	// Namespace todo
	Namespace namespace.Service
	// Permission 权限服务
	Permission permission.Service
	// Counter 自增ID服务
	Counter counter.Service
	// LDAP ldap服务
	LDAP provider.LDAP
	// GEOIP geoip服务
	GEOIP geoip.Service
	// IP2Region ip位置查询
	IP2Region ip2region.Service
	// Storage 对象存储服务
	Storage storage.Service
	// Session 审计服务
	Session session.Service
	// System 系统服务
	System system.Service
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
	case micro.Service:
		if Micro != nil {
			registryError(name)
		}
		Micro = value
		addService(name, svr)
	case role.Service:
		if Role != nil {
			registryError(name)
		}
		Role = value
		addService(name, svr)
	case endpoint.Service:
		if Endpoint != nil {
			registryError(name)
		}
		Endpoint = value
		addService(name, svr)
	case policy.Service:
		if Policy != nil {
			registryError(name)
		}
		Policy = value
		addService(name, svr)
	case department.Service:
		if Department != nil {
			registryError(name)
		}
		Department = value
		addService(name, svr)
	case namespace.Service:
		if Namespace != nil {
			registryError(name)
		}
		Namespace = value
		addService(name, svr)
	case permission.Service:
		if Permission != nil {
			registryError(name)
		}
		Permission = value
		addService(name, svr)
	case counter.Service:
		if Counter != nil {
			registryError(name)
		}
		Counter = value
		addService(name, svr)
	case provider.LDAP:
		if LDAP != nil {
			registryError(name)
		}
		LDAP = value
		addService(name, svr)
	case geoip.Service:
		if LDAP != nil {
			registryError(name)
		}
		GEOIP = value
		addService(name, svr)
	case storage.Service:
		if Storage != nil {
			registryError(name)
		}
		Storage = value
		addService(name, svr)
	case ip2region.Service:
		if IP2Region != nil {
			registryError(name)
		}
		IP2Region = value
		addService(name, svr)
	case session.Service:
		if Session != nil {
			registryError(name)
		}
		Session = value
		addService(name, svr)
	case system.Service:
		if System != nil {
			registryError(name)
		}
		System = value
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
