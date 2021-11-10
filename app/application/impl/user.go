package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/app/application"
)

type userimpl struct {
	*service
	application.UnimplementedApplicationServiceServer
}

func (s *userimpl) CreateUserApplication(ctx context.Context, req *application.CreateApplicatonRequest) (
	*application.Application, error) {
	app, err := application.NewUserApplicartion(req)
	if err != nil {
		return nil, err
	}

	return s.save(app)
}

func (s *userimpl) DescribeApplication(ctx context.Context, req *application.DescribeApplicationRequest) (
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

		return nil, exception.NewInternalServerError("find application %s error, %s", req.Id, err)
	}

	return app, nil
}

func (s *userimpl) QueryApplication(ctx context.Context, req *application.QueryApplicationRequest) (*application.Set, error) {
	r := newPaggingQuery(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find domain error, error is %s", err)
	}

	appSet := application.NewApplicationSet(request.NewPageRequest(uint(req.Page.PageSize), uint(req.Page.PageNumber)))
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

func (s *userimpl) DeleteApplication(ctx context.Context, req *application.DeleteApplicationRequest) (*application.Application, error) {
	result, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": req.Id})
	if err != nil {
		return nil, exception.NewInternalServerError("delete application(%s) error, %s", req.Id, err)
	}

	if result.DeletedCount == 0 {
		return nil, exception.NewNotFound("app %s not found", req.Id)
	}
	return nil, nil
}
