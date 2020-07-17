package ldap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/keyauth/pkg/provider/ldap"
	"github.com/infraboard/mcube/logger/zap"
)

func TestConn(t *testing.T) {
	should := assert.New(t)

	conf := ldap.NewDefaultConfig()

	p := ldap.NewProvider(conf)
	ok, err := p.CheckUserPassword("admin", "admin")
	should.NoError(err)
	should.True(ok)
}

func init() {
	zap.DevelopmentSetup()
}
