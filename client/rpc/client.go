package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcenter/client/rpc/auth"
	"github.com/infraboard/mcenter/client/rpc/resolver"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/infraboard/keyauth/apps/application"
	"github.com/infraboard/keyauth/apps/department"
	"github.com/infraboard/keyauth/apps/domain"
	"github.com/infraboard/keyauth/apps/endpoint"
	"github.com/infraboard/keyauth/apps/mconf"
	"github.com/infraboard/keyauth/apps/micro"
	"github.com/infraboard/keyauth/apps/namespace"
	"github.com/infraboard/keyauth/apps/permission"
	"github.com/infraboard/keyauth/apps/policy"
	"github.com/infraboard/keyauth/apps/role"
	"github.com/infraboard/keyauth/apps/session"
	"github.com/infraboard/keyauth/apps/tag"
	"github.com/infraboard/keyauth/apps/token"
	"github.com/infraboard/keyauth/apps/user"
	"github.com/infraboard/keyauth/apps/verifycode"
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
func NewClient(conf *rpc.Config) (*Client, error) {
	zap.DevelopmentSetup()
	log := zap.L()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 连接到服务
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s://%s", resolver.Scheme, "keyauth"), // Dial to "mcenter://keyauth"
		grpc.WithPerRPCCredentials(auth.NewAuthentication(conf.ClientID, conf.ClientSecret)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithBlock(),
	)

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
	conf *rpc.Config
	conn *grpc.ClientConn
	log  logger.Logger
}

// GetClientID todo
func (c *Client) GetClientID() string {
	return c.conf.ClientID
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
func (c *Client) Mconf() mconf.ServiceClient {
	return mconf.NewServiceClient(c.conn)
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
