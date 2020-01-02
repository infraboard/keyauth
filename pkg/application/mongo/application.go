package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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

func (s *service) DescriptionApplication(req *application.DescriptApplicationRequest) (
	*application.Application, error) {
	app := new(application.Application)

	if err := s.col.FindOne(context.TODO(), bson.M{"_id": req.ID}).Decode(app); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("applicaiton %s not found", req.ID)
		}

		return nil, exception.NewInternalServerError("find application %s error, %s", req.ID, err)
	}

	return app, nil
}

func (s *service) QueryApplication(req *application.QueryApplicationRequest) (
	apps []*application.Application, totalPage int64, err error) {
	r := newPaggingQuery(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, 0, exception.NewInternalServerError("find domain error, error is %s", err)
	}

	// 循环
	for resp.Next(context.TODO()) {
		app := new(application.Application)
		if err := resp.Decode(app); err != nil {
			return nil, 0, exception.NewInternalServerError("decode domain error, error is %s", err)
		}

		apps = append(apps, app)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, 0, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	totalPage = count

	return apps, totalPage, nil
}

func (s *service) DeleteApplication(id string) error {
	_, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return exception.NewInternalServerError("delete application(%s) error, %s", id, err)
	}
	return nil
}

func newPaggingQuery(req *application.QueryApplicationRequest) *request {
	return &request{req}
}

type request struct {
	*application.QueryApplicationRequest
}

func (r *request) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{"create_at", -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *request) FindFilter() bson.M {
	filter := bson.M{}

	return filter
}
