package system

// Service 存储服务
type Service interface {
	CreateConfig() (*Config, error)
	UpdateConfig() (*Config, error)
	GetConfig() (*Config, error)
}
