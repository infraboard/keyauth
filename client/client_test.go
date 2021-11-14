package client_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/keyauth/app/domain"
	"github.com/infraboard/keyauth/app/user"
	"github.com/infraboard/keyauth/client"
)

func Test_Client(t *testing.T) {
	should := assert.New(t)
	conf := client.NewDefaultConfig()
	conf.SetClientCredentials("pz3HiVQA3indzSHzFKtLHaJW", "vDvlAtqN3rS9CZcHugXp6QBuk28zRjud")
	c, err := client.NewClient(conf)

	if should.NoError(err) {
		req := user.NewQueryAccountRequest()
		req.Domain = domain.AdminDomainName
		eps, err := c.User().QueryAccount(context.Background(), req)
		if should.NoError(err) {
			t.Logf("get users: %s ", eps)
		}
	}
}
