package user

// PrimaryAccountService 主账号服务
type PrimaryAccountService interface {
	// 新建主账号
	CreatePrimayAccount(req *CreateUserRequest) (*User, error)
	// 更新用户密码
	UpdateAccountPassword(userName, oldPass, newPass string) error
}

// RAMAccountService 子账号服务
type RAMAccountService interface {
	ListRAMAccount(domainID string) ([]*User, error)
	// CreateRAMAccount 创建子账号
	// RAM (Resource Access Management)提供的用户身份管理与访问控制服务
	CreateRAMAccount(domainID string, req *CreateUserRequest) (*User, error)
	// 注销子账号
	DeleteRAMAccount(userID string) error
}
