//go:generate  mcube enum -m

package audit

// Result todo
type Result uint

const (
	// Success (success) todo
	Success Result = iota
	// Failed (failed) todo
	Failed
)

// ActionType 动作类型
type ActionType uint

const (
	// LoginAction (login) 登录动作
	LoginAction ActionType = iota
	// LogoutAction (logout) 登出
	LogoutAction
)
