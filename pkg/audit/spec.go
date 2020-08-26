package audit

import (
	"fmt"
	"net/http"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/http/request"
)

// Service todo
type Service interface {
	SaveLoginRecord(*LoginLogData)
	QueryLoginRecord(*QueryLoginRecordRequest) (*LoginRecordSet, error)
}

// NewQueryLoginRecordRequestFromHTTP 列表查询请求
func NewQueryLoginRecordRequestFromHTTP(r *http.Request) (*QueryLoginRecordRequest, error) {
	page := request.NewPageRequestFromHTTP(r)
	qs := r.URL.Query()

	req := &QueryLoginRecordRequest{
		Session:       token.NewSession(),
		PageRequest:   page,
		Account:       qs.Get("account"),
		ApplicationID: qs.Get("application_id"),
	}

	gtStr := qs.Get("grant_type")
	if gtStr != "" {
		gt, err := token.ParseGrantTypeFromString(gtStr)
		if err != nil {
			return nil, err
		}
		req.GrantType = gt
	}

	return req, nil
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

// NewQueryOperateRecordRequest 列表查询请求
func NewQueryOperateRecordRequest(pageReq *request.PageRequest) *QueryOperateRecordRequest {
	return &QueryOperateRecordRequest{
		Session:     token.NewSession(),
		PageRequest: pageReq,
	}
}

// QueryOperateRecordRequest todo
type QueryOperateRecordRequest struct {
	*token.Session
	*request.PageRequest
	Account       string
	ApplicationID string
	Result        *Result
}

// Validate todo
func (req *QueryOperateRecordRequest) Validate() error {
	if req.GetToken() == nil {
		return fmt.Errorf("token required")
	}

	return nil
}
