package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/apps/service"
)

func (i *impl) save(ctx context.Context, ins *service.Service) error {
	if _, err := i.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted Service(%s) document error, %s",
			ins.Spec.Name, err)
	}
	return nil
}

func (i *impl) update(ctx context.Context, ins *service.Service) error {
	if _, err := i.col.UpdateByID(ctx, ins.Id, ins); err != nil {
		return exception.NewInternalServerError("inserted Service(%s) document error, %s",
			ins.Spec.Name, err)
	}

	return nil
}

func newQueryRequest(r *service.QueryServiceRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

type queryRequest struct {
	*service.QueryServiceRequest
}

func (r *queryRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort: bson.D{
			{Key: "create_at", Value: -1},
		},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryRequest) FindFilter() bson.M {
	filter := bson.M{}
	// if r.Keywords != "" {
	// 	filter["$or"] = bson.A{
	// 		bson.M{"data.name": bson.M{"$regex": r.Keywords, "$options": "im"}},
	// 		bson.M{"data.author": bson.M{"$regex": r.Keywords, "$options": "im"}},
	// 	}
	// }
	return filter
}

func (i *impl) query(ctx context.Context, req *queryRequest) (*service.ServiceSet, error) {
	resp, err := i.col.Find(ctx, req.FindFilter(), req.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find book error, error is %s", err)
	}

	ServiceSet := service.NewServiceSet()
	// 循环
	for resp.Next(ctx) {
		ins := service.NewDefaultService()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode book error, error is %s", err)
		}

		ServiceSet.Add(ins)
	}

	// count
	count, err := i.col.CountDocuments(ctx, req.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get Service count error, error is %s", err)
	}
	ServiceSet.Total = count

	return ServiceSet, nil
}

func (i *impl) get(ctx context.Context, req *service.DescribeServiceRequest) (*service.Service, error) {
	filter := bson.M{}
	switch req.DescribeBy {
	case service.DescribeBy_SERVICE_ID:
		filter["_id"] = req.Id
	case service.DescribeBy_SERVICE_CLIENT_ID:
		filter["credential.client_id"] = req.ClientId
	case service.DescribeBy_SERVICE_NAME:
		filter["spec.name"] = req.Name
	}

	ins := service.NewDefaultService()
	if err := i.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("Service %s not found", req)
		}

		return nil, exception.NewInternalServerError("find Service %s error, %s", req, err)
	}

	return ins, nil
}

func (i *impl) delete(ctx context.Context, ins *service.Service) error {
	if ins == nil || ins.Id == "" {
		return fmt.Errorf("service is nil")
	}

	result, err := i.col.DeleteOne(ctx, bson.M{"_id": ins.Id})
	if err != nil {
		return exception.NewInternalServerError("delete Service(%s) error, %s", ins.Id, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("Service %s not found", ins.Id)
	}

	return nil
}
