package token

// Service token管理服务
type Service interface {
	IssueToken(req *IssueTokenRequest) (*Token, error)
	ValidateToken(accessToken, endpoint string) (*Token, error)
	RevolkToken(accessToken string) error
}
