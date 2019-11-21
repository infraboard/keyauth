package application

// Application is oauth2's client: https://tools.ietf.org/html/rfc6749#section-2
type Application struct {
	ID          string `json:"id"`                   // 唯一ID
	Name        string `json:"name"`                 // 应用名称
	UserID      string `json:"user_id"`              // 应用属于那个用户
	Website     string `json:"website,omitempty"`    // 应用的网站地址
	LogoImage   string `json:"logo_image,omitempty"` // 应用的LOGO
	Description string `json:"description"`          // 应用简单的描述
	CreateAt    int64  `json:"create_at"`            // 应用创建的时间
	UpdateAt    int64  `json:"update_at"`            // 应用更新的时间

	RedirectURI       string `json:"redirect_uri"`        // 应用重定向URI, Oauht2时需要改参数
	ClientID          string `json:"client_id"`           // 应用客户端ID
	ClientSecret      string `json:"client_secret"`       // 应用客户端秘钥
	Locked            bool   `json:"locked"`              // 是否冻结应用, 冻结应用后, 该应用无法通过凭证获取访问凭证(token)
	LastLoginTime     int64  `json:"last_login_time"`     // 应用最近一次登录的时间
	LastLoginIP       string `json:"last_login_ip"`       // 最近一次登录的IP
	LoginFailedTimes  int    `json:"login_failed_times"`  // 应用最近一次连续登录失败的次数, 成功过后清零
	LoginSuccessTimes int64  `json:"login_success_times"` // 应用成功登录的次数
	TokenExpireTime   int64  `json:"token_expire_time"`   // 应用申请的token的过期时间
}
