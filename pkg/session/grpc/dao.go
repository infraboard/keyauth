package grpc

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/infraboard/keyauth/pkg/session"
)

func (s *service) updateSession(sess *session.Session) error {
	_, err := s.col.UpdateOne(context.TODO(), bson.M{"_id": sess.Id}, bson.M{"$set": sess})
	if err != nil {
		return exception.NewInternalServerError("update session(%s) error, %s", sess.Id, err)
	}

	return nil
}

func (s *service) saveSession(sess *session.Session) error {
	if _, err := s.col.InsertOne(context.TODO(), sess); err != nil {
		return exception.NewInternalServerError("inserted session document error, %s", err)
	}

	return nil
}
