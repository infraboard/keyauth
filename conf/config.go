package conf

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mclient *mongo.Client
	monce   sync.Once
)

func newConfig() *Config {
	return &Config{
		Mongo: newDefaultMongoDB(),
	}
}

// Config 应用配置
type Config struct {
	Mongo *mongodb `toml:"mongodb"`
}

type mongodb struct {
	Endpoints []string `toml:"endpoints"`
	UserName  string   `toml:"username"`
	Password  string   `toml:"password"`
	Database  string   `toml:"database"`
}

func newDefaultMongoDB() *mongodb {
	return &mongodb{
		Database:  "keyauth",
		Endpoints: []string{"127.0.0.1:27017"},
	}
}

// Client 获取一个全局的mongodb客户端连接
func (m *mongodb) Client() *mongo.Client {
	monce.Do(func() {
		client, err := m.getClient()
		if err != nil {
			panic(err)
		}

		mclient = client
	})

	return mclient
}

func (m *mongodb) getClient() (*mongo.Client, error) {
	cred := options.Credential{
		AuthSource: m.Database,
	}

	if m.UserName != "" && m.Password != "" {
		cred.Username = m.UserName
		cred.Password = m.Password
		cred.PasswordSet = true
	} else {
		cred.PasswordSet = false
	}
	opts := options.Client()
	opts.SetHosts(m.Endpoints)
	opts.SetAuth(cred)
	opts.SetConnectTimeout(5 * time.Second)

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, fmt.Errorf("new mongodb client error, %s", err)
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, fmt.Errorf("ping mongodb server error, %s", err)
	}

	return client, nil
}
