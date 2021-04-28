package conf

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcube/bus/broker/kafka"
	"github.com/infraboard/mcube/bus/broker/nats"
	"github.com/infraboard/mcube/cache/memory"
	"github.com/infraboard/mcube/cache/redis"
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
		Cache: newDefaultCache(),
		Nats:  nats.NewDefaultConfig(),
		Kafka: kafka.NewDefultConfig(),
	}
}

// Config 应用配置
type Config struct {
	App   *app          `toml:"app"`
	Log   *log          `toml:"log"`
	Mongo *mongodb      `toml:"mongodb"`
	Nats  *nats.Config  `toml:"nats"`
	Kafka *kafka.Config `toml:"kafka"`
	Cache *_cache       `toml:"cache"`
}

// InitGloabl 注入全局变量
func (c *Config) InitGloabl() error {
	// 加载全局配置单例
	global = c

	// 加载全局数据量单例
	mclient, err := c.Mongo.getClient()
	if err != nil {
		return err
	}
	mgoclient = mclient
	return nil
}

type app struct {
	Name     string `toml:"name" env:"K_APP_NAME"`
	Host     string `toml:"host" env:"K_APP_HOST"`
	HTTPPort string `toml:"http_port" env:"K_HTTP_PORT"`
	GRPCPort string `toml:"grpc_port" env:"K_GRPC_PORT"`
	Key      string `toml:"key" env:"K_APP_KEY"`
}

func (a *app) HTTPAddr() string {
	return a.Host + ":" + a.HTTPPort
}

func (a *app) GRPCAddr() string {
	return a.Host + ":" + a.GRPCPort
}

func newDefaultAPP() *app {
	return &app{
		Name:     "keyauth",
		Host:     "127.0.0.1",
		HTTPPort: "8050",
		GRPCPort: "18050",
		Key:      "default",
	}
}

type log struct {
	Level   string    `toml:"level" env:"K_LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"K_LOG_PATH"`
	Format  LogFormat `toml:"format" env:"K_LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"K_LOG_TO"`
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
	Endpoints []string `toml:"endpoints" env:"K_MONGO_ENDPOINTS" envSeparator:","`
	UserName  string   `toml:"username" env:"K_MONGO_USERNAME"`
	Password  string   `toml:"password" env:"K_MONGO_PASSWORD"`
	Database  string   `toml:"database" env:"K_MONGO_DATABASE"`
	AuthDB    string   `toml:"auth_db" env:"K_MONGO_AUTHDB"`
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

func (m *mongodb) authDB() string {
	if m.AuthDB != "" {
		return m.AuthDB
	}

	return m.Database
}

func (m *mongodb) GetDB() *mongo.Database {
	return m.Client().Database(m.Database)
}

func (m *mongodb) getClient() (*mongo.Client, error) {
	opts := options.Client()

	cred := options.Credential{
		AuthSource: m.authDB(),
	}

	if m.UserName != "" && m.Password != "" {
		cred.Username = m.UserName
		cred.Password = m.Password
		cred.PasswordSet = true
		opts.SetAuth(cred)
	}
	opts.SetHosts(m.Endpoints)
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

func newDefaultCache() *_cache {
	return &_cache{
		Type:   "memory",
		Memory: memory.NewDefaultConfig(),
		Redis:  redis.NewDefaultConfig(),
	}
}

type _cache struct {
	Type   string         `toml:"type" json:"type" yaml:"type" env:"K_CACHE_TYPE"`
	Memory *memory.Config `toml:"memory" json:"memory" yaml:"memory"`
	Redis  *redis.Config  `toml:"redis" json:"redis" yaml:"redis"`
}
