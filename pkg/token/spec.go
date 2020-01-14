package token

import (
	"errors"
	"fmt"
)

// Service token管理服务
type Service interface {
	IssueToken(req *IssueTokenRequest) (*Token, error)
	ValidateToken(req *ValidateTokenRequest) (*Token, error)
	RevolkToken(req *DescribeTokenRequest) error
}

// NewIssueTokenRequest 默认请求
func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{}
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
		if req.RefreshToken == "" {
			return fmt.Errorf("use %s grant type, refresh_token required", REFRESH)
		}
	case Access:
		if req.AccessToken == "" {
			return fmt.Errorf("use %s grant type, access_token required", Access)
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
	*DescribeTokenRequest
	Endpoint string `json:"endpoint,omitempty" validate:"lte=400"` // 判断
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

// NewDescribeTokenRequest 实例化
func NewDescribeTokenRequest() *DescribeTokenRequest {
	return &DescribeTokenRequest{}
}

// DescribeTokenRequest 撤销请求
type DescribeTokenRequest struct {
	ClientID     string `json:"client_id,omitempty" validate:"required,lte=80"`     // 客户端ID
	ClientSecret string `json:"client_secret,omitempty" validate:"required,lte=80"` // 客户端凭证
	AccessToken  string `json:"access_token,omitempty" validate:"lte=80"`           // 访问凭证
	RefreshToken string `json:"refresh_token,omitempty" validate:"lte=80"`          // 访问凭证
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
