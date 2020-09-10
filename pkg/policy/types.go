//go:generate  mcube enum -m

package policy

// Type 更新模式
type Type uint

const (
	// BuildInPolicy (build_in) 系统内部逻辑, 不允许用户看到并修改
	BuildInPolicy Type = iota
	// CustomPolicy (custom)) 用户自己定义的策略
	CustomPolicy
)
