package user

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
	"golang.org/x/crypto/bcrypt"

	common "github.com/infraboard/keyauth/common/types"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/token/session"
	"github.com/infraboard/keyauth/pkg/user/types"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// New 实例
func New(req *CreateAccountRequest) (*User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	pass, err := NewHashedPassword(req.Password)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &User{
		CreateAt:             ftime.Now(),
		UpdateAt:             ftime.Now(),
		CreateAccountRequest: req,
		HashedPassword:       pass,
		Status: &Status{
			Locked: false,
		},
	}, nil
}

// NewDefaultUser 实例
func NewDefaultUser() *User {
	return &User{
		CreateAccountRequest: NewCreateUserRequest(),
		Status: &Status{
			Locked: false,
		},
	}
}

// User info
type User struct {
	CreateAt              ftime.Time     `bson:"create_at" json:"create_at,omitempty"` // 用户创建的时间
	UpdateAt              ftime.Time     `bson:"update_at" json:"update_at,omitempty"` // 修改时间
	Domain                string         `bson:"domain" json:"domain,omitempty"`       // 如果是子账号和服务账号 都需要继承主用户Domain
	Type                  types.UserType `bson:"type"  json:"type"`                    // 是否是主账号
	Roles                 []string       `bson:"-" json:"roles,omitempty"`             // 用户的角色(当携带Namesapce查询时会有)
	*CreateAccountRequest `bson:",inline"`

	HashedPassword *Password              `bson:"password" json:"password,omitempty"` // 密码相关信息
	Status         *Status                `bson:"status" json:"status,omitempty"`     // 用户状态
	Department     *department.Department `bson:"-" json:"department,omitempty"`      // 部门
}

// Block 锁用户
func (u *User) Block(reason string) {
	u.Status.Locked = true
	u.Status.LockedReson = reason
	u.Status.LockedTime = ftime.Now()
}

// Desensitize 关键数据脱敏
func (u *User) Desensitize() {
	if u.HashedPassword != nil {
		u.HashedPassword.Password = ""
		u.HashedPassword.History = []string{}
	}
	return
}

// ChangePassword 修改用户密码
func (u *User) ChangePassword(old, new string, maxHistory uint, needReset bool) error {
	// 确认旧密码
	if err := u.HashedPassword.CheckPassword(old); err != nil {
		return err
	}

	// 修改新密码
	newPass, err := NewHashedPassword(new)
	if err != nil {
		return exception.NewBadRequest(err.Error())
	}
	u.HashedPassword.Update(newPass, maxHistory, needReset)
	return nil
}

// CreateAccountRequest 创建用户请求
type CreateAccountRequest struct {
	*session.Session `bson:"-" json:"-"`
	*Profile         `bson:",inline"`
	CreateType       CreateType `bson:"create_type" json:"create_type"`               // 创建方式
	Password         string     `bson:"-" json:"password" validate:"required,lte=80"` // 密码相关信息
}

// NewProfile todo
func NewProfile() *Profile {
	return &Profile{}
}

// Profile todo
type Profile struct {
	DepartmentID  string `bson:"department_id" json:"department_id" validate:"lte=200"` // 用户所属部门
	Account       string `bson:"_id" json:"account" validate:"required,lte=60"`         // 用户账号名称
	Phone         string `bson:"phone" json:"phone" validate:"lte=30"`                  // 手机号码, 用户可以通过手机进行注册和密码找回, 还可以通过手机号进行登录
	Email         string `bson:"email" json:"email" validate:"lte=30"`                  // 邮箱, 用户可以通过邮箱进行注册和照明密码
	Address       string `bson:"address" json:"address" validate:"lte=120"`             // 用户住址
	RealName      string `bson:"real_name" json:"real_name" validate:"lte=10"`          // 用户真实姓名
	NickName      string `bson:"nick_name" json:"nick_name" validate:"lte=30"`          // 用户昵称, 用于在界面进行展示
	Gender        Gender `bson:"gender" json:"gender" validate:"lte=10"`                // 性别
	Avatar        string `bson:"avatar" json:"avatar" validate:"lte=300"`               // 头像
	Language      string `bson:"language" json:"language" validate:"lte=40"`            // 用户使用的语言
	City          string `bson:"city" json:"city" validate:"lte=40"`                    // 用户所在的城市
	Province      string `bson:"province" json:"province" validate:"lte=40"`            // 用户所在的省
	ExpiresDays   int    `bson:"expires_days" json:"expires_days"`                      // 用户多久未登录时(天), 冻结改用户, 防止僵尸用户的账号被利用'
	IsInitialized bool   `bson:"is_initialized" json:"is_initialized"`                  // 用户是否初始化
}

// ValidateInitialized 判断初始化数据是否准备好了
func (req *Profile) ValidateInitialized() error {
	if req.Email != "" && req.Phone != "" {
		return nil
	}

	return fmt.Errorf("email and phone required when initial")
}

// HasDepartment todo
func (req *Profile) HasDepartment() bool {
	return req.DepartmentID != ""
}

// Patch todo
func (req *Profile) Patch(data *Profile) {
	patchData, _ := json.Marshal(data)
	json.Unmarshal(patchData, req)
}

// Validate 校验请求是否合法
func (req *CreateAccountRequest) Validate() error {
	tk := req.GetToken()

	if tk == nil {
		return fmt.Errorf("token required")
	}

	// 非管理员, 主账号 可以创建子账号
	if !tk.UserType.Is(types.UserType_SUPPER, types.UserType_PRIMARY) {
		return fmt.Errorf("%s user can't create sub account", tk.UserType)
	}

	return validate.Struct(req)
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
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
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
	Password    string     `bson:"password" json:"password,omitempty"`    // hash过后的密码
	CreateAt    ftime.Time `bson:"create_at" json:"create_at,omitempty" ` // 密码创建时间
	UpdateAt    ftime.Time `bson:"update_at" json:"update_at,omitempty"`  // 密码更新时间
	NeedReset   bool       `bson:"need_reset" json:"need_reset"`          // 密码需要被重置
	ResetReason string     `bson:"reset_reason" json:"reset_reason"`      // 需要重置的原因
	History     []string   `bson:"history" json:"history,omitempty"`      // 历史密码

	IsExpired bool `bson:"-" json:"is_expired"` // 是否过期
}

// SetExpired 密码过期
func (p *Password) SetExpired() {
	p.IsExpired = true
}

// SetNeedReset 需要被重置
func (p *Password) SetNeedReset(format string, a ...interface{}) {
	p.NeedReset = true
	p.ResetReason = fmt.Sprintf(format, a...)
}

// CheckPassword 判断password 是否正确
func (p *Password) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
	if err != nil {
		return exception.NewUnauthorized("user or password not connrect")
	}
	return nil
}

// IsHistory 检测是否是历史密码
func (p *Password) IsHistory(password string) bool {
	for _, pass := range p.History {
		err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
		if err == nil {
			return true
		}
	}

	return false
}

// HistoryCount 保存了几个历史密码
func (p *Password) HistoryCount() int {
	return len(p.History)
}

func (p *Password) rotaryHistory(maxHistory uint) {
	if uint(p.HistoryCount()) < maxHistory {
		p.History = append(p.History, p.Password)
	} else {
		remainHistry := p.History[:maxHistory]
		p.History = []string{p.Password}
		p.History = append(p.History, remainHistry...)
	}
}

// Update 更新密码
func (p *Password) Update(new *Password, maxHistory uint, needReset bool) {
	p.rotaryHistory(maxHistory)
	p.Password = new.Password
	p.NeedReset = needReset
	p.UpdateAt = ftime.Now()
	if !needReset {
		p.ResetReason = ""
	}
}

// NewUserSet 实例
func NewUserSet(req *request.PageRequest) *Set {
	return &Set{
		PageRequest: req,
		Items:       []*User{},
	}
}

// Set 用户列表
type Set struct {
	*request.PageRequest

	Total int64   `json:"total"`
	Items []*User `json:"items"`
}

// Add todo
func (s *Set) Add(u *User) {
	s.Items = append(s.Items, u)
}

// NewPutAccountRequest todo
func NewPutAccountRequest() *UpdateAccountRequest {
	return &UpdateAccountRequest{
		Session:    session.NewSession(),
		UpdateMode: common.PutUpdateMode,
		Profile:    NewProfile(),
	}
}

// NewPatchAccountRequest todo
func NewPatchAccountRequest() *UpdateAccountRequest {
	return &UpdateAccountRequest{
		Session:    session.NewSession(),
		UpdateMode: common.PatchUpdateMode,
		Profile:    NewProfile(),
	}
}

// UpdateAccountRequest todo
type UpdateAccountRequest struct {
	*session.Session `bson:"-" json:"-"`
	UpdateMode       common.UpdateMode `json:"update_mode"`
	*Profile         `bson:",inline"`
}

// Validate 更新请求校验
func (req *UpdateAccountRequest) Validate() error {
	tk := req.GetToken()

	if tk == nil {
		return fmt.Errorf("token required")
	}

	// 用户初始化要判断初始化信息填写完整
	if err := req.ValidateInitialized(); req.IsInitialized && err != nil {
		return err
	}

	return validate.Struct(req)
}
