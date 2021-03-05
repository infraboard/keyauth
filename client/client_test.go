package client_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/mcube/http/request"
)

func Test_Client(t *testing.T) {
	should := assert.New(t)
	conf := client.NewDefaultConfig()
	conf.SetClientCredentials("VYizVq1fsK7olinqVHrBvFOl", "qS9FGBoFGRaVfbgeqFVDRcgH7nNJi9fp")
	c, err := client.NewClient(conf)
	if should.NoError(err) {
		page := request.NewPageRequest(20, 1)
		meta := metadata.Pairs("access_token", "NEjvVOhmhAQXFuYSrZdJaBsH")
		ctx := metadata.NewOutgoingContext(context.Background(), meta)
		eps, err := c.Endpoint().QueryEndpoints(ctx, endpoint.NewQueryEndpointRequest(page))
		if should.NoError(err) {
			t.Logf("get eps: %s ", eps)
		}
	}
}
