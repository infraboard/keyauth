package client_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/keyauth/app/endpoint"
	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/mcube/http/request"
)

func Test_Client(t *testing.T) {
	should := assert.New(t)
	conf := client.NewDefaultConfig()
	conf.SetClientCredentials("pz3HiVQA3indzSHzFKtLHaJW", "vDvlAtqN3rS9CZcHugXp6QBuk28zRjud")
	c, err := client.NewClient(conf)
	if should.NoError(err) {
		page := request.NewPageRequest(20, 1)
		eps, err := c.Endpoint().QueryEndpoints(context.Background(), endpoint.NewQueryEndpointRequest(page))
		if should.NoError(err) {
			t.Logf("get eps: %s ", eps)
		}
	}
}
