package user

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

// Type 用户类型
type Type string

const (
	// SupperAdmin 超级管理员
	SupperAdmin Type = "supper"
	// PrimaryAccount 主账号
	PrimaryAccount = "primary"
	// SubAccount 子账号
	SubAccount = "sub"
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
	Type               Type       `bson:"type"  json:"type"`                    // 是否是主账号
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
