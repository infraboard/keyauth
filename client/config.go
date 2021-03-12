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

// SetAddress todo
func (c *Config) SetAddress(addr string) {
	c.address = addr
}

// Address 地址
func (c *Config) Address() string {
	return c.address
}
