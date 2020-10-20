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
