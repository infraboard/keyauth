package client

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/app/application"
	"github.com/infraboard/keyauth/app/department"
	"github.com/infraboard/keyauth/app/domain"
	"github.com/infraboard/keyauth/app/endpoint"
	"github.com/infraboard/keyauth/app/mconf"
	"github.com/infraboard/keyauth/app/micro"
	"github.com/infraboard/keyauth/app/namespace"
	"github.com/infraboard/keyauth/app/permission"
	"github.com/infraboard/keyauth/app/policy"
	"github.com/infraboard/keyauth/app/role"
	"github.com/infraboard/keyauth/app/session"
	"github.com/infraboard/keyauth/app/tag"
	"github.com/infraboard/keyauth/app/token"
	"github.com/infraboard/keyauth/app/user"
	"github.com/infraboard/keyauth/app/verifycode"
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
func (c *Client) Application() application.ServiceClient {
	return application.NewServiceClient(c.conn)
}

// Department todo
func (c *Client) Department() department.ServiceClient {
	return department.NewServiceClient(c.conn)
}

// Domain todo
func (c *Client) Domain() domain.ServiceClient {
	return domain.NewServiceClient(c.conn)
}

// Endpoint todo
func (c *Client) Endpoint() endpoint.ServiceClient {
	return endpoint.NewServiceClient(c.conn)
}

// Micro todo
func (c *Client) Micro() micro.ServiceClient {
	return micro.NewServiceClient(c.conn)
}

// Micro todo
func (c *Client) Mconf() mconf.ConfigServiceClient {
	return mconf.NewConfigServiceClient(c.conn)
}

// Namespace todo
func (c *Client) Namespace() namespace.ServiceClient {
	return namespace.NewServiceClient(c.conn)
}

// Permission todo
func (c *Client) Permission() permission.ServiceClient {
	return permission.NewServiceClient(c.conn)
}

// Policy todo
func (c *Client) Policy() policy.ServiceClient {
	return policy.NewServiceClient(c.conn)
}

// Role todo
func (c *Client) Role() role.ServiceClient {
	return role.NewServiceClient(c.conn)
}

// Tag todo
func (c *Client) Tag() tag.ServiceClient {
	return tag.NewServiceClient(c.conn)
}

// SessionAdmin todo
func (c *Client) Session() session.ServiceClient {
	return session.NewServiceClient(c.conn)
}

// Token todo
func (c *Client) Token() token.ServiceClient {
	return token.NewServiceClient(c.conn)
}

// User todo
func (c *Client) User() user.ServiceClient {
	return user.NewServiceClient(c.conn)
}

// Verifycode todo
func (c *Client) Verifycode() verifycode.ServiceClient {
	return verifycode.NewServiceClient(c.conn)
}
