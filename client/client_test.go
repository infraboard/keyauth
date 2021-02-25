package client_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/mcube/http/request"
)

func Test_Client(t *testing.T) {
	should := assert.New(t)
	conf := client.NewDefaultConfig()
	conf.SetPasswordCredentials("VYizVq1fsK7olinqVHrBvFOl", "qS9FGBoFGRaVfbgeqFVDRcgH7nNJi9fp")
	c, err := client.NewClient(conf)
	if should.NoError(err) {
		page := request.NewPageRequest(20, 1)
		eps, err := c.Endpoint().QueryEndpoints(context.Background(), endpoint.NewQueryEndpointRequest(page))
		if should.NoError(err) {
			t.Logf("get eps: %s", eps)
		}
	}
}
