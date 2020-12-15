package token

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/infraboard/mcube/http/request"
)

var (
	// DefaultForwareHeaderKey 协商forward ip 的 hander key名称
	defaultScanForwareHeaderKey = []string{"X-Forwarded-For", "X-Real-IP"}
)

// Service token管理服务
type Service interface {
	IssueToken(*IssueTokenRequest) (*Token, error)
	ValidateToken(*ValidateTokenRequest) (*Token, error)
	DescribeToken(*DescribeTokenRequest) (*Token, error)
	RevolkToken(*RevolkTokenRequest) error
	QueryToken(*QueryTokenRequest) (*Set, error)
	BlockToken(*BlockTokenRequest) (*Token, error)
}

// NewIssueTokenRequest 默认请求
func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{}
}

// NewIssueTokenByPassword todo
func NewIssueTokenByPassword(clientID, clientSecret, user, pass string) *IssueTokenRequest {
	return &IssueTokenRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Username:     user,
		Password:     pass,
		GrantType:    PASSWORD,
	}
}

// IssueTokenRequest 颁发token请求
type IssueTokenRequest struct {
	ClientID     string    `json:"client_id,omitempty" validate:"required,lte=80"`     // 客户端ID
	ClientSecret string    `json:"client_secret,omitempty" validate:"required,lte=80"` // 客户端凭证
	Username     string    `json:"username,omitempty" validate:"lte=40"`               // 用户名
	Password     string    `json:"password,omitempty" validate:"lte=100"`              // 密码
	RefreshToken string    `json:"refresh_token,omitempty" validate:"lte=80"`          // 刷新凭证
	AccessToken  string    `json:"access_token,omitempty" validate:"lte=80"`           // 访问凭证
	AuthCode     string    `json:"code,omitempty" validate:"lte=40"`                   // https://tools.ietf.org/html/rfc6749#section-4.1.2
	State        string    `json:"state,omitempty" validate:"lte=40"`                  // https://tools.ietf.org/html/rfc6749#section-10.12
	GrantType    GrantType `json:"grant_type,omitempty" validate:"lte=20"`             // 授权的类型
	Type         Type      `json:"type,omitempty" validate:"lte=20"`                   // 令牌的类型 类型包含: bearer/jwt  (默认为bearer)
	Scope        string    `json:"scope,omitempty" validate:"lte=100"`                 // 令牌的作用范围: detail https://tools.ietf.org/html/rfc6749#section-3.3

	ua string
	ip string
}

// AbnormalUserCheckKey todo
func (req *IssueTokenRequest) AbnormalUserCheckKey() string {
	return "abnormal_" + req.Username
}

// WithUserAgent todo
func (req *IssueTokenRequest) WithUserAgent(userAgent string) {
	req.ua = userAgent
}

// GetUserAgent todo
func (req *IssueTokenRequest) GetUserAgent() string {
	return req.ua
}

// WithRemoteIPFromHTTP todo
func (req *IssueTokenRequest) WithRemoteIPFromHTTP(r *http.Request) {
	req.ip = request.GetRemoteIP(r)
}

// WithRemoteIP todo
func (req *IssueTokenRequest) WithRemoteIP(ip string) {
	req.ip = ip
}

// GetRemoteIP todo
func (req *IssueTokenRequest) GetRemoteIP() string {
	return req.ip
}

// GetDomainNameFromAccount todo
func (req *IssueTokenRequest) GetDomainNameFromAccount() string {
	d := strings.Split(req.Username, "@")
	if len(d) == 2 {
		return d[1]
	}

	return ""
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
		if req.AccessToken == "" {
			return fmt.Errorf("use %s grant type, access_token required", REFRESH)
		}
		if req.RefreshToken == "" {
			return fmt.Errorf("use %s grant type, refresh_token required", REFRESH)
		}
	case ACCESS:
		if req.AccessToken == "" {
			return fmt.Errorf("use %s grant type, access_token required", ACCESS)
		}
	case LDAP:
		if req.Username == "" || req.Password == "" {
			return fmt.Errorf("use %s grant type, username and password required", LDAP)
		}
	case CLIENT:
	case AUTHCODE:
		if req.AuthCode == "" {
			return fmt.Errorf("use %s grant type, code required", AUTHCODE)
		}
	default:
		return fmt.Errorf("unknown grant type %s", req.GrantType)
	}

	return nil
}

// NewValidateTokenRequest 实例化
func NewValidateTokenRequest() *ValidateTokenRequest {
	return &ValidateTokenRequest{
		DescribeTokenRequest: NewDescribeTokenRequest(),
	}
}

// ValidateTokenRequest 校验token
type ValidateTokenRequest struct {
	NamesapceID string `json:"namespace_id,omitempty" validate:"lte=100"` // Namespace ID
	EndpointID  string `json:"endpoint_id,omitempty" validate:"lte=400"`  // Endpoint ID(hash ID)
	*DescribeTokenRequest
}

// Validate 校验参数
func (req *ValidateTokenRequest) Validate() error {
	if req.DescribeTokenRequest == nil {
		return errors.New("DescribeTokenRequest required")
	}
	if err := req.DescribeTokenRequest.Validate(); err != nil {
		return err
	}

	return nil
}

// NewQueryTokenRequest 请求实例
func NewQueryTokenRequest(page *request.PageRequest) *QueryTokenRequest {
	return &QueryTokenRequest{
		PageRequest: page,
	}
}

// QueryTokenRequest 查询Token列表
type QueryTokenRequest struct {
	*request.PageRequest
	ApplicationID string    `json:"application_id,omitempty"`
	GrantType     GrantType `json:"grant_type,omitempty"`
}

// NewRevolkTokenRequest 撤销Token请求
func NewRevolkTokenRequest(clientID, clientSecret string) *RevolkTokenRequest {
	return &RevolkTokenRequest{
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		LogoutSession:        true,
		DescribeTokenRequest: NewDescribeTokenRequest(),
	}
}

// RevolkTokenRequest 撤销Token的请求
type RevolkTokenRequest struct {
	ClientSecret  string `json:"client_secret,omitempty" validate:"required,lte=80"` // 客户端凭证
	ClientID      string `json:"client_id,omitempty" validate:"required,lte=80"`     // 客户端ID
	LogoutSession bool   `json:"logout_session"`                                     // 是否退出会话, 当刷新token时 不需要退出会话
	*DescribeTokenRequest
}

// NewDescribeTokenRequest 实例化
func NewDescribeTokenRequest() *DescribeTokenRequest {
	return &DescribeTokenRequest{}
}

// NewDescribeTokenRequestWithAccessToken 实例化
func NewDescribeTokenRequestWithAccessToken(at string) *DescribeTokenRequest {
	req := NewDescribeTokenRequest()
	req.AccessToken = at
	return req
}

// DescribeTokenRequest 撤销请求
type DescribeTokenRequest struct {
	AccessToken  string `json:"access_token,omitempty" validate:"lte=80"`  // 访问凭证
	RefreshToken string `json:"refresh_token,omitempty" validate:"lte=80"` // 访问凭证
}

// Validate 校验
func (req *DescribeTokenRequest) Validate() error {
	if err := validate.Struct(req); err != nil {
		return err
	}

	if req.AccessToken == "" && req.RefreshToken == "" {
		return errors.New("access_token and refresh_token required one")
	}

	return nil
}

// NewBlockTokenRequest todo
func NewBlockTokenRequest(accessToken string, bt BlockType, reason string) *BlockTokenRequest {
	return &BlockTokenRequest{
		AccessToken: accessToken,
		BlcokType:   bt,
		BlockReson:  reason,
	}
}

// BlockTokenRequest 禁用请求
type BlockTokenRequest struct {
	AccessToken string
	BlockReson  string
	BlcokType   BlockType
}
