//go:generate  mcube enum -m

package user

const (
	// Unknown (unknown) todo
	Unknown Gender = iota
	// Male (male) 男
	Male
	// Female (female) 女
	Female
)

// Gender 性别
type Gender uint

const (
	// DomainAdmin (domain_admin) 域管理员创建的用户
	DomainAdmin CreateType = iota
	// LDAPSync (ldap_sync) LDAP同步的用户
	LDAPSync
	// UserRegistry (user_registry) 用户自己注册的用户
	UserRegistry
)

// CreateType todo
type CreateType uint
