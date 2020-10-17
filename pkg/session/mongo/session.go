package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/keyauth/pkg/token"
)

func (s *service) Login(tk *token.Token) (*session.Session, error) {
	sess, err := session.NewSession(s.ip, tk)
	if err != nil {
		return nil, err
	}

	if _, err := s.login.InsertOne(context.TODO(), sess); err != nil {
		s.log.Errorf("inserted session document error, %s", err)
	}

	return sess, nil
}

func (s *service) Logout(req *session.LogoutRequest) error {
	descReq := session.NewDescribeSessionRequestWithID(req.SessionID)
	sess, err := s.DescribeSession(descReq)
	if err != nil {
		return fmt.Errorf("query session error, %s", err)
	}

	sess.LogoutAt = req.LogoutAt

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

func (s *service) updateSession(sess *session.Session) error {
	_, err := s.login.UpdateOne(context.TODO(), bson.M{"_id": sess.ID}, bson.M{"$set": sess})
	if err != nil {
		return exception.NewInternalServerError("update session(%s) error, %s", sess.ID, err)
	}

	return nil
}
