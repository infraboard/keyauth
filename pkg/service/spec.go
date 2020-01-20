package service

import "github.com/infraboard/mcube/http/router"

// Service token管理服务
type Service interface {
	Registry(serviceToken string, entrySet router.EntrySet)
}
