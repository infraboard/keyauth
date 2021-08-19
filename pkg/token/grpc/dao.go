package grpc

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/infraboard/keyauth/pkg/token"
)

func (s *service) saveToken(tk *token.Token) error {
	if _, err := s.col.InsertOne(context.TODO(), tk); err != nil {
		return exception.NewInternalServerError("inserted token(%s) document error, %s",
			tk.AccessToken, err)
	}

	return nil
}

func (s *service) updateToken(tk *token.Token) error {
	_, err := s.col.UpdateOne(context.TODO(), bson.M{"_id": tk.AccessToken}, bson.M{"$set": tk})
	if err != nil {
		return exception.NewInternalServerError("update token(%s) error, %s", tk.AccessToken, err)
	}

	return nil
}
