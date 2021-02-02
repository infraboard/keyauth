//go:generate  mcube enum -m

package geoip

import (
	"fmt"
	"strings"
)

// DBFileContentType 数据文件内容类型
type DBFileContentType uint

const (
	// IPv4Content (ipv4) ipv4 文件
	IPv4Content DBFileContentType = iota
	// LocationContent (location) ip 地域信息
	LocationContent
)

// ParseDBFileContentType Parse DBFileContentType from string
func ParseDBFileContentType(str string) (DBFileContentType, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := DBFileContentType_value[key]
	if !ok {
		return 0, fmt.Errorf("unknown Status: %s", str)
	}

	return v, nil
}
