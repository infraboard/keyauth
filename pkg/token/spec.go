package token

// Service token管理服务
type Service interface {
	IssueToken(req *IssueTokenRequest) (*Token, error)
	DescribeToken(accessToken string) (*Token, error)
	RevolkToken(accessToken string) error
}
