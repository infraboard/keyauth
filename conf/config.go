package conf

// Config 应用配置
type Config struct {
	Mongo *mongo `toml:"mongodb"`
}

type mongo struct {
	Endpoints []string `toml:"endpoints"`
	UserName  string   `toml:"username"`
	Password  string   `toml:"password"`
	Database  string   `toml:"database"`
}

func newConfig() *Config {
	return &Config{
		Mongo: new(mongo),
	}
}
