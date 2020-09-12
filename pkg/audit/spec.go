package audit

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
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
		LoginIP:       qs.Get("login_ip"),
		LoginCity:     qs.Get("login_city"),
	}

	gtStr := qs.Get("grant_type")
	if gtStr != "" {
		gt, err := token.ParseGrantTypeFromString(gtStr)
		if err != nil {
			return nil, err
		}
		req.GrantType = gt
	}

	startTime := qs.Get("start_time")
	endTime := qs.Get("end_time")
	if startTime != "" {
		startTS, err := strconv.ParseInt(startTime, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse login start time error, %s", err)
		}

		if startTS != 0 {
			st := ftime.T(time.Unix(startTS/1000, 0))
			req.StartLoginTime = &st
		}
	}
	if endTime != "" {
		endTS, err := strconv.ParseInt(endTime, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse login start time error, %s", err)
		}

		if endTS != 0 {
			et := ftime.T(time.Unix(endTS/1000, 0))
			req.EndLoginTime = &et
		}
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
	return &QueryLoginRecordRequest{
		Session:       token.NewSession(),
		PageRequest:   request.NewPageRequest(1, 1),
		Account:       req.Account,
		ApplicationID: req.ApplicationID,
		GrantType:     req.GrantType,
	}
}

// QueryLoginRecordRequest todo
type QueryLoginRecordRequest struct {
	*token.Session
	*request.PageRequest
	Account        string
	LoginIP        string
	LoginCity      string
	ApplicationID  string
	GrantType      token.GrantType
	StartLoginTime *ftime.Time
	EndLoginTime   *ftime.Time
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
