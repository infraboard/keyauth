package role

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	// BuildInType 内建角色, 系统初始时创建
	BuildInType Type = iota + 1
	// GlobalType 管理员创建的一些角色, 全局可用
	GlobalType
	// CustomType 用户自定义的角色
	CustomType
)

// Type 角色类型
type Type uint

func (t Type) String() string {
	switch t {
	case BuildInType:
		return "build_in"
	case GlobalType:
		return "global"
	case CustomType:
		return "custom"
	default:
		return "unknown"
	}
}

// MarshalJSON todo
func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON todo
func (t *Type) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case "build_in":
		*t = BuildInType
	case "global":
		*t = GlobalType
	case "custom":
		*t = CustomType
	default:
		return fmt.Errorf("unknown role type %s", string(b))
	}
	return nil
}

const (
	// Allow 允许访问
	Allow EffectType = iota + 1
	// Deny 拒绝访问
	Deny
)

// EffectType 授权效力包括两种：允许（Allow）和拒绝（Deny）
type EffectType uint

func (t EffectType) String() string {
	switch t {
	case Allow:
		return "allow"
	case Deny:
		return "deny"
	default:
		return "unknown"
	}
}

// MarshalJSON todo
func (t EffectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON todo
func (t *EffectType) UnmarshalJSON(b []byte) error {
	switch strings.Trim(string(b), `"`) {
	case "allow":
		*t = Allow
	case "deny":
		*t = Deny
	default:
		return fmt.Errorf("unknown effect type %s", string(b))
	}
	return nil
}
