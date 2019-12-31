package token

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// GrantType is the type for OAuth2 param `grant_type`
type GrantType string

// oauth2 Authorization Grant: https://tools.ietf.org/html/rfc6749#section-1.3
const (
	// AUTHCODE oauth2 Authorization Code Grant
	AUTHCODE GrantType = "authorization_code"
	// IMPLICIT oauth2 Implicit Grant
	IMPLICIT GrantType = "implicit"
	// PASSWORD oauth2 Resource Owner Password Credentials Grant
	PASSWORD GrantType = "password"
	// CLIENT oauth2 Client Credentials Grant
	CLIENT GrantType = "client_credentials"
	// REFRESH oauth2 Refreshing an Access Token
	REFRESH GrantType = "refresh_token"
	// UPSCOPE is an custom grant for use unscope token acquire scope token
	UPSCOPE GrantType = "upgrade_scope"
	// WeChat is an custom grant for use unscope token acquire scope token
	WeChat GrantType = "wechat"
)

// Type token type
type Type string

// oauth2 Token Type: https://tools.ietf.org/html/rfc6749#section-7.1
const (
	// Bearer detail: https://tools.ietf.org/html/rfc6750
	Bearer Type = "bearer"
	// MAC detail: https://tools.ietf.org/html/rfc6749#ref-OAuth-HTTP-MAC
	MAC Type = "mac"
	// JWT detail:  https://tools.ietf.org/html/rfc7519
	JWT Type = "jwt"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Token is user's access resource token
type Token struct {
	AccessToken   string `bson:"access_token" json:"access_token"`             // 服务访问令牌
	RefreshToken  string `bson:"refresh_token" json:"refresh_token,omitempty"` // 用于刷新访问令牌的凭证, 刷新过后, 原先令牌将会被删除
	Name          string `bson:"name" json:"name,omitempty"`                   // 独立颁发给SDK使用时, 命名token
	Description   string `bson:"description" json:"description,omitempty"`     // 独立颁发给SDK使用时, 令牌的描述信息, 方便定位与取消
	ApplicationID string `json:"application_id,omitempty"`                     // 用户应用ID, 如果凭证是颁发给应用的, 应用在删除时需要删除所有的令牌, 应用禁用时, 该应用令牌验证会不通过
	UserID        string `json:"user_id,omitempty"`                            // 用户ID
	CreatedAt     int64  `json:"create_at,omitempty"`                          // 凭证创建时间
	ExpiresIn     int64  `json:"ttl,omitempty"`                                // 凭证过期的时间
	ExpiresAt     int64  `json:"expires_at,omitempty"`                         // 还有多久过期

	CurrentProject string `json:"current_project,omitempty"` // 当前所在项目
	DomainID       string `json:"domain_id,omitempty"`       // 用户所在的域的ID, 用户可以切换域(如果用户加入了多个域)
	ServiceID      string `json:"service_id,omitempty"`      // 服务ID, 如果凭证是颁发给内部服务使用时, 服务删除时,颁发给它的令牌需要删除, 服务禁用时, 令牌验证不通过
}

// IssueTokenRequest 颁发token请求
type IssueTokenRequest struct {
	ClientID     string    `json:"client_id,omitempty" validate:"required,lte=80"`     // 客户端ID
	ClientSecret string    `json:"client_secret,omitempty" validate:"required,lte=80"` // 客户端凭证
	Username     string    `json:"username,omitempty" validate:"lte=40"`               // 用户名
	Password     string    `json:"password,omitempty" validate:"lte=100"`              // 密码
	AccessToken  string    `json:"access_token,omitempty" validate:"lte=80"`           // 访问凭证
	RefreshToken string    `json:"refresh_token,omitempty" validate:"lte=80"`          // 刷新凭证
	AuthCode     string    `json:"code,omitempty" validate:"lte=40"`                   // https://tools.ietf.org/html/rfc6749#section-4.1.2
	State        string    `json:"state,omitempty" validate:"lte=40"`                  // https://tools.ietf.org/html/rfc6749#section-10.12
	GrantType    GrantType `json:"grant_type,omitempty" validate:"lte=20"`             // 授权的类型
	Type         Type      `json:"type,omitempty" validate:"lte=20"`                   // 令牌的类型 类型包含: bearer/jwt  (默认为bearer)
	Scope        string    `json:"scope,omitempty" validate:"lte=100"`                 // 令牌的作用范围: detail https://tools.ietf.org/html/rfc6749#section-3.3
}

// Validate 校验请求
func (req *IssueTokenRequest) Validate() error {
	if err := validate.Struct(req); err != nil {
		return err
	}

	switch req.GrantType {
	case PASSWORD:
		if req.Username == "" || req.Password == "" {
			return fmt.Errorf("use %s grant type, username and password required", PASSWORD)
		}
	case REFRESH:
		if req.RefreshToken == "" || req.AccessToken == "" {
			return fmt.Errorf("use %s grant type, access_token and refresh_token required", REFRESH)
		}
	case AUTHCODE:
		if req.AuthCode == "" {
			return fmt.Errorf("use %s grant type, code required", AUTHCODE)
		}
	default:
		return fmt.Errorf("unknown grant type %s", req.GrantType)
	}

	return nil
}
