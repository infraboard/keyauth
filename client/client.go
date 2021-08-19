package client

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/mconf"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/keyauth/pkg/tag"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

var (
	client *Client
)

// SetGlobal todo
func SetGlobal(cli *Client) {
	client = cli
}

// C Global
func C() *Client {
	return client
}

// NewClient todo
func NewClient(conf *Config) (*Client, error) {
	zap.DevelopmentSetup()
	log := zap.L()

	conn, err := grpc.Dial(conf.address, grpc.WithInsecure(), grpc.WithPerRPCCredentials(conf.Authentication))
	if err != nil {
		return nil, err
	}

	return &Client{
		conf: conf,
		conn: conn,
		log:  log,
	}, nil
}

// Client 客户端
type Client struct {
	conf *Config
	conn *grpc.ClientConn
	log  logger.Logger
}

// GetClientID todo
func (c *Client) GetClientID() string {
	return c.conf.clientID
}

// ApplicationAdmin todo
func (c *Client) Application() application.ApplicationServiceClient {
	return application.NewApplicationServiceClient(c.conn)
}

// Department todo
func (c *Client) Department() department.DepartmentServiceClient {
	return department.NewDepartmentServiceClient(c.conn)
}

// Domain todo
func (c *Client) Domain() domain.DomainServiceClient {
	return domain.NewDomainServiceClient(c.conn)
}

// Endpoint todo
func (c *Client) Endpoint() endpoint.EndpointServiceClient {
	return endpoint.NewEndpointServiceClient(c.conn)
}

// Micro todo
func (c *Client) Micro() micro.MicroServiceClient {
	return micro.NewMicroServiceClient(c.conn)
}

// Micro todo
func (c *Client) Mconf() mconf.MicroConfigServiceClient {
	return mconf.NewMicroConfigServiceClient(c.conn)
}

// Namespace todo
func (c *Client) Namespace() namespace.NamespaceServiceClient {
	return namespace.NewNamespaceServiceClient(c.conn)
}

// Permission todo
func (c *Client) Permission() permission.PermissionServiceClient {
	return permission.NewPermissionServiceClient(c.conn)
}

// Policy todo
func (c *Client) Policy() policy.PolicyServiceClient {
	return policy.NewPolicyServiceClient(c.conn)
}

// Role todo
func (c *Client) Role() role.RoleServiceClient {
	return role.NewRoleServiceClient(c.conn)
}

// Tag todo
func (c *Client) Tag() tag.TagServiceClient {
	return tag.NewTagServiceClient(c.conn)
}

// SessionAdmin todo
func (c *Client) SessionAdmin() session.AdminServiceClient {
	return session.NewAdminServiceClient(c.conn)
}

// SessionUser todo
func (c *Client) SessionUser() session.UserServiceClient {
	return session.NewUserServiceClient(c.conn)
}

// Token todo
func (c *Client) Token() token.TokenServiceClient {
	return token.NewTokenServiceClient(c.conn)
}

// User todo
func (c *Client) User() user.UserServiceClient {
	return user.NewUserServiceClient(c.conn)
}

// Verifycode todo
func (c *Client) Verifycode() verifycode.VerifyCodeServiceClient {
	return verifycode.NewVerifyCodeServiceClient(c.conn)
}
