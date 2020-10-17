package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"

	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/keyauth/pkg/token"
)

func (s *service) Login(tk *token.Token) (*session.Session, error) {
	if tk.IsRefresh() {
		sess, err := s.DescribeSession(session.NewDescribeSessionRequestWithID(tk.SessionID))
		if err != nil {
			return nil, err
		}

		sess.AccessToken = tk.AccessToken
		if err := s.updateSession(sess); err != nil {
			return nil, err
		}
		return sess, nil
	}

	// 关闭之前的session
	s.closeOldSession(tk)

	sess, err := session.NewSession(s.ip, tk)
	if err != nil {
		return nil, err
	}

	if err := s.saveSession(sess); err != nil {
		return nil, err
	}

	return sess, nil
}

// 判断用户之前的会话是否正常退出
// 如果access token已经过期, 则已access token过期时间为登出数据 结束该会话
// 如果该会话的刷新token已经过期, 则已刷新结束时间为登出时间 结束该会话
// 如果token正常, 则已当前时间为登出时间 结束该会话
// 结束会话后, 禁用该token
func (s *service) closeOldSession(tk *token.Token) {
	descReq := session.NewDescribeSessionRequestWithToken(tk)
	sess, err := s.DescribeSession(descReq)
	if err != nil {
		s.log.Errorf("query session error, %s", err)
		return
	}

	// 禁用该token
	if tk.IsBlock {
		sess.LogoutAt = tk.BlockAt
	} else {
		sess.LogoutAt = tk.Block(token.Normal, "session closed by other login")
	}

	if err := s.updateSession(sess); err != nil {
		s.log.Errorf("block session error, %s", err)
	}
}

func (s *service) Logout(req *session.LogoutRequest) error {
	descReq := session.NewDescribeSessionRequestWithID(req.SessionID)
	sess, err := s.DescribeSession(descReq)
	if err != nil {
		return fmt.Errorf("query session error, %s", err)
	}

	if err := s.updateSession(sess); err != nil {
		s.log.Errorf("update session error, %s", err)
	}
	return nil
}

func (s *service) DescribeSession(*session.DescribeSessionRequest) (*session.Session, error) {
	return nil, nil
}

func (s *service) QuerySession(req *session.QuerySessionRequest) (*session.Set, error) {
	r, err := newQueryLoginLogRequest(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate query session request error, %s", err)
	}

	resp, err := s.login.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find session error, %s", err)
	}

	set := session.NewSessionSet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		ins := session.NewDefaultSession()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode session error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.login.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get session count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}
