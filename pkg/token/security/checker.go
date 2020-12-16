package security

import (
	"fmt"

	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/ip2region"
	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
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
	OtherPlaceLoggedInChecK(*token.IssueTokenRequest) error
	NotLoginDaysChecK(*token.IssueTokenRequest) error
}

// IPProtectChecker todo
type IPProtectChecker interface {
	IPProtectCheck(*token.IssueTokenRequest) error
}

// NewChecker todo
func NewChecker() (Checker, error) {
	if pkg.Domain == nil {
		return nil, fmt.Errorf("denpence domain service required")
	}
	if pkg.User == nil {
		return nil, fmt.Errorf("denpence user service required")
	}
	if pkg.Session == nil {
		return nil, fmt.Errorf("denpence session service required")
	}
	if pkg.IP2Region == nil {
		return nil, fmt.Errorf("denpence ip2region service required")
	}
	c := cache.C()
	if c == nil {
		return nil, fmt.Errorf("denpence cache service is nil")
	}

	return &checker{
		domain:   pkg.Domain,
		user:     pkg.User,
		session:  pkg.Session,
		cache:    c,
		ip2Regin: pkg.IP2Region,
		log:      zap.L().Named("Login Security"),
	}, nil
}

type checker struct {
	domain   domain.Service
	user     user.Service
	session  session.Service
	cache    cache.Cache
	ip2Regin ip2region.Service
	log      logger.Logger
}

func (c *checker) MaxFailedRetryCheck(req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySetting(req)
	if !ss.LoginSecurity.RetryLock {
		c.log.Debugf("retry lock check disabled, don't check")
		return nil
	}

	var count uint
	err := c.cache.Get(req.AbnormalUserCheckKey(), count)
	if err != nil {
		c.log.Errorf("get key %s from cache error, %s", req.AbnormalUserCheckKey())
	}

	rc := ss.LoginSecurity.RetryLockConfig
	if count > rc.RetryLimite {
		return fmt.Errorf("retry %d times, reach the max(%d) retry limit", count, rc.RetryLimite)
	}

	return nil
}

func (c *checker) UpdateFailedRetry(req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySetting(req)
	if !ss.LoginSecurity.RetryLock {
		c.log.Debugf("retry lock check disabled, don't check")
		return nil
	}

	var count int
	if c.cache.IsExist(req.AbnormalUserCheckKey()) {
		// 之前已经登陆失败过
		err := c.cache.Put(req.AbnormalUserCheckKey(), count+1)
		if err != nil {
			c.log.Errorf("set key %s to cache error, %s", req.AbnormalUserCheckKey())
		}
	} else {
		// 首次登陆失败
		err := c.cache.PutWithTTL(
			req.AbnormalUserCheckKey(),
			count+1,
			ss.LoginSecurity.RetryLockConfig.LockedMiniteDuration(),
		)
		if err != nil {
			c.log.Errorf("set key %s to cache error, %s", req.AbnormalUserCheckKey())
		}
	}
	return nil
}

func (c *checker) OtherPlaceLoggedInChecK(req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySetting(req)
	if !ss.LoginSecurity.ExceptionLock {
		c.log.Debugf("exception check disabled, don't check")
		return nil
	}

	if !ss.LoginSecurity.ExceptionLockConfig.OtherPlaceLogin {
		c.log.Debugf("other place login check disabled, don't check")
		return nil
	}

	// 查询当前登陆IP地域
	login, err := c.ip2Regin.LookupIP(req.GetRemoteIP())
	if err != nil {
		return err
	}

	// 查询出用户上次登陆的地域
	sessSet, err := c.session.QuerySession(session.NewQuerySessionRequest(request.NewPageRequest(1, 1)))
	if err != nil {
		return err
	}

	if sessSet.Length() > 0 {
		last := sessSet.Items[0]
		c.log.Debugf("user last login city: %s (%d)", last.City, last.CityID)
		if login.CityID != last.CityID {
			return exception.NewOtherPlaceLoggedIn("异地登录, 请输入验证码后再次提及")
		}
	}

	return nil
}

func (c *checker) NotLoginDaysChecK(req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySetting(req)
	if !ss.LoginSecurity.ExceptionLock {
		c.log.Debugf("exception check disabled, don't check")
		return nil
	}

	// 查询出用户上次登陆的地域
	sessSet, err := c.session.QuerySession(session.NewQuerySessionRequest(request.NewPageRequest(1, 1)))
	if err != nil {
		return err
	}

	if sessSet.Length() > 0 {
		last := sessSet.Items[0]
		c.log.Debugf("user last login city: %s (%d)", last.City, last.CityID)
	}

	return nil
}

func (c *checker) IPProtectCheck(req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySetting(req)
	if !ss.LoginSecurity.IPLimite {
		c.log.Debugf("ip limite check disabled, don't check")
		return nil
	}

	return nil
}

func (c *checker) getOrDefaultSecuritySetting(req *token.IssueTokenRequest) *domain.SecuritySetting {
	ss := domain.NewDefaultSecuritySetting()
	u, err := c.user.DescribeAccount(user.NewDescriptAccountRequestWithAccount(req.Username))
	if err != nil {
		c.log.Errorf("get user account error, %s, use default setting to check", err)
		return ss
	}

	d, err := c.domain.DescriptionDomain(domain.NewDescriptDomainRequestWithName(u.Domain))
	if err != nil {
		c.log.Errorf("get domain error, %s, use default setting to check", err)
		return ss
	}

	return d.SecuritySetting
}
