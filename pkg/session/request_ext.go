package session

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg/token"
)

// NewQuerySessionRequestFromHTTP 列表查询请求
func NewQuerySessionRequestFromHTTP(r *http.Request) (*QuerySessionRequest, error) {
	page := request.NewPageRequestFromHTTP(r)
	qs := r.URL.Query()

	req := &QuerySessionRequest{
		Page:          &page.PageRequest,
		Account:       qs.Get("account"),
		ApplicationId: qs.Get("application_id"),
		LoginIp:       qs.Get("login_ip"),
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

		req.StartLoginTime = startTS / 1000
	}
	if endTime != "" {
		endTS, err := strconv.ParseInt(endTime, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse login start time error, %s", err)
		}

		req.EndLoginTime = endTS / 1000
	}

	return req, nil
}

// NewQuerySessionRequest 列表查询请求
func NewQuerySessionRequest(pageReq *request.PageRequest) *QuerySessionRequest {
	return &QuerySessionRequest{
		Page: &pageReq.PageRequest,
	}
}

// NewQuerySessionRequestFromToken 列表查询请求
func NewQuerySessionRequestFromToken(tk *token.Token) *QuerySessionRequest {
	return &QuerySessionRequest{
		Page:          &request.NewPageRequest(1, 1).PageRequest,
		Account:       tk.Account,
		ApplicationId: tk.ApplicationId,
		GrantType:     tk.GrantType,
	}
}

// Validate todo
func (req *QuerySessionRequest) Validate() error {
	return nil
}

// NewDescribeSessionRequestWithToken todo
func NewDescribeSessionRequestWithToken(tk *token.Token) *DescribeSessionRequest {
	return &DescribeSessionRequest{
		Domain:  tk.Domain,
		Account: tk.Account,
		Login:   true,
	}
}

// NewDescribeSessionRequestWithID todo
func NewDescribeSessionRequestWithID(id string) *DescribeSessionRequest {
	return &DescribeSessionRequest{
		SessionId: id,
	}
}

// Validate todo
func (req *DescribeSessionRequest) Validate() error {
	if req.SessionId == "" && !req.HasAccount() {
		return fmt.Errorf("id or (domain and account) requried")
	}

	return nil
}

// HasAccount todo
func (req *DescribeSessionRequest) HasAccount() bool {
	if req.Domain != "" && req.Account != "" {
		return true
	}

	return false
}

// NewLogoutRequest todo
func NewLogoutRequest(sessionID string) *LogoutRequest {
	return &LogoutRequest{
		SessionId: sessionID,
	}
}

// NewQueryUserLastSessionRequest todo
func NewQueryUserLastSessionRequest(account string) *QueryUserLastSessionRequest {
	return &QueryUserLastSessionRequest{
		Account: account,
	}
}

// Validate todo
func (req *QueryUserLastSessionRequest) Validate() error {
	return validate.Struct(req)
}
