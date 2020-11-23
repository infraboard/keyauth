package system

// Service 存储服务
type Service interface {
	CreateConfig() (*Config, error)
	UpdateDomain() (*Config, error)
	DescriptionDomain() (*Config, error)
}
