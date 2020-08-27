package mongo

import "fmt"

// NewFailedLogin todo
func NewFailedLogin() *FailedLogin {
	return &FailedLogin{}
}

// FailedLogin 记录
type FailedLogin struct {
	Count int `json:"count"`
}

// Inc todo
func (f *FailedLogin) Inc() {
	f.Count++
}

// CheckBlook 判断是否被阻断
func (f *FailedLogin) CheckBlook() error {
	if f.Count > 5 {
		return fmt.Errorf("max retry(5)")
	}
	return nil
}
