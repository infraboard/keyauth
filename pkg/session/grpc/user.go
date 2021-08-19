package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/keyauth/pkg/token"
)

type userimpl struct {
	*service
	session.UnimplementedUserServiceServer
}

func (s *userimpl) Login(ctx context.Context, tk *token.Token) (*session.Session, error) {
	if tk.IsRefresh() {
		sess, err := s.DescribeSession(ctx, session.NewDescribeSessionRequestWithID(tk.SessionId))
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
	s.closeOldSession(ctx, tk)

	sess, err := session.NewSession(s.ip, tk)
	if err != nil {
		return nil, err
	}

	// 填充session IP信息
	sess.IpInfo, err = s.parseRemoteIPInfo(tk.GetRemoteIp())
	if err != nil {
		s.log.Errorf("parse remote ip error, %s", err)
	}

	if err := s.saveSession(sess); err != nil {
		return nil, err
	}
	s.log.Infof("user(%s) session: %s login at: %s", sess.Account, sess.Id, sess.LoginAt)
	return sess, nil
}

func (s *userimpl) parseRemoteIPInfo(ip string) (*session.IPInfo, error) {
	if ip == "" {
		return nil, nil
	}

	info, err := s.ip.LookupIP(ip)
	if err != nil {
		return nil, fmt.Errorf("parse ipinfo error, %s", err)
	}

	return &session.IPInfo{
		CityId:   info.CityID,
		Country:  info.Country,
		Region:   info.Region,
		Province: info.Province,
		City:     info.City,
		Isp:      info.ISP,
	}, nil
}

// 判断用户之前的会话是否正常退出
// 如果access token已经过期, 则已access token过期时间为登出数据 结束该会话
// 如果该会话的刷新token已经过期, 则已刷新结束时间为登出时间 结束该会话
// 如果token正常, 则已当前时间为登出时间 结束该会话
// 结束会话后, 禁用该token
func (s *userimpl) closeOldSession(ctx context.Context, tk *token.Token) {
	descReq := session.NewDescribeSessionRequestWithToken(tk)
	sess, err := s.DescribeSession(ctx, descReq)
	if err != nil {
		s.log.Errorf("query session error, %s", err)
		return
	}

	blockReq := token.NewBlockTokenRequest(sess.AccessToken, token.BlockType_OTHER_CLIENT_LOGGED_IN, "session closed by other login")
	preTK, err := s.token.BlockToken(nil, blockReq)
	if err != nil {
		s.log.Errorf("block previous token error, %s", err)
		return
	}
	sess.LogoutAt = ftime.Time(time.Unix(preTK.EndAt()/1000, 0)).Timestamp()

	if err := s.updateSession(sess); err != nil {
		s.log.Errorf("block session error, %s", err)
	}
	s.log.Infof("user(%s) session: %s logout at: %s", sess.Account, sess.Id, sess.LogoutAt)
}

func (s *userimpl) Logout(ctx context.Context, req *session.LogoutRequest) (*session.Session, error) {
	descReq := session.NewDescribeSessionRequestWithID(req.SessionId)
	sess, err := s.DescribeSession(ctx, descReq)
	if err != nil {
		return nil, fmt.Errorf("query session error, %s", err)
	}

	sess.LogoutAt = ftime.Now().Timestamp()
	if err := s.updateSession(sess); err != nil {
		s.log.Errorf("update session error, %s", err)
	}
	s.log.Infof("user(%s) session: %s logout at: %s", sess.Account, sess.Id, sess.LogoutAt)
	return sess, nil
}

func (s *userimpl) DescribeSession(ctx context.Context, req *session.DescribeSessionRequest) (*session.Session, error) {
	r, err := newDescribeSession(req)
	if err != nil {
		return nil, err
	}

	ins := session.NewDefaultSession()
	if err := s.col.FindOne(context.TODO(), r.FindFilter(), r.FindOptions()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("session %s not found", req)
		}

		return nil, exception.NewInternalServerError("find session %s error, %s", req.SessionId, err)
	}
	return ins, nil
}

func (s *userimpl) QuerySession(ctx context.Context, req *session.QuerySessionRequest) (*session.Set, error) {
	r, err := newQueryLoginLogRequest(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate query session request error, %s", err)
	}

	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find session error, %s", err)
	}

	set := session.NewSessionSet()
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
