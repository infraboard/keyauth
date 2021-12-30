package wxwork

import (
"fmt"
"github.com/infraboard/keyauth/app/provider/auth/wxwork"
"github.com/infraboard/keyauth/app/token/session"
"github.com/infraboard/mcube/http/request"
"github.com/infraboard/mcube/types/ftime"
)

type WechatWork interface {
	SaveConfig(*SaveConfRequest) (*WechatWorkConfig, error)
	QueryConfig(*QueryConfigRequest) (*WechatWorkSet, error)
	DescribeConfig(*DescribeWechatWorkConf) (*WechatWorkConfig, error)
	DeleteConfig(*DescribeWechatWorkConf) error
}

func NewSaveConfRequest() *SaveConfRequest {
	return &SaveConfRequest{
		Session: session.NewSession(),
		Enabled: true,
		Config: wxwork.NewDefaultConfig(),
	}
}

// NewWechatWorkConfig todo
func NewWechatWorkConfig(req *SaveConfRequest) (*WechatWorkConfig, error) {
	tk := req.GetToken()
	if tk == nil {
		return nil, fmt.Errorf("token requird")
	}

	if err := req.Config.Validate(); err != nil {
		return nil, err
	}

	ins := &WechatWorkConfig{
		Domain:                tk.Domain,
		Creater:               tk.Account,
		CreateAt:              ftime.Now(),
		UpdateAt:              ftime.Now(),
		SaveConfRequest: req,
	}
	return ins, nil
}

// NewDefaultWechatWorkConfig todo
func NewDefaultWechatWorkConfig() *WechatWorkConfig {
	return &WechatWorkConfig{
		SaveConfRequest: NewSaveConfRequest(),
	}
}

type WechatWorkConfig struct {
	Domain                 string     `bson:"_id" json:"domain,omitempty"`          // 所属域ID
	Creater                string     `bson:"creater" json:"creater,omitempty"`     // 创建人
	CreateAt               ftime.Time `bson:"create_at" json:"create_at,omitempty"` // 创建时间
	UpdateAt               ftime.Time `bson:"update_at" json:"update_at,omitempty"` // 更新时间
	*SaveConfRequest `bson:",inline"`
}

// NewConfigSet 实例化
func NewConfigSet(req *request.PageRequest) *WechatWorkSet {
	return &WechatWorkSet{
		PageRequest: req,
		Items:       []*WechatWorkConfig{},
	}
}

// WechatWorkSet 列表
type WechatWorkSet struct {
	*request.PageRequest

	Total int64         `json:"total"`
	Items []*WechatWorkConfig `json:"items"`
}

// Add 添加应用
func (s *WechatWorkSet) Add(item *WechatWorkConfig) {
	s.Items = append(s.Items, item)
}

// NewDescribeConfWithDomain todo
func NewDescribeConfWithDomain(domain string) *DescribeWechatWorkConf {
	return &DescribeWechatWorkConf{
		Domain: domain,
	}
}

// DescribeWechatWorkConf 描述配置
type DescribeWechatWorkConf struct {
	Domain string
}

func (req *DescribeWechatWorkConf) Validate() error {
	if req.Domain == "" {
		return fmt.Errorf("domain required")
	}

	return nil
}

// NewQueryConfigRequest todo
func NewQueryConfigRequest(pageReq *request.PageRequest) *QueryConfigRequest {
	return &QueryConfigRequest{
		Session:     session.NewSession(),
		PageRequest: pageReq,
	}
}

// QueryConfigRequest 查询企业微信配置
type QueryConfigRequest struct {
	*session.Session
	*request.PageRequest
}


type SaveConfRequest struct {
	Enabled          bool `bson:"enabled" json:"enabled"`
	*wxwork.Config     `bson:",inline"`
	*session.Session `bson:"-" json:"-"`
}

// QueryWechatWorkConfRequest 查询企业微信配置
type QueryWechatWorkConfRequest struct {
	*session.Session
	*request.PageRequest
}

// DeleteWechatWorkConf todo
type DeleteWechatWorkConf struct {
	ID string
}
