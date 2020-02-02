package service

import (
	"github.com/go-playground/validator/v10"
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

// MicroService is service provider
type MicroService struct {
	*CreateServiceRequest
	ID           string     `json:"id"`                      // 唯一ID
	CreateAt     ftime.Time `json:"create_at,omitempty"`     // 创建的时间
	UpdateAt     ftime.Time `json:"update_at,omitempty"`     // 更新时间
	ClientID     string     `json:"client_id,omitempty"`     // 客户端ID
	ClientSecret string     `json:"client_secret,omitempty"` // 客户端秘钥
	Features     []*Feature `json:"features,omitempty"`      // 服务功能列表
}

// New 创建服务
func New(req *CreateServiceRequest) (*MicroService, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := &MicroService{
		CreateServiceRequest: req,
		ID:                   xid.New().String(),
		CreateAt:             ftime.Now(),
		UpdateAt:             ftime.Now(),
	}

	return ins, nil
}

// CreateServiceRequest 服务创建请求
type CreateServiceRequest struct {
	Name            string   `json:"name" validate:"required,lte=80"` // 名称
	Label           []string `json:"label" validate:"lte=80"`         // 服务标签
	Description     string   `json:"description,omitempty"`           // 描述信息
	Enabled         bool     `json:"enabled"`                         // 是否启用该服务
	TokenExpireTime int64    `json:"token_expire_time,omitempty"`     // 凭证申请的token的过期时间
}

// Validate 校验请求是否合法
func (req *CreateServiceRequest) Validate() error {
	return validate.Struct(req)
}

// Feature Service's features
type Feature struct {
	CreateAt  ftime.Time `json:"create_at,omitempty"`  // 创建时间
	UpdateAt  ftime.Time `json:"update_at,omitempty"`  // 更新时间
	ServiceID string     `json:"service_id,omitempty"` // 该功能属于那个服务
	Version   string     `json:"version,omitempty"`    // 服务那个版本的功能
	router.Entry
}

// NewApplicationSet 实例化
func NewApplicationSet(req *request.PageRequest) *MicroServiceSet {
	return &MicroServiceSet{

		PageRequest: req,
	}
}

// MicroServiceSet 列表
type MicroServiceSet struct {
	*request.PageRequest

	Total int64           `json:"total"`
	Items []*MicroService `json:"items"`
}

// Add 添加
func (s *MicroServiceSet) Add(e *MicroService) {
	s.Items = append(s.Items, e)
}
