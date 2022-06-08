package rpc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/keyauth/apps/domain"
	"github.com/infraboard/keyauth/apps/user"
	"github.com/infraboard/keyauth/client/rpc"
	mcenter "github.com/infraboard/mcenter/client/rpc"
)

func Test_Client(t *testing.T) {
	should := assert.New(t)
	conf := mcenter.NewDefaultConfig()
	c, err := rpc.NewClient(conf)

	if should.NoError(err) {
		req := user.NewQueryAccountRequest()
		req.Domain = domain.AdminDomainName
		eps, err := c.User().QueryAccount(context.Background(), req)
		if should.NoError(err) {
			t.Logf("get users: %s ", eps)
		}
	}
}
