package application

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Service 用户服务
type Service interface {
	CreateUserApplication(userID string, req *CreateApplicatonRequest) (*Application, error)
	DeleteApplication(id string) error
	DescriptionApplication(req *DescriptApplicationRequest) (*Application, error)
	QueryApplication(req *QueryApplicationRequest) ([]*Application, int64, error)
}

// NewDescriptApplicationRequest new实例
func NewDescriptApplicationRequest() *DescriptApplicationRequest {
	return &DescriptApplicationRequest{}
}

// DescriptApplicationRequest 查询应用详情
type DescriptApplicationRequest struct {
	ID       string
	ClientID string
}

// Validate 校验详情查询请求
func (req *DescriptApplicationRequest) Validate() error {
	if req.ID == "" && req.ClientID == "" {
		return errors.New("id or client_id is required")
	}

	return nil
}

// NewQueryApplicationRequest 列表查询请求
func NewQueryApplicationRequest(pageReq *request.PageRequest) *QueryApplicationRequest {
	return &QueryApplicationRequest{
		PageRequest: pageReq,
	}
}

// QueryApplicationRequest 查询应用列表
type QueryApplicationRequest struct {
	*request.PageRequest
	UserID string
}

// NewCreateApplicatonRequest 请求
func NewCreateApplicatonRequest() *CreateApplicatonRequest {
	return &CreateApplicatonRequest{}
}

// CreateApplicatonRequest 创建应用请求
type CreateApplicatonRequest struct {
	Name            string `bson:"name" json:"name,omitempty" validate:"required,lte=30"`         // 应用名称
	Website         string `bson:"website" json:"website,omitempty" validate:"lte=200"`           // 应用的网站地址
	LogoImage       string `bson:"logo_image" json:"logo_image,omitempty" validate:"lte=200"`     // 应用的LOGO
	Description     string `bson:"description" json:"description,omitempty" validate:"lte=1000"`  // 应用简单的描述
	RedirectURI     string `bson:"redirect_uri" json:"redirect_uri,omitempty" validate:"lte=200"` // 应用重定向URI, Oauht2时需要该参数
	TokenExpireTime int64  `bson:"token_expire_time" json:"token_expire_time"`                    // 应用申请的token的过期时间
}

// Validate 请求校验
func (req *CreateApplicatonRequest) Validate() error {
	return validate.Struct(req)
}
