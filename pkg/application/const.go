package application

const (
	// AdminWebApplicationName 颁发给后台管理Web端的凭证
	AdminWebApplicationName = "admin-web"
	// AdminServiceApplicationName 颁发给服务管理的应用凭证
	AdminServiceApplicationName = "admin-micro"
)

const (
	// DefaultAccessTokenExpireSecond token默认过期时长
	DefaultAccessTokenExpireSecond = 3600
	// DefaultRefreshTokenExpiredSecond 刷新token默认过期时间
	DefaultRefreshTokenExpiredSecond = DefaultAccessTokenExpireSecond * 4
)
