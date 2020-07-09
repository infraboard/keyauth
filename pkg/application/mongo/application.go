package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/application"
)

func (s *service) CreateUserApplication(req *application.CreateApplicatonRequest) (
	*application.Application, error) {
	userID := req.GetToken().UserID
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
	_, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return exception.NewInternalServerError("delete application(%s) error, %s", id, err)
	}
	return nil
}

func newPaggingQuery(req *application.QueryApplicationRequest) *queryRequest {
	return &queryRequest{req}
}

type queryRequest struct {
	*application.QueryApplicationRequest
}

func (r *queryRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{"create_at", -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryRequest) FindFilter() bson.M {
	filter := bson.M{}

	return filter
}

func newDescribeQuery(req *application.DescriptApplicationRequest) (*describeRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeRequest{req}, nil
}

type describeRequest struct {
	*application.DescriptApplicationRequest
}

func (r *describeRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.ID != "" {
		filter["_id"] = r.ID
	}
	if r.Name != "" {
		filter["name"] = r.Name
	}
	if r.ClientID != "" {
		filter["client_id"] = r.ClientID
	}

	return filter
}
