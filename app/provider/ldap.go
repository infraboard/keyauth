package provider

import (
	"encoding/json"
	"fmt"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
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

	// 补充默认BaseDN
	if req.BaseDN == "" {
		req.BaseDN = req.GetBaseDNFromUser()
	}

	ins := &LDAPConfig{
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
	Domain                 string     `bson:"_id" json:"domain,omitempty"`          // 所属域ID
	Creater                string     `bson:"creater" json:"creater,omitempty"`     // 创建人
	CreateAt               ftime.Time `bson:"create_at" json:"create_at,omitempty"` // 创建时间
	UpdateAt               ftime.Time `bson:"update_at" json:"update_at,omitempty"` // 更新时间
	*SaveLDAPConfigRequest `bson:",inline"`
}

// Merge todo
func (ldap *LDAPConfig) Merge(data *LDAPConfig) {
	mergeData, _ := json.Marshal(data)
	json.Unmarshal(mergeData, ldap)
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
