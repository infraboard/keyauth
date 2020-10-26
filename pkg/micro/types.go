//go:generate  mcube enum -m

package micro

import (
	"github.com/infraboard/keyauth/version"
)

var (
	systemMicro = []string{version.ServiceName}
)

// IsSystemMicro todo
func IsSystemMicro(name string) bool {
	for _, sm := range systemMicro {
		if name == sm {
			return true
		}
	}

	return false
}

const (
	// Custom (cumstom) 自定义的服务
	Custom Type = iota
	// BuildIn (build_in) 系统内建的服务
	BuildIn
)

// Type 服务类型
type Type uint
