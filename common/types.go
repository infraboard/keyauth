//go:generate  mcube enum -m

package common

// UpdateMode 更新模式
type UpdateMode uint

const (
	// PutUpdateMode (put) Patch 模式
	PutUpdateMode UpdateMode = iota
	// PatchUpdateMode (patch) Put 模式
	PatchUpdateMode
)
