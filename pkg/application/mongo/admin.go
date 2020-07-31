package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/application"
)

func (s *service) GetBuildInApplication(name string) (*application.Application, error) {
	app := new(application.Application)
	if err := s.col.FindOne(context.TODO(), bson.M{"name": name, "build_in": true}).Decode(app); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("applicaiton %s not found", name)
		}

		return nil, exception.NewInternalServerError("find application %s error, %s", name, err)
	}

	return app, nil
}

func (s *service) CreateBuildInApplication(req *application.CreateApplicatonRequest) (
	*application.Application, error) {
	account := req.GetToken().Account
	app, err := application.NewBuildInApplication(account, req)
	if err != nil {
		return nil, err
	}

	return s.save(app)
}
