package token

import "fmt"

// oauth2 Authorization Grant: https://tools.ietf.org/html/rfc6749#section-1.3
const (
	UNKNOWN GrantType = "unknwon"
	// AUTHCODE oauth2 Authorization Code Grant
	AUTHCODE = "authorization_code"
	// IMPLICIT oauth2 Implicit Grant
	IMPLICIT = "implicit"
	// PASSWORD oauth2 Resource Owner Password Credentials Grant
	PASSWORD = "password"
	// CLIENT oauth2 Client Credentials Grant
	CLIENT = "client_credentials"
	// REFRESH oauth2 Refreshing an Access Token
	REFRESH = "refresh_token"
	// ACCESS is an custom grant for use use generate personal private token
	ACCESS = "access_token"
	// LDAP 通过ldap认证
	LDAP = "ldap"
)

// ParseGrantTypeFromString todo
func ParseGrantTypeFromString(str string) (GrantType, error) {
	switch str {
	case "authorization_code":
		return AUTHCODE, nil
	case "implicit":
		return IMPLICIT, nil
	case "password":
		return PASSWORD, nil
	case "client_credentials":
		return CLIENT, nil
	case "refresh_token":
		return REFRESH, nil
	case "access_token":
		return ACCESS, nil
	case "ldap":
		return LDAP, nil
	default:
		return UNKNOWN, fmt.Errorf("unknown Grant type: %s", str)
	}
}

// GrantType is the type for OAuth2 param ` grant_type`
type GrantType string

// Is 判断类型
func (t GrantType) Is(tps ...GrantType) bool {
	for i := range tps {
		if tps[i] == t {
			return true
		}
	}
	return false
}

// oauth2 Token Type: https://tools.ietf.org/html/rfc6749#section-7.1
const (
	// Bearer detail: https://tools.ietf.org/html/rfc6750
	Bearer Type = "bearer"
	// MAC detail: https://tools.ietf.org/html/rfc6749#ref-OAuth-HTTP-MAC
	MAC = "mac"
	// JWT detail:  https://tools.ietf.org/html/rfc7519
	JWT = "jwt"
)

// Type token type
type Type string

const (
	// OtherClientLoggedIn 启动端登录
	OtherClientLoggedIn BlockType = "other_client_logged_in"
	// SessionTerminated 会话中断
	SessionTerminated = "session_terminated"
	// OtherPlaceLoggedIn 异地登录保护
	OtherPlaceLoggedIn = "other_place_logged_in"
	// OtherIPLoggedIn 登录IP保护
	OtherIPLoggedIn = "other_ip_logged_in"
)

// BlockType 禁用类型
type BlockType string
