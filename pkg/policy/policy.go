package policy

import (
	"crypto/sha1"
	"fmt"

	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
)

// New 新实例
func New(createrID string, req *CreatePolicyRequest) (*Policy, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &Policy{
		ID:                  req.hashedID(),
		CreateAt:            ftime.Now(),
		UpdateAt:            ftime.Now(),
		CreaterID:           createrID,
		CreatePolicyRequest: req,
	}, nil
}

// NewDefaultPolicy todo
func NewDefaultPolicy() *Policy {
	return &Policy{
		CreatePolicyRequest: NewCreatePolicyRequest(),
	}
}

// Policy 权限策略
type Policy struct {
	ID        string     `bson:"_id" json:"id"`                // 策略ID
	CreateAt  ftime.Time `bson:"create_at" json:"create_at"`   // 创建时间
	UpdateAt  ftime.Time `bson:"update_at" json:"update_at"`   // 更新时间
	CreaterID string     `bson:"creater_id" json:"creater_id"` // 创建者ID
	UserType  user.Type  `bson:"user_type" json:"user_type"`   // 用户类型

	*CreatePolicyRequest `bson:",inline"`
}

// NewCreatePolicyRequest 请求实例
func NewCreatePolicyRequest() *CreatePolicyRequest {
	return &CreatePolicyRequest{}
}

// CreatePolicyRequest 创建策略的请求
type CreatePolicyRequest struct {
	UserID      string      `bson:"user_id" json:"user_id" validate:"required,lte=120"`    // 用户ID
	RoleName    string      `bson:"role_name" json:"role_name" validate:"required,lte=40"` // 角色名称
	ExpiredTime *ftime.Time `bson:"expired_time" json:"expired_time"`                      // 策略过期时间
	Namespace   string      `bson:"namespace" json:"namespace" validate:"lte=120"`         // 范围
}

// Validate 校验请求合法
func (req *CreatePolicyRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreatePolicyRequest) hashedID() string {
	inst := sha1.New()
	hashedStr := fmt.Sprintf("%s-%s-%s",
		req.Namespace, req.UserID, req.RoleName)
	inst.Write([]byte(hashedStr))
	return fmt.Sprintf("%x", inst.Sum([]byte("")))
}

// NewPolicySet todo
func NewPolicySet(req *request.PageRequest) *Set {
	return &Set{
		PageRequest: req,
	}
}

// Set 列表
type Set struct {
	*request.PageRequest

	Total int64     `json:"total"`
	Items []*Policy `json:"items"`
}

// Add 添加
func (s *Set) Add(e *Policy) {
	s.Items = append(s.Items, e)
}

// UserRoles 获取用户的角色
func (s *Set) UserRoles(userID string) []string {
	rns := []string{}
	for i := range s.Items {
		item := s.Items[i]
		if item.UserID == userID {
			rns = append(rns, item.RoleName)
		}
	}

	if len(rns) == 0 {
		rns = append(rns, "vistor")
	}

	return rns
}
