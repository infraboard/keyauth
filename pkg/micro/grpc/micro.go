package grpc

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/token"
)

func (s *service) CreateService(ctx context.Context, req *micro.CreateMicroRequest) (
	*micro.Micro, error) {
	ins, err := micro.New(req)
	if err != nil {
		return nil, err
	}

	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	ins.Creater = tk.Account
	ins.Domain = tk.Domain

	if _, err := s.scol.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a service document error, %s", err)
	}
	return ins, nil
}

func (s *service) QueryService(ctx context.Context, req *micro.QueryMicroRequest) (*micro.Set, error) {
	r := newPaggingQuery(req)
	fmt.Println(r.FindFilter())
	resp, err := s.scol.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find service error, error is %s", err)
	}

	set := micro.NewMicroSet()
	// 循环
	for resp.Next(context.TODO()) {
		ins := new(micro.Micro)
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode service error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.scol.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) DescribeService(ctx context.Context, req *micro.DescribeMicroRequest) (
	*micro.Micro, error) {
	r, err := newDescribeQuery(req)
	if err != nil {
		return nil, err
	}

	ins := new(micro.Micro)
	if err := s.scol.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("service %s not found", req)
		}

		return nil, exception.NewInternalServerError("find service %s error, %s", req, err)
	}
	return ins, nil
}

func (s *service) ValidateClientCredential(ctx context.Context, req *micro.ValidateClientCredentialRequest) (
	*micro.Micro, error) {
	descReq := micro.NewDescribeServiceRequestWithClientID(req.ClientId)
	ins, err := s.DescribeService(ctx, descReq)
	if err != nil {
		return nil, err
	}
	if err := ins.ValiateClientCredential(req.ClientSecret); err != nil {
		return nil, err
	}
	ins.Desensitize()
	return ins, nil
}

func (s *service) RefreshServiceClientSecret(ctx context.Context, req *micro.DescribeMicroRequest) (
	*micro.Micro, error) {
	ins, err := s.DescribeService(ctx, req)
	if err != nil {
		return nil, err
	}

	if !ins.ClientEnabled {
		return nil, exception.NewBadRequest("client is not enabled")
	}

	ins.ClientSecret = token.MakeBearer(32)
	ins.ClientRefreshAt = ftime.Now().Timestamp()
	if err := s.update(ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) DeleteService(ctx context.Context, req *micro.DeleteMicroRequest) (*micro.Micro, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate delete service error, %s", err)
	}

	describeReq := micro.NewDescribeServiceRequest()
	describeReq.Id = req.Id
	svr, err := s.DescribeService(ctx, describeReq)
	if err != nil {
		return nil, err
	}

	if micro.Type.IsIn(micro.Type_BUILD_IN) {
		return nil, exception.NewBadRequest("service %s is system service, your can't delete", svr.Name)
	}

	// 清除服务实体
	_, err = s.scol.DeleteOne(context.TODO(), bson.M{"_id": req.Id})
	if err != nil {
		return nil, exception.NewInternalServerError("delete service(%s) error, %s", req.Id, err)
	}

	// 删除服务注册的Endpoint
	deReq := endpoint.NewDeleteEndpointRequestWithServiceID(svr.Id)
	_, err = s.endpoint.DeleteEndpoint(ctx, deReq)
	if err != nil {
		s.log.Errorf("delete service endpoint error, %s", err)
	}

	return svr, nil
}
