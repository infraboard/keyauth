package user

import "fmt"

// User info
type User struct {
	ID                string    `json:"id,omitempty"`                  // 用户UUID
	Account           string    `json:"account,omitempty"`             // 用户账号名称
	Mobile            string    `json:"mobile,omitempty"`              // 手机号码, 用户可以通过手机进行注册和密码找回, 还可以通过手机号进行登录
	Email             string    `json:"email,omitempty"`               // 邮箱, 用户可以通过邮箱进行注册和照明密码
	Phone             string    `json:"phone,omitempty"`               // 用户的座机号码
	Address           string    `json:"address,omitempty"`             // 用户住址
	RealName          string    `json:"real_name,omitempty"`           // 用户真实姓名
	NickName          string    `json:"nick_name,omitempty"`           // 用户昵称, 用于在界面进行展示
	Gender            string    `json:"gender,omitempty"`              // 性别
	Avatar            string    `json:"avatar,omitempty"`              // 头像
	Language          string    `json:"language,omitempty"`            // 用户使用的语言
	City              string    `json:"city,omitempty"`                // 用户所在的城市
	Province          string    `json:"province,omitempty"`            // 用户所在的省
	Locked            string    `json:"locked,omitempty"`              // 是否冻结次用户
	CreateAt          int64     `json:"create_at,omitempty"`           // 用户创建的时间
	ExpiresActiveDays int       `json:"expires_active_days,omitempty"` // 用户多久未登录时(天), 冻结改用户, 防止僵尸用户的账号被利用
	Password          *Password `json:"password,omitempty"`            // 密码相关信息
	IsDomainOwner     bool      `json:"is_domain_owner,omitempty"`
}

// Password user's password
type Password struct {
	UserID   string `json:"-"`
	Password string `json:"-"`
	ExpireAt int64  `json:"expire_at"`           // 密码过期时间
	CreateAt int64  `json:"create_at"`           // 密码创建时间
	UpdateAt int64  `json:"update_at,omitempty"` // 密码更新时间
}

func (u *User) String() string {
	return fmt.Sprint(*u)
}
