package audit

// LoginRecord 登录日志
type LoginRecord struct {
	LastLoginTime     int64  `json:"last_login_time"`     // 应用最近一次登录的时间
	LastLoginIP       string `json:"last_login_ip"`       // 最近一次登录的IP
	LoginFailedTimes  int    `json:"login_failed_times"`  // 应用最近一次连续登录失败的次数, 成功过后清零
	LoginSuccessTimes int64  `json:"login_success_times"` // 应用成功登录的次数
}
