package security

import (
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
	MaxFailedRetryCheck(*token.IssueTokenRequest) error
	UpdateFailedRetry(*token.IssueTokenRequest) error
}

// ExceptionLockChecKer 异地登录限制
type ExceptionLockChecKer interface {
	OtherPlaceLoggedInChecK(*token.Token) error
	NotLoginDaysChecK(*token.Token) error
}

// IPProtectChecker todo
type IPProtectChecker interface {
	IPProtectCheck(*token.IssueTokenRequest) error
}
