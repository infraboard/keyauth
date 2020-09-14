package provider

import (
	"fmt"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

// NewLDAPConfig todo
func NewLDAPConfig(req *SaveLDAPConfigRequest) (*LDAPConfig, error) {
	tk := req.GetToken()
	if tk == nil {
		return nil, fmt.Errorf("token requird")
	}

	if err := req.Config.Validate(); err != nil {
		return nil, err
	}

	ins := &LDAPConfig{
		ID:                    xid.New().String(),
		Domain:                tk.Domain,
		Creater:               tk.Account,
		CreateAt:              ftime.Now(),
		UpdateAt:              ftime.Now(),
		SaveLDAPConfigRequest: req,
	}
	return ins, nil
}

// NewDefaultLDAPConfig todo
func NewDefaultLDAPConfig() *LDAPConfig {
	return &LDAPConfig{
		SaveLDAPConfigRequest: NewSaveLDAPConfigRequest(),
	}
}

// LDAPConfig todo
type LDAPConfig struct {
	ID                     string     `bson:"_id" json:"id,omitempty"`              // 唯一ID
	Domain                 string     `bson:"domain" json:"domain_id,omitempty"`    // 所属域ID
	Creater                string     `bson:"creater" json:"creater,omitempty"`     // 创建人
	CreateAt               ftime.Time `bson:"create_at" json:"create_at,omitempty"` // 创建时间
	UpdateAt               ftime.Time `bson:"update_at" json:"update_at,omitempty"` // 更新时间
	*SaveLDAPConfigRequest `bson:",inline"`
}

// NewLDAPSet 实例化
func NewLDAPSet(req *request.PageRequest) *LDAPSet {
	return &LDAPSet{
		PageRequest: req,
		Items:       []*LDAPConfig{},
	}
}

// LDAPSet 列表
type LDAPSet struct {
	*request.PageRequest

	Total int64         `json:"total"`
	Items []*LDAPConfig `json:"items"`
}

// Add 添加应用
func (s *LDAPSet) Add(item *LDAPConfig) {
	s.Items = append(s.Items, item)
}
