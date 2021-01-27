package security

import (
	"context"

	"github.com/infraboard/keyauth/pkg/token"
)

// Checker 安全检测
type Checker interface {
	MaxTryChecker
	ExceptionLockChecKer
	IPProtectChecker
}

// MaxTryChecker todo 失败重试限制
type MaxTryChecker interface {
	MaxFailedRetryCheck(context.Context, *token.IssueTokenRequest) error
	UpdateFailedRetry(context.Context, *token.IssueTokenRequest) error
}

// ExceptionLockChecKer 异地登录限制
type ExceptionLockChecKer interface {
	OtherPlaceLoggedInChecK(context.Context, *token.Token) error
	NotLoginDaysChecK(context.Context, *token.Token) error
}

// IPProtectChecker todo
type IPProtectChecker interface {
	IPProtectCheck(context.Context, *token.IssueTokenRequest) error
}
