package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/infraboard/keyauth/common/password"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/mcube/exception"
)

// NewDefaultSecuritySetting todo
func NewDefaultSecuritySetting() *SecuritySetting {
	return &SecuritySetting{
		PasswordSecurity: NewDefaulPasswordSecurity(),
		LoginSecurity:    NewDefaultLoginSecurity(),
	}
}

// GetPasswordRepeateLimite todo
func (s *SecuritySetting) GetPasswordRepeateLimite() uint {
	if s.PasswordSecurity == nil {
		return 0
	}
	return uint(s.PasswordSecurity.RepeateLimite)
}

// Patch todo
func (s *SecuritySetting) Patch(data *SecuritySetting) {
	patchData, _ := json.Marshal(data)
	json.Unmarshal(patchData, s)
}

// NewDefaulPasswordSecurity todo
func NewDefaulPasswordSecurity() *PasswordSecurity {
	return &PasswordSecurity{
		Length:                  8,
		IncludeNumber:           true,
		IncludeLowerLetter:      true,
		IncludeUpperLetter:      false,
		IncludeSymbols:          false,
		RepeateLimite:           1,
		PasswordExpiredDays:     90,
		BeforeExpiredRemindDays: 10,
	}
}

// Validate 校验对象合法性
func (p *PasswordSecurity) Validate() error {
	return validate.Struct(p)
}

// IsPasswordExpired todo
func (p *PasswordSecurity) IsPasswordExpired(pass *user.Password) error {
	if p.PasswordExpiredDays == 0 {
		return nil
	}

	delta := p.expiredDelta(time.Unix(pass.UpdateAt/1000, 0))
	if delta > 0 {
		return exception.NewPasswordExired("password expired %d days", delta)
	}

	return nil
}

// SetPasswordNeedReset todo
func (p *PasswordSecurity) SetPasswordNeedReset(pass *user.Password) {
	// 密码用不过期, 不需要重置
	if p.PasswordExpiredDays == 0 {
		return
	}

	// 计算密码是否过期
	delta := p.expiredDelta(time.Unix(pass.UpdateAt/1000, 0))
	if delta > 0 {
		pass.SetExpired()
		return
	}

	// 计算是否即将过期, 需要用户重置
	if -delta < int(p.BeforeExpiredRemindDays) {
		pass.SetNeedReset("密码%d天后过期, 请重置密码", -delta)
	}
}

func (p *PasswordSecurity) expiredDelta(updateAt time.Time) int {
	updateBefore := uint(time.Now().Sub(updateAt).Hours() / 24)
	return int(updateBefore) - int(p.PasswordExpiredDays)
}

// Check todo
func (p *PasswordSecurity) Check(pass string) error {
	v := password.NewValidater(pass)

	if ok := v.LengthOK(int(p.Length)); !ok {
		return fmt.Errorf("password length less than %d", p.Length)
	}
	if p.IncludeNumber {
		if ok := v.IncludeNumbers(); !ok {
			return fmt.Errorf("must include numbers")
		}
	}
	if p.IncludeLowerLetter {
		if ok := v.IncludeLowercaseLetters(); !ok {
			return fmt.Errorf("must include lower letter")
		}
	}
	if p.IncludeUpperLetter {
		if ok := v.IncludeUppercaseLetters(); !ok {
			return fmt.Errorf("must include upper letter")
		}
	}
	if p.IncludeSymbols {
		if ok := v.IncludeSymbols(); !ok {
			return fmt.Errorf("must include symbols")
		}
	}

	return nil
}

// NewDefaultLoginSecurity todo
func NewDefaultLoginSecurity() *LoginSecurity {
	return &LoginSecurity{
		ExceptionLock: false,
		ExceptionLockConfig: &ExceptionLockConfig{
			OtherPlaceLogin: true,
			NotLoginDays:    30,
		},
		RetryLock: true,
		RetryLockConfig: &RetryLockConfig{
			RetryLimite:  5,
			LockedMinite: 30,
		},
		IpLimite: false,
		IpLimiteConfig: &IPLimiteConfig{
			Ip: []string{},
		},
	}
}

// LockedMiniteDuration todo
func (c *RetryLockConfig) LockedMiniteDuration() time.Duration {
	return time.Duration(c.LockedMinite) * time.Minute
}

// GenRandomPasswordConfig todo
func (p *PasswordSecurity) GenRandomPasswordConfig() password.Config {
	return password.Config{
		Length:                  int(p.Length) + 4,
		IncludeSymbols:          true,
		IncludeNumbers:          true,
		IncludeLowercaseLetters: true,
		IncludeUppercaseLetters: true,
	}
}
