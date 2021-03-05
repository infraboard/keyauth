package pkg

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/router"
	httpb "github.com/infraboard/mcube/pb/http"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/keyauth/version"
)

// GetInternalAdminTokenCtx 内部调用时的模拟token
func GetInternalAdminTokenCtx(account string) context.Context {
	return session.WithTokenContext(context.Background(), &token.Token{
		Account:  account,
		Domain:   domain.AdminDomainName,
		UserType: types.UserType_INTERNAL,
	})
}

// NewInternalAuther 内部使用的auther
func NewInternalAuther() router.Auther {
	return &httpAuther{}
}

// internal todo
type httpAuther struct{}

func (i *httpAuther) Auth(r *http.Request, entry httpb.Entry) (
	authInfo interface{}, err error) {
	var tk *token.Token
	if entry.AuthEnable {

	}

	if entry.PermissionEnable && tk != nil {
		// 如果是超级管理员不做权限校验, 直接放行
		if tk.UserType.IsIn(types.UserType_SUPPER) {
			return tk, nil
		}

		// 其他比如服务类型, 主账号类型, 子账号类型
		// 如果开启权限认证都需要检查
		if Permission == nil {
			return nil, fmt.Errorf("permission service not load")
		}

		req := permission.NewCheckPermissionrequest()
		req.EndpointId = i.endpointHashID(entry)
		ctx := session.WithTokenContext(context.Background(), tk)
		_, err = Permission.CheckPermission(ctx, req)
		if err != nil {
			return nil, exception.NewPermissionDeny("no permission")
		}
	}

	return tk, nil
}

func (i *httpAuther) endpointHashID(entry httpb.Entry) string {
	return endpoint.GenHashID(version.ServiceName, entry.GrpcPath)
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
