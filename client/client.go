package client

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/pkg/endpoint"
)

// NewClient todo
func NewClient(conf *Config) (*Client, error) {
	zap.DevelopmentSetup()
	log := zap.L()

	conn, err := grpc.Dial(conf.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
		log:  log,
	}, nil
}

// Client 客户端
type Client struct {
	conn *grpc.ClientConn
	log  logger.Logger
}

// Endpoint todo
func (c *Client) Endpoint() endpoint.EndpointServiceClient {
	return endpoint.NewEndpointServiceClient(c.conn)
}
