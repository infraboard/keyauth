package application

import "github.com/go-playground/validator/v10"

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Service 用户服务
type Service interface {
	CreateUserApplication(userID string, req *CreateApplicatonRequest) (*Application, error)
	DeleteApplication(id string) error
	DescriptionApplication(req *DescriptApplicationRequest) (*Application, error)
}

// DescriptApplicationRequest 查询应用详情
type DescriptApplicationRequest struct {
	ID string `json:"id,omitempty"`
}

// NewCreateApplicatonRequest 请求
func NewCreateApplicatonRequest() *CreateApplicatonRequest {
	return &CreateApplicatonRequest{}
}

// CreateApplicatonRequest 创建应用请求
type CreateApplicatonRequest struct {
	Name            string `bson:"name" json:"name" validate:"required,lte=30"`               // 应用名称
	Website         string `bson:"website" json:"website,omitempty" validate:"lte=200"`       // 应用的网站地址
	LogoImage       string `bson:"logo_image" json:"logo_image,omitempty" validate:"lte=200"` // 应用的LOGO
	Description     string `bson:"description" json:"description" validate:"lte=1000"`        // 应用简单的描述
	RedirectURI     string `bson:"redirect_uri" json:"redirect_uri" validate:"lte=200"`       // 应用重定向URI, Oauht2时需要该参数
	TokenExpireTime int64  `bson:"token_expire_time" json:"token_expire_time"`                // 应用申请的token的过期时间
}

// Validate 请求校验
func (req *CreateApplicatonRequest) Validate() error {
	return validate.Struct(req)
}
