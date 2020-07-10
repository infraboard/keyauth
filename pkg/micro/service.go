package micro

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Type 服务类型
type Type string

// Micro is service provider
type Micro struct {
	ID                  string `bson:"_id" json:"id"` // 微服务ID
	*CreateMicroRequest `bson:",inline"`
	CreaterID           string     `bson:"creater_id" json:"creater_id"`
	CreateAt            ftime.Time `bson:"create_at" json:"create_at,omitempty"` // 创建的时间
	UpdateAt            ftime.Time `bson:"update_at" json:"update_at,omitempty"` // 更新时间
	AccountID           string     `bson:"account_id" json:"account_id"`         // 服务账号
	AccessToken         string     `bson:"access_token" json:"access_token"`     // 服务访问凭证
	RefreshToken        string     `bson:"refresh_token" json:"-"`               // 服务刷新凭证
	Features            []*Feature `bson:"features" json:"features,omitempty"`   // 服务功能列表
}

// New 创建服务
func New(req *CreateMicroRequest) (*Micro, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := &Micro{
		ID:                 xid.New().String(),
		CreateAt:           ftime.Now(),
		UpdateAt:           ftime.Now(),
		CreateMicroRequest: req,
		Features:           []*Feature{},
	}

	return ins, nil
}

// NewCreateMicroRequest todo
func NewCreateMicroRequest() *CreateMicroRequest {
	return &CreateMicroRequest{
		Session: token.NewSession(),
		Enabled: true,
		Label:   map[string]string{},
	}
}

// CreateMicroRequest 服务创建请求
type CreateMicroRequest struct {
	*token.Session  `bson:"-" json:"-"`
	Name            string            `bson:"name" json:"name" validate:"required,lte=200"`         // 名称
	Label           map[string]string `bson:"label" json:"label" validate:"lte=80"`                 // 服务标签
	Description     string            `bson:"description" json:"description,omitempty"`             // 描述信息
	Enabled         bool              `bson:"enabled" json:"enabled"`                               // 是否启用该服务
	TokenExpireTime int64             `bson:"token_expire_time" json:"token_expire_time,omitempty"` // 凭证申请的token的过期时间                               // 创建者ID
}

// Validate 校验请求是否合法
func (req *CreateMicroRequest) Validate() error {
	if req.GetToken() == nil {
		return fmt.Errorf("session context required")
	}

	return validate.Struct(req)
}

// Feature Service's features
type Feature struct {
	CreateAt     ftime.Time `bson:"create_at" json:"create_at,omitempty"`     // 创建时间
	UpdateAt     ftime.Time `bson:"update_at" json:"update_at,omitempty"`     // 更新时间
	ServiceName  string     `bson:"service_name" json:"service_id,omitempty"` // 该功能属于那个服务
	Version      string     `bson:"version" json:"version,omitempty"`         // 服务那个版本的功能
	router.Entry `bson:",inline"`
}

// NewMicroSet 实例化
func NewMicroSet(req *request.PageRequest) *Set {
	return &Set{
		PageRequest: req,
	}
}

// Set 列表
type Set struct {
	*request.PageRequest

	Total int64    `json:"total"`
	Items []*Micro `json:"items"`
}

// Add 添加
func (s *Set) Add(e *Micro) {
	s.Items = append(s.Items, e)
}
