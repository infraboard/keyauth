package token

import (
	"errors"
	fmt "fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/pb/page"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewBlockTokenRequest todo
func NewBlockTokenRequest(accessToken string, bt BlockType, reason string) *BlockTokenRequest {
	return &BlockTokenRequest{
		AccessToken: accessToken,
		BlockType:   bt,
		BlockReason: reason,
	}
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

// Validate 校验
func (m *DescribeTokenRequest) Validate() error {
	if err := validate.Struct(m); err != nil {
		return err
	}

	if m.AccessToken == "" && m.RefreshToken == "" {
		return errors.New("access_token and refresh_token required one")
	}

	return nil
}

// NewRevolkTokenRequest 撤销Token请求
func NewRevolkTokenRequest(clientID, clientSecret string) *RevolkTokenRequest {
	return &RevolkTokenRequest{
		ClientId:      clientID,
		ClientSecret:  clientSecret,
		LogoutSession: true,
	}
}

// NewQueryTokenRequest 请求实例
func NewQueryTokenRequest(page *page.PageRequest) *QueryTokenRequest {
	return &QueryTokenRequest{}
}

// Validate 校验参数
func (m *ValidateTokenRequest) Validate() error {
	if err := validate.Struct(m); err != nil {
		return err
	}

	if m.AccessToken == "" && m.RefreshToken == "" {
		return errors.New("access_token and refresh_token required one")
	}

	return nil
}

// NewValidateTokenRequest 实例化
func NewValidateTokenRequest() *ValidateTokenRequest {
	return &ValidateTokenRequest{}
}

// Validate 校验请求
func (m *IssueTokenRequest) Validate() error {
	if err := validate.Struct(m); err != nil {
		return err
	}

	switch m.GrantType {
	case GrantType_PASSWORD:
		if m.Username == "" || m.Password == "" {
			return fmt.Errorf("use %s grant type, username and password required", GrantType_PASSWORD)
		}
	case GrantType_REFRESH:
		if m.AccessToken == "" {
			return fmt.Errorf("use %s grant type, access_token required", GrantType_REFRESH)
		}
		if m.RefreshToken == "" {
			return fmt.Errorf("use %s grant type, refresh_token required", GrantType_REFRESH)
		}
	case GrantType_ACCESS:
		if m.AccessToken == "" {
			return fmt.Errorf("use %s grant type, access_token required", GrantType_ACCESS)
		}
	case GrantType_LDAP:
		if m.Username == "" || m.Password == "" {
			return fmt.Errorf("use %s grant type, username and password required", GrantType_LDAP)
		}
	case GrantType_CLIENT:
	case GrantType_AUTH_CODE:
		if m.AuthCode == "" {
			return fmt.Errorf("use %s grant type, code required", GrantType_AUTH_CODE)
		}
	default:
		return fmt.Errorf("unknown grant type %s", m.GrantType)
	}

	return nil
}

// AbnormalUserCheckKey todo
func (m *IssueTokenRequest) AbnormalUserCheckKey() string {
	return "abnormal_" + m.Username
}

// WithUserAgent todo
func (m *IssueTokenRequest) WithUserAgent(userAgent string) {
	m.UserAgent = userAgent
}

// WithRemoteIPFromHTTP todo
func (m *IssueTokenRequest) WithRemoteIPFromHTTP(r *http.Request) {
	m.RemoteIp = request.GetRemoteIP(r)
}

// WithRemoteIP todo
func (m *IssueTokenRequest) WithRemoteIP(ip string) {
	m.RemoteIp = ip
}

// GetDomainNameFromAccount todo
func (m *IssueTokenRequest) GetDomainNameFromAccount() string {
	d := strings.Split(m.Username, "@")
	if len(d) == 2 {
		return d[1]
	}

	return ""
}

// NewIssueTokenRequest 默认请求
func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{}
}

// NewIssueTokenByPassword todo
func NewIssueTokenByPassword(clientID, clientSecret, user, pass string) *IssueTokenRequest {
	return &IssueTokenRequest{
		ClientId:     clientID,
		ClientSecret: clientSecret,
		Username:     user,
		Password:     pass,
		GrantType:    GrantType_PASSWORD,
	}
}

// MakeDescribeTokenRequest todo
func (m *ValidateTokenRequest) MakeDescribeTokenRequest() *DescribeTokenRequest {
	req := NewDescribeTokenRequest()
	req.AccessToken = m.AccessToken
	req.RefreshToken = m.RefreshToken
	return req
}

// Validate todo
func (m *RevolkTokenRequest) Validate() error {
	if err := validate.Struct(m); err != nil {
		return err
	}

	return nil
}

// MakeDescribeTokenRequest todo
func (m *RevolkTokenRequest) MakeDescribeTokenRequest() *DescribeTokenRequest {
	req := NewDescribeTokenRequest()
	req.AccessToken = m.AccessToken
	req.RefreshToken = m.RefreshToken
	return req
}
