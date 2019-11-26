package conf

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mgoclient *mongo.Client
)

func newConfig() *Config {
	return &Config{
		App:   newDefaultAPP(),
		Log:   newDefaultLog(),
		Mongo: newDefaultMongoDB(),
	}
}

// Config 应用配置
type Config struct {
	App   *app     `toml:"app"`
	Log   *log     `toml:"log"`
	Mongo *mongodb `toml:"mongodb"`
}

type app struct {
	Name string `toml:"name"`
	Host string `toml:"host"`
	Port string `toml:"port"`
	Key  string `toml:"key"`
}

func (a *app) Addr() string {
	return a.Host + ":" + a.Port
}

func newDefaultAPP() *app {
	return &app{
		Name: "keyauth",
		Host: "127.0.0.1",
		Port: "8050",
		Key:  "default",
	}
}

type log struct {
	Level   string    `toml:"level"`
	PathDir string    `toml:"path_dir"`
	Format  LogFormat `toml:"format"`
	To      LogTo     `toml:"to"`
}

func newDefaultLog() *log {
	return &log{
		Level:   "debug",
		PathDir: "logs",
		Format:  "text",
		To:      "stdout",
	}
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
	if mgoclient == nil {
		panic("please load mongo client first")
	}

	return mgoclient
}

func (m *mongodb) GetDB() *mongo.Database {
	return m.Client().Database(m.Database)
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

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("new mongodb client error, %s", err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("ping mongodb server(%s) error, %s", m.Endpoints, err)
	}

	return client, nil
}
