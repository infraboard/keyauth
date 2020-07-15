package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/application"
)

func (s *service) CreateUserApplication(req *application.CreateApplicatonRequest) (
	*application.Application, error) {
	userID := req.GetToken().UserID
	app, err := application.NewUserApplicartion(userID, req)
	if err != nil {
		return nil, err
	}

	return s.save(app)
}

func (s *service) DescriptionApplication(req *application.DescriptApplicationRequest) (
	*application.Application, error) {
	r, err := newDescribeQuery(req)
	if err != nil {
		return nil, err
	}

	app := new(application.Application)
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(app); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("applicaiton %s not found", req)
		}

		return nil, exception.NewInternalServerError("find application %s error, %s", req.ID, err)
	}

	return app, nil
}

func (s *service) QueryApplication(req *application.QueryApplicationRequest) (*application.Set, error) {
	r := newPaggingQuery(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find domain error, error is %s", err)
	}

	appSet := application.NewApplicationSet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		app := new(application.Application)
		if err := resp.Decode(app); err != nil {
			return nil, exception.NewInternalServerError("decode domain error, error is %s", err)
		}

		appSet.Add(app)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	appSet.Total = count

	return appSet, nil
}

func (s *service) DeleteApplication(id string) error {
	result, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return exception.NewInternalServerError("delete application(%s) error, %s", id, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("app %s not found", id)
	}
	return nil
}
