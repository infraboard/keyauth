package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
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
		s.log.Errorf("inserted login document error, %s", err)
	}

	return sess, nil
}

func (s *service) Logout(tk *token.Token) (*session.Session, error) {
	queryReq := session.NewQuerySessionRequestFromToken(tk)
	set, err := s.QuerySession(queryReq)
	if err != nil {
		return nil, fmt.Errorf("login record query error, %s", err)
	}

	if set.IsEmpty() {
		return nil, exception.NewBadRequest("login record ont found")
	}

	sess := set.Items[0]
	sess.LogoutAt = ftime.Now()

	if tk.CheckRefreshIsExpired() {
		sess.LogoutAt = tk.RefreshExpiredAt
	}

	if err := s.updateSession(sess); err != nil {
		s.log.Errorf("update session error, %s", err)
	}
	return nil, nil
}

func (s *service) QuerySession(req *session.QuerySessionRequest) (*session.Set, error) {
	r, err := newQueryLoginLogRequest(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate query login record request error, %s", err)
	}

	resp, err := s.login.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find login record error, error is %s", err)
	}

	set := session.NewSessionSet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		ins := session.NewDefaultSession()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode login record error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.login.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get login record count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (s *service) updateSession(sess *session.Session) error {
	_, err := s.login.UpdateOne(context.TODO(), bson.M{"_id": sess.ID}, bson.M{"$set": sess})
	if err != nil {
		return exception.NewInternalServerError("update login record(%s) error, %s", sess.ID, err)
	}

	return nil
}
