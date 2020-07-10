package policy

import (
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
	ID        string     `json:"id"`         // 策略ID
	CreateAt  ftime.Time `json:"create_at"`  // 创建时间
	UpdateAt  ftime.Time `json:"update_at"`  // 更新时间
	CreaterID string     `json:"creater_id"` // 创建者ID
	*CreatePolicyRequest
}

// NewPolicySet todo
func NewPolicySet(req *request.PageRequest) *PolicySet {
	return &PolicySet{
		PageRequest: req,
	}
}

// PolicySet 列表
type PolicySet struct {
	*request.PageRequest

	Total int64     `json:"total"`
	Items []*Policy `json:"items"`
}

// Add 添加
func (s *PolicySet) Add(e *Policy) {
	s.Items = append(s.Items, e)
}

// UserRoles 获取用户的角色
func (s *PolicySet) UserRoles(userID string) []string {
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
