package application

// Service 用户服务
type Service interface {
	// 创建app
	CreateUserApplication(userID string, req *CreateApplicatonRequest) (*Application, error)
	// 删除app
	DeleteApplication(id string) error
}
