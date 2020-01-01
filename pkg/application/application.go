package application

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/pkg/token"
)

// ClientType 客户端类型
type ClientType string

const (
	// Confidential (server-based) https://tools.ietf.org/html/rfc6749#section-2.1
	Confidential ClientType = "confidential"
	// Public （client-based)
	Public ClientType = "public"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

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

// NewUserApplicartion 新建实例
func NewUserApplicartion(userID string, t ClientType, req *CreateApplicatonRequest) (*Application, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	app := newDeafultApplication(req)
	app.UserID = userID

	return app, nil
}

// NewServiceApplicartion 新建实例
func NewServiceApplicartion(serviceID string, t ClientType, req *CreateApplicatonRequest) (*Application, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	app := newDeafultApplication(req)
	app.ServiceID = serviceID

	return app, nil
}

func newDeafultApplication(req *CreateApplicatonRequest) *Application {
	return &Application{
		ID:                      xid.New().String(),
		CreateAt:                ftime.Now(),
		UpdateAt:                ftime.Now(),
		ClientType:              Public,
		ClientID:                token.MakeBearer(16),
		ClientSecret:            token.MakeBearer(32),
		CreateApplicatonRequest: req,
	}
}

// Application is oauth2's client: https://tools.ietf.org/html/rfc6749#section-2
type Application struct {
	ID                       string     `bson:"_id" json:"id,omitempty"`                      // 唯一ID
	UserID                   string     `bson:"user_id" json:"user_id,omitempty"`             // 应用属于那个用户
	ServiceID                string     `bson:"service_id" json:"service_id,omitempty"`       // 应用属于一个service
	CreateAt                 ftime.Time `bson:"create_at" json:"create_at,omitempty"`         // 应用创建的时间
	UpdateAt                 ftime.Time `bson:"update_at" json:"update_at,omitempty"`         // 应用更新的时间
	ClientType               ClientType `bson:"client_type" json:"client_type,omitempty"`     // 客户端类型
	ClientID                 string     `bson:"client_id" json:"client_id,omitempty"`         // 应用客户端ID
	ClientSecret             string     `bson:"client_secret" json:"client_secret,omitempty"` // 应用客户端秘钥
	Locked                   bool       `bson:"locked" json:"locked,omitempty"`               // 是否冻结应用, 冻结应用后, 该应用无法通过凭证获取访问凭证(token)
	*CreateApplicatonRequest `bson:",inline"`
}

// CheckClientSecret 判断凭证是否合法
func (a *Application) CheckClientSecret(secret string) error {
	if a.ClientSecret != secret {
		return errors.New("client_secret is not correct")
	}

	return nil
}
