package security

import (
	"fmt"
	"time"

	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/keyauth/pkg/domain"
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
func NewChecker(
	cache cache.Cache,
	domain domain.Service,
	maxRetry int,
	retryTTL time.Duration,
) Checker {
	return &checker{
		domain:   domain,
		cache:    cache,
		maxRetry: maxRetry,
		retryTTL: retryTTL,
		log:      zap.L().Named("Login Security"),
	}
}

type checker struct {
	domain   domain.Service
	cache    cache.Cache
	maxRetry int
	retryTTL time.Duration
	log      logger.Logger
}

func (c *checker) MaxFailedRetryCheck(req *token.IssueTokenRequest) error {
	count := 0
	err := c.cache.Get(req.AbnormalUserCheckKey(), count)
	if err != nil {
		c.log.Errorf("get key %s from cache error, %s", req.AbnormalUserCheckKey())
	}

	maxRetry, _ := c.getRetryConfig(req)
	if count > maxRetry {
		return fmt.Errorf("max retry(5)")
	}

	return nil
}

func (c *checker) UpdateFailedRetry(req *token.IssueTokenRequest) error {
	count := 0
	if c.cache.IsExist(req.AbnormalUserCheckKey()) {
		// 之前已经登陆失败过
		err := c.cache.Put(req.AbnormalUserCheckKey(), count+1)
		if err != nil {
			c.log.Errorf("set key %s to cache error, %s", req.AbnormalUserCheckKey())
		}
	} else {
		// 首次登陆失败
		_, retryTTL := c.getRetryConfig(req)
		err := c.cache.PutWithTTL(req.AbnormalUserCheckKey(), count+1, retryTTL)
		if err != nil {
			c.log.Errorf("set key %s to cache error, %s", req.AbnormalUserCheckKey())
		}
	}
	return nil
}

func (c *checker) OtherPlaceLoggedInChecK(req *token.IssueTokenRequest) error {
	return nil
}

func (c *checker) IPProtectCheck(req *token.IssueTokenRequest) error {
	return nil
}

// 如果子用户包含DomainName, 则获取Doman的设置
func (c *checker) getRetryConfig(req *token.IssueTokenRequest) (int, time.Duration) {
	domainName := req.GetDomainNameFromAccount()
	if domainName != "" {
		dm, err := c.domain.DescriptionDomain(domain.NewDescriptDomainRequestWithName(domainName))
		if err != nil {
			c.log.Errorf("get domain by name error, %s", err)
			return c.maxRetry, c.retryTTL
		}

		rc := dm.SecuritySetting.LoginSecurity.RetryLockConfig
		return rc.RetryLimiteInt(), rc.LockedMiniteDuration()
	}

	return c.maxRetry, c.retryTTL
}
