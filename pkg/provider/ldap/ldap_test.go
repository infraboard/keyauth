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

func TestGetBaseDNFromUser(t *testing.T) {
	should := assert.New(t)

	conf := ldap.NewDefaultConfig()
	conf.User = "cn=admin,dc=example,dc=org"
	baseDN := conf.GetBaseDNFromUser()

	should.Equal("dc=example,dc=org", baseDN)
}

func init() {
	zap.DevelopmentSetup()
}
