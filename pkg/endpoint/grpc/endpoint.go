package grpc

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/micro"
)

func (s *service) DescribeEndpoint(ctx context.Context, req *endpoint.DescribeEndpointRequest) (
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

		return nil, exception.NewInternalServerError("find endpoint %s error, %s", req.Id, err)
	}

	return ins, nil
}

func (s *service) QueryEndpoints(ctx context.Context, req *endpoint.QueryEndpointRequest) (
	*endpoint.Set, error) {
	r := newQueryEndpointRequest(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find endpoint error, error is %s", err)
	}

	set := endpoint.NewEndpointSet()
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

func (s *service) Registry(ctx context.Context, req *endpoint.RegistryRequest) (*endpoint.RegistryResponse, error) {
	rctx, err := pkg.GetGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 查询该服务
	svr, err := s.micro.DescribeService(ctx, micro.NewDescribeServiceRequestWithClientID(rctx.GetClientID()))
	if err != nil {
		return nil, err
	}
	s.log.Debugf("service %s registry endpoints", svr.Name)

	if err := svr.ValiateClientCredential(rctx.GetClientSecret()); err != nil {
		return nil, err
	}

	// 生产该服务的Endpoint
	endpoints := req.Endpoints(svr.Id)

	// 更新已有的记录
	news := make([]interface{}, 0, len(endpoints))
	for i := range endpoints {
		if err := s.col.FindOneAndReplace(context.TODO(), bson.M{"_id": endpoints[i].Id}, endpoints[i]).Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				news = append(news, endpoints[i])
			} else {
				return nil, err
			}
		}
	}

	// 插入新增记录
	if len(news) > 0 {
		if _, err := s.col.InsertMany(context.TODO(), news); err != nil {
			return nil, exception.NewInternalServerError("inserted a service document error, %s", err)
		}
	}

	return endpoint.NewRegistryResponse("ok"), nil
}

func (s *service) DeleteEndpoint(ctx context.Context, req *endpoint.DeleteEndpointRequest) (*endpoint.Endpoint, error) {
	result, err := s.col.DeleteOne(context.TODO(), bson.M{"service_id": req.ServiceId})
	if err != nil {
		return nil, exception.NewInternalServerError("delete service(%s) endpoint error, %s", req.ServiceId, err)
	}

	s.log.Infof("delete service %s endpoint success, total count: %d", req.ServiceId, result.DeletedCount)
	return nil, nil
}
