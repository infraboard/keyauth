package password

import (
	"regexp"
)

// Validater 校验器
type Validater interface {
	Reset(pass string)
	IncludeNumbers() bool
	IncludeLowercaseLetters() bool
	IncludeUppercaseLetters() bool
	IncludeSymbols() bool
	LengthOK(limit int) bool
}

// NewValidater 生产检查器
func NewValidater(pass string) Validater {
	return &validater{
		pass:       pass,
		numReg:     `[0-9]{1}`,
		lowerReg:   `[a-z]{1}`,
		upperReg:   `[A-Z]{1}`,
		symbolsReg: `[!@#~$%^&*()+|_]{1}`,
	}
}

// Validater 校验密码强度
type validater struct {
	pass       string
	numReg     string
	lowerReg   string
	upperReg   string
	symbolsReg string
}

func (v *validater) Reset(pass string) {
	v.pass = pass
}

// IncludeNumbers 是否包含数字
func (v *validater) IncludeNumbers() bool {
	return v.match(v.numReg)
}

// IncludeLowercaseLetters 是否包含小写字母
func (v *validater) IncludeLowercaseLetters() bool {
	return v.match(v.lowerReg)
}

// IncludeUppercaseLetters 是否包含大写字母
func (v *validater) IncludeUppercaseLetters() bool {
	return v.match(v.upperReg)
}

// IncludeSymbols 是否包含特殊字符
func (v *validater) IncludeSymbols() bool {
	return v.match(v.symbolsReg)
}

// LengthOK 长度是否合法
func (v *validater) LengthOK(limit int) bool {
	return len(v.pass) >= limit
}

func (v *validater) match(reg string) bool {
	if b, err := regexp.MatchString(reg, v.pass); !b || err != nil {
		return false
	}

	return true
}
