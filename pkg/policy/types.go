//go:generate  mcube enum -m

package policy

// Type 更新模式
type Type int32

const (
	// CustomPolicy (custom) 用户自己定义的策略
	CustomPolicy Type = iota
	// BuildInPolicy (build_in) 系统内部逻辑, 不允许用户看到并修改
	BuildInPolicy
)
