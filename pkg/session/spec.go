package session

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/token/session"
)

// Service todo
type Service interface {
	UserService
	AdminService
}

// UserService 用户端接口
type UserService interface {
	Login(*token.Token) (*Session, error)
	Logout(*LogoutRequest) error
	DescribeSession(*DescribeSessionRequest) (*Session, error)
	QuerySession(*QuerySessionRequest) (*Set, error)
}

// AdminService admin接口
type AdminService interface {
	QueryUserLastSession(*QueryUserLastSessionRequest) (*Session, error)
}

// NewQuerySessionRequestFromHTTP 列表查询请求
func NewQuerySessionRequestFromHTTP(r *http.Request) (*QuerySessionRequest, error) {
	page := request.NewPageRequestFromHTTP(r)
	qs := r.URL.Query()

	req := &QuerySessionRequest{
		Session:       session.NewSession(),
		PageRequest:   page,
		Account:       qs.Get("account"),
		ApplicationID: qs.Get("application_id"),
		LoginIP:       qs.Get("login_ip"),
		LoginCity:     qs.Get("login_city"),
	}

	gtStr := qs.Get("grant_type")
	if gtStr != "" {
		req.GrantType = token.ParseGrantTypeFromString(gtStr)
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

// NewQuerySessionRequest 列表查询请求
func NewQuerySessionRequest(pageReq *request.PageRequest) *QuerySessionRequest {
	return &QuerySessionRequest{
		Session:     session.NewSession(),
		PageRequest: pageReq,
	}
}

// NewQuerySessionRequestFromToken 列表查询请求
func NewQuerySessionRequestFromToken(tk *token.Token) *QuerySessionRequest {
	return &QuerySessionRequest{
		Session:       session.NewSession(),
		PageRequest:   request.NewPageRequest(1, 1),
		Account:       tk.Account,
		ApplicationID: tk.ApplicationId,
		GrantType:     tk.GrantType,
	}
}

// QuerySessionRequest todo
type QuerySessionRequest struct {
	*session.Session
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
func (req *QuerySessionRequest) Validate() error {
	if req.GetToken() == nil {
		return fmt.Errorf("token required")
	}

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
		SessionID: id,
	}
}

// DescribeSessionRequest todo
type DescribeSessionRequest struct {
	SessionID string
	Domain    string
	Account   string
	Login     bool
}

// Validate todo
func (req *DescribeSessionRequest) Validate() error {
	if req.SessionID == "" && !req.HasAccount() {
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
		SessionID: sessionID,
	}
}

// LogoutRequest 登出请求
type LogoutRequest struct {
	SessionID string
	Account   string
}

// NewQueryUserLastSessionRequest todo
func NewQueryUserLastSessionRequest(account string) *QueryUserLastSessionRequest {
	return &QueryUserLastSessionRequest{
		Account: account,
	}
}

// QueryUserLastSessionRequest todo
type QueryUserLastSessionRequest struct {
	Account string `validate:"required"`
}

// Validate todo
func (req *QueryUserLastSessionRequest) Validate() error {
	return validate.Struct(req)
}
