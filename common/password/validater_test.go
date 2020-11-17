package password_test

import (
	"testing"

	"github.com/infraboard/keyauth/common/password"
	"github.com/stretchr/testify/assert"
)

func TestValidater(t *testing.T) {
	should := assert.New(t)
	v := password.NewValidater("xx1X*")
	should.True(v.LengthOK(2))
	should.True(v.IncludeLowercaseLetters())
	should.True(v.IncludeUppercaseLetters())
	should.True(v.IncludeNumbers())
	should.True(v.IncludeSymbols())
}
