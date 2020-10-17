package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/mcube/exception"
)

func (s *service) updateSession(sess *session.Session) error {
	_, err := s.login.UpdateOne(context.TODO(), bson.M{"_id": sess.ID}, bson.M{"$set": sess})
	if err != nil {
		return exception.NewInternalServerError("update session(%s) error, %s", sess.ID, err)
	}

	return nil
}

func (s *service) saveSession(sess *session.Session) error {
	if _, err := s.login.InsertOne(context.TODO(), sess); err != nil {
		return exception.NewInternalServerError("inserted session document error, %s", err)
	}

	return nil
}
