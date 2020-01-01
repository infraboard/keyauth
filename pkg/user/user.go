package user

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// New 实例
func New(req *CreateUserRequest) (*User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	pass, err := NewHashedPassword(req.Password)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &User{
		ID:                xid.New().String(),
		CreateAt:          ftime.Now(),
		UpdateAt:          ftime.Now(),
		CreateUserRequest: req,
		HashedPassword:    pass,
	}, nil
}

// NewDescribeUser 实例
func NewDescribeUser() *User {
	return &User{
		CreateUserRequest: NewCreateUserRequest(),
	}
}

// User info
type User struct {
	ID                 string     `bson:"_id" json:"id,omitempty"`              // 用户UUID
	CreateAt           ftime.Time `bson:"create_at" json:"create_at,omitempty"` // 用户创建的时间
	UpdateAt           ftime.Time `bson:"update_at" json:"update_at,omitempty"` // 修改时间
	Primary            bool       `bson:"primary"  json:"primary"`              // 是否是主账号
	*CreateUserRequest `bson:",inline"`

	HashedPassword *Password `bson:"password" json:"password,omitempty"` // 密码相关信息
	Status         *Status   `bson:"status" json:"status,omitempty"`     // 用户状态
}

func (u *User) String() string {
	return fmt.Sprint(*u)
}

// Status 用户状态
type Status struct {
	Locked      bool       `bson:"locked" json:"locked"`                       // 是否冻结
	LockedTime  ftime.Time `bson:"locked_time" json:"locked_time,omitempty"`   // 冻结时间
	LockedReson string     `bson:"locked_reson" json:"locked_reson,omitempty"` // 冻结原因
	UnLockTime  ftime.Time `bson:"unlock_time" json:"unlock_time,omitempty"`   // 解冻时间
}

// NewCreateUserRequest 创建请求
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{}
}

// Validate 校验请求是否合法
func (req *CreateUserRequest) Validate() error {
	return validate.Struct(req)
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Account     string `bson:"account" json:"account,omitempty" validate:"required,lte=60"` // 用户账号名称
	Mobile      string `bson:"mobile" json:"mobile,omitempty" validate:"lte=30"`            // 手机号码, 用户可以通过手机进行注册和密码找回, 还可以通过手机号进行登录
	Email       string `bson:"email" json:"email,omitempty" validate:"lte=30"`              // 邮箱, 用户可以通过邮箱进行注册和照明密码
	Phone       string `bson:"phone" json:"phone,omitempty" validate:"lte=30"`              // 用户的座机号码
	Address     string `bson:"address" json:"address,omitempty" validate:"lte=120"`         // 用户住址
	RealName    string `bson:"real_name" json:"real_name,omitempty" validate:"lte=10"`      // 用户真实姓名
	NickName    string `bson:"nick_name" json:"nick_name,omitempty" validate:"lte=30"`      // 用户昵称, 用于在界面进行展示
	Gender      string `bson:"gender" json:"gender,omitempty" validate:"lte=10"`            // 性别
	Avatar      string `bson:"avatar" json:"avatar,omitempty" validate:"lte=300"`           // 头像
	Language    string `bson:"language" json:"language,omitempty" validate:"lte=40"`        // 用户使用的语言
	City        string `bson:"city" json:"city,omitempty" validate:"lte=40"`                // 用户所在的城市
	Province    string `bson:"province" json:"province,omitempty" validate:"lte=40"`        // 用户所在的省
	ExpiresDays int    `bson:"expires_days" json:"expires_days,omitempty"`                  // 用户多久未登录时(天), 冻结改用户, 防止僵尸用户的账号被利用
	Password    string `bson:"-" json:"password,omitempty" validate:"required,lte=80"`      // 密码相关信息
}

// NewHashedPassword 生产hash后的密码对象
func NewHashedPassword(password string) (*Password, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	return &Password{
		Password: string(bytes),
		CreateAt: ftime.Now(),
		UpdateAt: ftime.Now(),
	}, nil
}

// Password user's password
type Password struct {
	Password string     `bson:"password" json:"password,omitempty"`    // hash过后的密码
	ExpireAt ftime.Time `bson:"expire_at" json:"expire_at,omitempty" ` // 密码过期时间
	CreateAt ftime.Time `bson:"create_at" json:"create_at,omitempty" ` // 密码创建时间
	UpdateAt ftime.Time `bson:"update_at" json:"update_at,omitempty"`  // 密码更新时间
}

// CheckPassword 判断password 是否正确
func (p *Password) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
}
