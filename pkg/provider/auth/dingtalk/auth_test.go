package dingtalk_test

import (
	"testing"

	"github.com/infraboard/keyauth/pkg/provider/auth/dingtalk"
)

func TestAuthCode(t *testing.T) {
	dd := dingtalk.Dingtalk{
		AppID:     "xxx",
		AppSecret: "xxxx",
	}
	dd.CodeAuth(nil)
}
