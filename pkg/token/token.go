package token

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/types/ftime"
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
	AccessToken      string     `bson:"access_token" json:"access_token"`                       // 服务访问令牌
	RefreshToken     string     `bson:"refresh_token" json:"refresh_token,omitempty"`           // 用于刷新访问令牌的凭证, 刷新过后, 原先令牌将会被删除
	Name             string     `bson:"name" json:"name,omitempty"`                             // 独立颁发给SDK使用时, 命名token
	Description      string     `bson:"description" json:"description,omitempty"`               // 独立颁发给SDK使用时, 令牌的描述信息, 方便定位与取消
	UserID           string     `bson:"user_id" json:"user_id,omitempty"`                       // 用户ID
	ApplicationID    string     `bson:"application_id" json:"application_id,omitempty"`         // 用户应用ID, 如果凭证是颁发给应用的, 应用在删除时需要删除所有的令牌, 应用禁用时, 该应用令牌验证会不通过
	CreatedAt        ftime.Time `bson:"create_at" json:"create_at,omitempty"`                   // 凭证创建时间
	AccessExpiredAt  ftime.Time `bson:"access_expired_at" json:"access_expires_at,omitempty"`   // 还有多久过期
	RefreshExpiredAt ftime.Time `bson:"refresh_expired_at" json:"refresh_expired_at,omitempty"` // 刷新token过期时间

	ClientID  string    `bson:"client_id" json:"client_id,omitempty"`   // 客户端ID
	GrantType GrantType `bson:"grant_type" json:"grant_type,omitempty"` // 授权的类型
	Type      Type      `bson:"type" json:"type,omitempty"`             // 令牌的类型 类型包含: bearer/jwt  (默认为bearer)
	Scope     string    `bson:"scope" json:"scope,omitempty"`           // 令牌的作用范围: detail https://tools.ietf.org/html/rfc6749#section-3.3, 格式 resource-ro@k=*, resource-rw@k=*
}

// CheckAccessIsExpired 检测token是否过期
func (t *Token) CheckAccessIsExpired() bool {
	return t.AccessExpiredAt.T().Before(time.Now())
}

// CheckRefreshIsExpired 检测刷新token是否过期
func (t *Token) CheckRefreshIsExpired() bool {
	return t.RefreshExpiredAt.T().Before(time.Now())
}
