package pkg

import (
	"context"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
)

// GetInternalAdminTokenCtx 内部调用时的模拟token
func GetInternalAdminTokenCtx(account string) context.Context {
	return session.WithTokenContext(context.Background(), &token.Token{
		Account:  account,
		Domain:   domain.AdminDomainName,
		UserType: types.UserType_INTERNAL,
	})
}
