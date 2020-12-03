//go:generate  mcube enum -m

package sms

const (
	// ProviderTenCent (tencent) 腾讯短信服务
	ProviderTenCent Provider = iota + 1
	// ProviderALI (ali) 阿里短信服务
	ProviderALI
)

// Provider 短信提供商
type Provider uint
