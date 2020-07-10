package user

const (
	// SupperAccount 超级管理员
	SupperAccount Type = "supper"
	// ServiceAccount 服务账号
	ServiceAccount = "service"
	// PrimaryAccount 主账号
	PrimaryAccount = "primary"
	// SubAccount 子账号
	SubAccount = "sub"
)

// Type 用户类型
type Type string

// Is todo
func (t *Type) Is(tps ...Type) bool {
	for _, tp := range tps {
		if *t == tp {
			return true
		}
	}

	return false
}
