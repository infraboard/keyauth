package security

import (
	"fmt"
	"time"

	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/logger"

	"github.com/infraboard/keyauth/pkg/token"
)

// Checker 安全检测
type Checker interface {
	MaxTryChecker
	OtherPlaceLoggedInChecKer
	IPProtectChecker
}

// MaxTryChecker todo 失败重试限制
type MaxTryChecker interface {
	MaxFailedRetryCheck(*token.IssueTokenRequest) error
	UpdateFailedRetry(*token.IssueTokenRequest) error
}

// OtherPlaceLoggedInChecKer 异地登录限制
type OtherPlaceLoggedInChecKer interface {
	OtherPlaceLoggedInChecK(*token.IssueTokenRequest) error
}

// IPProtectChecker todo
type IPProtectChecker interface {
	IPProtectCheck(*token.IssueTokenRequest) error
}

// NewChecker todo
func NewChecker() Checker {
	return &checker{}
}

// FailedLogin 记录
type checker struct {
	cache    cache.Cache
	maxRetry int
	retryTTL time.Duration
	log      logger.Logger
}

// CheckBlook 判断是否被阻断
func (c *checker) MaxFailedRetryCheck(req *token.IssueTokenRequest) error {
	count := 0
	err := c.cache.Get(req.AbnormalUserCheckKey(), count)
	if err != nil {
		c.log.Errorf("get key %s from cache error, %s", req.AbnormalUserCheckKey())
	}
	if count > c.maxRetry {
		return fmt.Errorf("max retry(5)")
	}

	return nil
}

func (c *checker) UpdateFailedRetry(req *token.IssueTokenRequest) error {
	count := 0
	err := c.cache.Get(req.AbnormalUserCheckKey(), count)
	if err != nil {
		c.log.Errorf("get key %s from cache error, %s", req.AbnormalUserCheckKey())
	}

	err = c.cache.PutWithTTL(req.AbnormalUserCheckKey(), count+1, c.retryTTL)
	if err != nil {
		c.log.Errorf("set key %s to cache error, %s", req.AbnormalUserCheckKey())
	}
	return nil
}

func (c *checker) OtherPlaceLoggedInChecK(req *token.IssueTokenRequest) error {
	return nil
}

func (c *checker) IPProtectCheck(req *token.IssueTokenRequest) error {
	return nil
}
