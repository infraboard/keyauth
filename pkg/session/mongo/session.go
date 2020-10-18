package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/mongo"

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
	s.log.Infof("user(%s) session: %s login at: %s", sess.Account, sess.ID, sess.LoginAt.T())
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

	blockReq := token.NewBlockTokenRequest(sess.AccessToken, token.OtherClientLoggedIn, "session closed by other login")
	preTK, err := s.token.BlockToken(blockReq)
	if err != nil {
		s.log.Errorf("block previous token error, %s", err)
		return
	}
	sess.LogoutAt = preTK.EndAt()

	if err := s.updateSession(sess); err != nil {
		s.log.Errorf("block session error, %s", err)
	}
	s.log.Infof("user(%s) session: %s logout at: %s", sess.Account, sess.ID, sess.LogoutAt.T())
}

func (s *service) Logout(req *session.LogoutRequest) error {
	descReq := session.NewDescribeSessionRequestWithID(req.SessionID)
	sess, err := s.DescribeSession(descReq)
	if err != nil {
		return fmt.Errorf("query session error, %s", err)
	}

	sess.LogoutAt = ftime.Now()
	if err := s.updateSession(sess); err != nil {
		s.log.Errorf("update session error, %s", err)
	}
	s.log.Infof("user(%s) session: %s logout at: %s", sess.Account, sess.ID, sess.LogoutAt.T())
	return nil
}

func (s *service) DescribeSession(req *session.DescribeSessionRequest) (*session.Session, error) {
	r, err := newDescribeSession(req)
	if err != nil {
		return nil, err
	}

	ins := session.NewDefaultSession()
	if err := s.col.FindOne(context.TODO(), r.FindFilter(), r.FindOptions()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("session %s not found", req)
		}

		return nil, exception.NewInternalServerError("find session %s error, %s", req.SessionID, err)
	}
	return ins, nil
}

func (s *service) QuerySession(req *session.QuerySessionRequest) (*session.Set, error) {
	r, err := newQueryLoginLogRequest(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate query session request error, %s", err)
	}

	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

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
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get session count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}
