package provider

import (
	"fmt"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/pkg/provider/ldap"
	"github.com/infraboard/keyauth/pkg/token"
)

// NewLDAPConfig todo
func NewLDAPConfig(tk *token.Token, conf *ldap.Config) (*LDAPConfig, error) {
	if tk == nil {
		return nil, fmt.Errorf("token requird")
	}

	if err := conf.Validate(); err != nil {
		return nil, err
	}
	ins := &LDAPConfig{
		ID:       xid.New().String(),
		Domain:   tk.Domain,
		Creater:  tk.Account,
		CreateAt: ftime.Now(),
		UpdateAt: ftime.Now(),
		Config:   conf,
	}
	return ins, nil
}

// NewDefaultLDAPConfig todo
func NewDefaultLDAPConfig() *LDAPConfig {
	return &LDAPConfig{
		Config: ldap.NewDefaultConfig(),
	}
}

// LDAPConfig todo
type LDAPConfig struct {
	ID           string     `bson:"_id" json:"id,omitempty"`              // 唯一ID
	Domain       string     `bson:"domain" json:"domain_id,omitempty"`    // 所属域ID
	Creater      string     `bson:"creater" json:"creater,omitempty"`     // 创建人
	CreateAt     ftime.Time `bson:"create_at" json:"create_at,omitempty"` // 创建时间
	UpdateAt     ftime.Time `bson:"update_at" json:"update_at,omitempty"` // 更新时间
	*ldap.Config `bson:",inline"`
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
