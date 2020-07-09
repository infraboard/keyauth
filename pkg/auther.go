package pkg

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg/token"
)

// NewInternalAuther 内部使用的auther
func NewInternalAuther() router.Auther {
	return &internal{}
}

// internal todo
type internal struct {
}

func (i *internal) Auth(r *http.Request, entry router.Entry) (
	authInfo interface{}, err error) {
	if entry.AuthEnable {
		req := token.NewValidateTokenRequest()
		// 获取需要校验的access token(用户的身份凭证)
		accessToken := r.Header.Get("x-oauth-token")
		if accessToken == "" {
			return nil, exception.NewUnauthorized("x-oauth-token header required")
		}
		req.AccessToken = accessToken

		tk, err := Token.ValidateToken(req)
		if err != nil {
			return nil, err
		}

		return tk, nil
	}

	return nil, nil
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	// Case insensitive prefix match. See Issue 22736.
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

// GetTokenFromContext 从上下文中获取Token
func GetTokenFromContext(r *http.Request) (*token.Token, error) {
	tk, ok := context.GetContext(r).AuthInfo.(*token.Token)
	if !ok {
		return nil, exception.NewInternalServerError("authInfo is not token pointer")
	}

	return tk, nil
}
