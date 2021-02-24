package client

// NewDefaultConfig todo
func NewDefaultConfig() *Config {
	return &Config{
		address:        "localhost:18050",
		Authentication: &Authentication{},
	}
}

// Config 客户端配置
type Config struct {
	address string
	*Authentication
}
