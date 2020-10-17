package token

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/pkg/user/types"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewDefaultToken todo
func NewDefaultToken() *Token {
	return &Token{}
}

// Token is user's access resource token
type Token struct {
	SessionID        string     `bson:"session_id" json:"session_id"`                           // 会话ID
	AccessToken      string     `bson:"_id" json:"access_token"`                                // 服务访问令牌
	RefreshToken     string     `bson:"refresh_token" json:"refresh_token,omitempty"`           // 用于刷新访问令牌的凭证, 刷新过后, 原先令牌将会被删除
	CreatedAt        ftime.Time `bson:"create_at" json:"create_at,omitempty"`                   // 凭证创建时间
	AccessExpiredAt  ftime.Time `bson:"access_expired_at" json:"access_expires_at,omitempty"`   // 还有多久过期
	RefreshExpiredAt ftime.Time `bson:"refresh_expired_at" json:"refresh_expired_at,omitempty"` // 刷新token过期时间

	Domain          string     `bson:"domain" json:"domain,omitempty"`                     // 用户所处域ID
	UserType        types.Type `bson:"user_type" json:"user_type,omitempty"`               // 用户类型
	Account         string     `bson:"account" json:"account,omitempty"`                   // 账户名称
	ApplicationID   string     `bson:"application_id" json:"application_id,omitempty"`     // 用户应用ID, 如果凭证是颁发给应用的, 应用在删除时需要删除所有的令牌, 应用禁用时, 该应用令牌验证会不通过
	ApplicationName string     `bson:"application_name" json:"application_name,omitempty"` // 应用名称
	ClientID        string     `bson:"client_id" json:"client_id,omitempty"`               // 客户端ID
	StartGrantType  *GrantType `bson:"start_grant_type" json:"start_grant_type,omitempty"` // 最开始授权类型
	GrantType       GrantType  `bson:"grant_type" json:"grant_type,omitempty"`             // 授权的类型
	Type            Type       `bson:"type" json:"type,omitempty"`                         // 令牌的类型 类型包含: bearer/jwt  (默认为bearer)
	Scope           string     `bson:"scope" json:"scope,omitempty"`                       // 令牌的作用范围: detail https://tools.ietf.org/html/rfc6749#section-3.3, 格式 resource-ro@k=*, resource-rw@k=*
	Description     string     `bson:"description" json:"description,omitempty"`           // 独立颁发给SDK使用时, 令牌的描述信息, 方便定位与取消
	IsBlock         bool       `bson:"is_block" json:"is_block"`                           // 是否被禁用
	BlockType       BlockType  `bson:"block_type" json:"block_type"`                       // 禁用类型
	BlockAt         ftime.Time `bson:"block_at" json:"block_at"`                           // 禁用时间
	BlockReason     string     `bson:"block_reason" json:"block_reason,omitempty"`         // 禁用原因

	remoteIP  string
	userAgent string
}

// IsRefresh todo
func (t *Token) IsRefresh() bool {
	return t.GrantType.Is(REFRESH)
}

// IsAvailable 判断一个token的可用性
func (t *Token) IsAvailable() error {
	if t.IsBlock {
		return fmt.Errorf("token is blocked")
	}

	if t.CheckAccessIsExpired() {
		return fmt.Errorf("token is expired")
	}

	return nil
}

// WithRemoteIP todo
func (t *Token) WithRemoteIP(ip string) {
	t.remoteIP = ip
}

// GetRemoteIP todo
func (t *Token) GetRemoteIP() string {
	return t.remoteIP
}

// WithUerAgent todo
func (t *Token) WithUerAgent(ua string) {
	t.userAgent = ua
}

// GetUserAgent todo
func (t *Token) GetUserAgent() string {
	return t.userAgent
}

// GetStartGrantType todo
func (t *Token) GetStartGrantType() *GrantType {
	if t.StartGrantType != nil {
		return t.StartGrantType
	}

	return &t.GrantType
}

// EndAt token结束时间
func (t *Token) EndAt() ftime.Time {
	if !t.IsBlock {
		return ftime.Time{}
	}

	if t.CheckRefreshIsExpired() {
		return t.RefreshExpiredAt
	}

	if t.CheckAccessIsExpired() {
		return t.AccessExpiredAt
	}

	return t.BlockAt
}

// CheckAccessIsExpired 检测token是否过期
func (t *Token) CheckAccessIsExpired() bool {
	if t.AccessExpiredAt.Timestamp() == 0 {
		return false
	}

	return t.AccessExpiredAt.T().Before(time.Now())
}

// CheckRefreshIsExpired 检测刷新token是否过期
func (t *Token) CheckRefreshIsExpired() bool {
	// 过期时间为0时, 标识不过期
	if t.RefreshExpiredAt.Timestamp() == 0 {
		return false
	}

	return t.RefreshExpiredAt.T().Before(time.Now())
}

// CheckTokenApplication 判断token是否属于该应用
func (t *Token) CheckTokenApplication(applicationID string) error {
	if t.ApplicationID != applicationID {
		return fmt.Errorf("the token is not issue by this application %s", applicationID)
	}

	return nil
}

// Desensitize 数据脱敏
func (t *Token) Desensitize() {
	t.RefreshToken = ""
}

// NewTokenSet 实例化
func NewTokenSet(req *request.PageRequest) *Set {
	return &Set{
		PageRequest: req,
	}
}

// Set token列表
type Set struct {
	*request.PageRequest

	Total int64    `json:"total"`
	Items []*Token `json:"items"`
}

// Add 添加
func (s *Set) Add(tk *Token) {
	s.Items = append(s.Items, tk)
}
