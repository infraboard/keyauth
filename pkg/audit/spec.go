package audit

import (
	"fmt"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/http/request"
)

// Service todo
type Service interface {
	LoginAudit
}

// LoginAudit 登录日志审计
type LoginAudit interface {
	SaveLoginRecord(*LoginLogData)
	QueryLoginRecord(*QueryLoginRecordRequest) (*LoginRecordSet, error)
}

// NewQueryLoginRecordRequest 列表查询请求
func NewQueryLoginRecordRequest(pageReq *request.PageRequest) *QueryLoginRecordRequest {
	return &QueryLoginRecordRequest{
		Session:     token.NewSession(),
		PageRequest: pageReq,
	}
}

// NewQueryLoginRecordRequestFromData 列表查询请求
func NewQueryLoginRecordRequestFromData(req *LoginLogData) *QueryLoginRecordRequest {
	sucess := Success
	return &QueryLoginRecordRequest{
		Session:       token.NewSession(),
		PageRequest:   request.NewPageRequest(1, 1),
		Account:       req.Account,
		ApplicationID: req.ApplicationID,
		GrantType:     req.GrantType,
		Result:        &sucess,
	}
}

// QueryLoginRecordRequest todo
type QueryLoginRecordRequest struct {
	*token.Session
	*request.PageRequest
	Account       string
	ApplicationID string
	GrantType     token.GrantType
	Result        *Result
}

// Validate todo
func (req *QueryLoginRecordRequest) Validate() error {
	if req.GetToken() == nil {
		return fmt.Errorf("token required")
	}

	return nil
}
