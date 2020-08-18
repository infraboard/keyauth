//go:generate  mcube enum -m

package geoip

// DBFileContentType 数据文件内容类型
type DBFileContentType uint

const (
	// IPv4Content (ipv4) ipv4 文件
	IPv4Content DBFileContentType = iota
	// LocationContent (location) ip 地域信息
	LocationContent
)
