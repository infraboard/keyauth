package audit

// Service todo
type Service interface {
	LoginAudit
}

// LoginAudit 登录日志审计
type LoginAudit interface {
	SaveLoginRecord(*LoginLogData) error
}
