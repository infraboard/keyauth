package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/endpoint"
)

func (s *service) DescribeEndpoint(req *endpoint.DescribeEndpointRequest) (
	*endpoint.Endpoint, error) {
	r, err := newDescribeEndpointRequest(req)
	if err != nil {
		return nil, err
	}

	ins := endpoint.NewDefaultEndpoint()
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("endpoint %s not found", req)
		}

		return nil, exception.NewInternalServerError("find endpoint %s error, %s", req.ID, err)
	}

	return ins, nil
}

func (s *service) QueryEndpoints(req *endpoint.QueryEndpointRequest) (
	*endpoint.Set, error) {
	r := newQueryEndpointRequest(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find endpoint error, error is %s", err)
	}

	set := endpoint.NewEndpointSet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		app := endpoint.NewDefaultEndpoint()
		if err := resp.Decode(app); err != nil {
			return nil, exception.NewInternalServerError("decode domain error, error is %s", err)
		}

		set.Add(app)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) Registry(req *endpoint.RegistryRequest) error {
	if err := req.Validate(); err != nil {
		return exception.NewBadRequest(err.Error())
	}

	tk := req.GetToken()

	endpoints := req.Endpoints(tk.Account)
	// 更新已有的记录
	news := make([]interface{}, 0, len(endpoints))
	for i := range endpoints {
		if err := s.col.FindOneAndReplace(context.TODO(), bson.M{"_id": endpoints[i].ID}, endpoints[i]).Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				news = append(news, endpoints[i])
			} else {
				return err
			}
		}
	}

	// 插入新增记录
	if len(news) > 0 {
		if _, err := s.col.InsertMany(context.TODO(), news); err != nil {
			return exception.NewInternalServerError("inserted a service document error, %s", err)
		}
	}

	return nil
}
