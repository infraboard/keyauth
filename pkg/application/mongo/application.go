package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/infraboard/keyauth/pkg/application"
)

func (s *service) CreateUserApplication(userID string, req *application.CreateApplicatonRequest) (
	*application.Application, error) {
	app, err := application.NewUserApplicartion(userID, application.Public, req)
	if err != nil {
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), app); err != nil {
		return nil, exception.NewInternalServerError("inserted application(%s) document error, %s",
			req.Name, err)
	}

	return app, nil
}

func (s *service) DeleteApplication(id string) error {
	_, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return exception.NewInternalServerError("delete application(%s) error, %s", id, err)
	}
	return nil
}
