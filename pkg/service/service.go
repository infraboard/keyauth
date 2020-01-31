package service

import (
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/types/ftime"
)

const (
	// Internal 内部调用的控制面类型的服务, 提供了RPC能力,需要注册到API 网关对内提供服务
	Internal Type = "internal"
	// Public 需要对外发布的控制面类型的服务, 提供了RPC能力, 需要注册到API 网关对外提供服务
	Public = "public"
)

// Type 服务类型
type Type string

// MicroService is service provider
type MicroService struct {
	*CreateServiceRequest
	ID           string     `json:"id"`                      // 唯一ID
	CreateAt     int64      `json:"create_at,omitempty"`     // 创建的时间
	UpdateAt     int64      `json:"update_at,omitempty"`     // 更新时间
	ClientID     string     `json:"client_id,omitempty"`     // 客户端ID
	ClientSecret string     `json:"client_secret,omitempty"` // 客户端秘钥
	Features     []*Feature `json:"features,omitempty"`      // 服务功能列表
}

// CreateServiceRequest 服务创建请求
type CreateServiceRequest struct {
	Name            string `json:"name" validate:"required,lte=80"` // 名称
	Type            Type   `json:"type" validate:"required,lte=80"` // 服务类型
	Description     string `json:"description,omitempty"`           // 描述信息
	Enabled         bool   `json:"enabled"`                         // 是否启用该服务
	TokenExpireTime int64  `json:"token_expire_time,omitempty"`     // 客户端凭证申请的token的过期时间
}

// Feature Service's features
type Feature struct {
	CreateAt  ftime.Time `json:"create_at,omitempty"`  // 创建时间
	UpdateAt  ftime.Time `json:"update_at,omitempty"`  // 更新时间
	ServiceID string     `json:"service_id,omitempty"` // 该功能属于那个服务
	Version   string     `json:"version,omitempty"`    // 服务那个版本的功能
	router.Entry
}
