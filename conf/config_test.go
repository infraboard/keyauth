package conf_test

import (
	"testing"

	"github.com/infraboard/keyauth/conf"
	"github.com/stretchr/testify/require"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := require.New(t)

	err := conf.LoadConfigFromToml("../etc/keyauth.toml")
	should.NoError(err)

	t.Log(conf.C().Mongo.Endpoints)
}

func TestMongoClient(t *testing.T) {
	should := require.New(t)

	err := conf.LoadConfigFromToml("../etc/keyauth.toml")
	should.NoError(err)

	conf.C().Mongo.Client()
}
