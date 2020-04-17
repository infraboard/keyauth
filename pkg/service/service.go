package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/types/ftime"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Type 服务类型
type Type string

// MicroService is service provider
type MicroService struct {
	*CreateServiceRequest `bson:",inline"`
	CreateAt              ftime.Time `bson:"create_at" json:"create_at,omitempty"`     // 创建的时间
	UpdateAt              ftime.Time `bson:"update_at" json:"update_at,omitempty"`     // 更新时间
	ServiceID             string     `bson:"service_id" json:"service_id,omitempty"`   // 客户端ID
	ServiceKey            string     `bson:"service_key" json:"service_key,omitempty"` // 客户端秘钥
	Features              []*Feature `bson:"features" json:"features,omitempty"`       // 服务功能列表
}

// CheckKey 校验服务key是否正确
func (ms *MicroService) CheckKey(key string) bool {
	return ms.ServiceKey == key
}

// New 创建服务
func New(req *CreateServiceRequest) (*MicroService, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := &MicroService{
		CreateServiceRequest: req,
		CreateAt:             ftime.Now(),
		UpdateAt:             ftime.Now(),
		ServiceID:            token.MakeBearer(24),
		ServiceKey:           token.MakeBearer(36),
	}

	return ins, nil
}

// NewCreateServiceRequest todo
func NewCreateServiceRequest() *CreateServiceRequest {
	return &CreateServiceRequest{
		Enabled:         true,
		TokenExpireTime: 3600,
	}
}

// CreateServiceRequest 服务创建请求
type CreateServiceRequest struct {
	Name            string   `bson:"_id" json:"name" validate:"required,lte=80"`           // 名称
	Label           []string `bson:"label" json:"label" validate:"lte=80"`                 // 服务标签
	Description     string   `bson:"description" json:"description,omitempty"`             // 描述信息
	Enabled         bool     `bson:"enabled" json:"enabled"`                               // 是否启用该服务
	TokenExpireTime int64    `bson:"token_expire_time" json:"token_expire_time,omitempty"` // 凭证申请的token的过期时间
	UserID          string   `bson:"user_id" json:"-"`                                     // 创建者ID
}

// Validate 校验请求是否合法
func (req *CreateServiceRequest) Validate() error {
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

// NewMicroServiceSet 实例化
func NewMicroServiceSet(req *request.PageRequest) *MicroServiceSet {
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
