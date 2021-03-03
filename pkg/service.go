package pkg

import (
	"fmt"

	"github.com/infraboard/mcube/pb/http"
	"google.golang.org/grpc"

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
	"github.com/infraboard/keyauth/pkg/verifycode"
)

var (
	// Domain 服务
	Domain domain.DomainServiceServer
	// User 用户服务
	User user.UserServiceServer
	// ApplicationUser 应用
	ApplicationUser application.UserServiceServer
	// ApplicationAdmin 应用
	ApplicationAdmin application.AdminServiceServer
	// Token 令牌服务
	Token token.TokenServiceServer
	// Micro todo
	Micro micro.MicroServiceServer
	// Role 角色服务
	Role role.RoleServiceServer
	// Endpoint 端点服务
	Endpoint endpoint.EndpointServiceServer
	// Policy 厕所里
	Policy policy.PolicyServiceServer
	// Department 部分服务
	Department department.DepartmentServiceServer
	// Namespace todo
	Namespace namespace.NamespaceServiceServer
	// Permission 权限服务
	Permission permission.PermissionServiceServer
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
	// SessionAdmin 审计服务
	SessionAdmin session.AdminServiceServer
	// SessionUser todo
	SessionUser session.UserServiceServer
	// System 系统服务
	System system.Service
	// VerifyCode 校验码服务
	VerifyCode verifycode.VerifyCodeServiceServer
)

var (
	servers       []Service
	successLoaded []string

	entrySet  = http.NewEntrySet()
	entryInit = false
)

// InitV1GRPCAPI 初始化API服务
func InitV1GRPCAPI(server *grpc.Server) {
	domain.RegisterDomainServiceServer(server, Domain)
	user.RegisterUserServiceServer(server, User)
	application.RegisterAdminServiceServer(server, ApplicationAdmin)
	application.RegisterUserServiceServer(server, ApplicationUser)
	token.RegisterTokenServiceServer(server, Token)
	micro.RegisterMicroServiceServer(server, Micro)
	role.RegisterRoleServiceServer(server, Role)
	endpoint.RegisterEndpointServiceServer(server, Endpoint)
	policy.RegisterPolicyServiceServer(server, Policy)
	department.RegisterDepartmentServiceServer(server, Department)
	namespace.RegisterNamespaceServiceServer(server, Namespace)
	permission.RegisterPermissionServiceServer(server, Permission)
	session.RegisterAdminServiceServer(server, SessionAdmin)
	session.RegisterUserServiceServer(server, SessionUser)
	verifycode.RegisterVerifyCodeServiceServer(server, VerifyCode)
	return
}

// HTTPEntry todo
func HTTPEntry() *http.EntrySet {
	if entryInit {
		return entrySet
	}

	addServiceEntry()
	entryInit = true
	return entrySet
}

func addServiceEntry() {
	entrySet.Merge(domain.HttpEntry())
	entrySet.Merge(user.HttpEntry())
	entrySet.Merge(application.HttpEntry())
	entrySet.Merge(token.HttpEntry())
	entrySet.Merge(micro.HttpEntry())
	entrySet.Merge(role.HttpEntry())
	entrySet.Merge(endpoint.HttpEntry())
	entrySet.Merge(policy.HttpEntry())
	entrySet.Merge(department.HttpEntry())
	entrySet.Merge(namespace.HttpEntry())
	entrySet.Merge(permission.HttpEntry())
	entrySet.Merge(session.HttpEntry())
	entrySet.Merge(verifycode.HttpEntry())
}

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
	case domain.DomainServiceServer:
		if Domain != nil {
			registryError(name)
		}
		Domain = value
		addService(name, svr)
	case user.UserServiceServer:
		if User != nil {
			registryError(name)
		}
		User = value
		addService(name, svr)
	case application.UserServiceServer:
		if ApplicationUser != nil {
			registryError(name)
		}
		ApplicationUser = value
		addService(name, svr)
	case application.AdminServiceServer:
		if ApplicationAdmin != nil {
			registryError(name)
		}
		ApplicationAdmin = value
		addService(name, svr)
	case token.TokenServiceServer:
		if Token != nil {
			registryError(name)
		}
		Token = value
		addService(name, svr)
	case micro.MicroServiceServer:
		if Micro != nil {
			registryError(name)
		}
		Micro = value
		addService(name, svr)
	case role.RoleServiceServer:
		if Role != nil {
			registryError(name)
		}
		Role = value
		addService(name, svr)
	case endpoint.EndpointServiceServer:
		if Endpoint != nil {
			registryError(name)
		}
		Endpoint = value
		addService(name, svr)
	case policy.PolicyServiceServer:
		if Policy != nil {
			registryError(name)
		}
		Policy = value
		addService(name, svr)
	case department.DepartmentServiceServer:
		if Department != nil {
			registryError(name)
		}
		Department = value
		addService(name, svr)
	case namespace.NamespaceServiceServer:
		if Namespace != nil {
			registryError(name)
		}
		Namespace = value
		addService(name, svr)
	case permission.PermissionServiceServer:
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
	case session.AdminServiceServer:
		if SessionAdmin != nil {
			registryError(name)
		}
		SessionAdmin = value
		addService(name, svr)
	case session.UserServiceServer:
		if SessionUser != nil {
			registryError(name)
		}
		SessionUser = value
		addService(name, svr)
	case system.Service:
		if System != nil {
			registryError(name)
		}
		System = value
		addService(name, svr)
	case verifycode.VerifyCodeServiceServer:
		if VerifyCode != nil {
			registryError(name)
		}
		VerifyCode = value
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
