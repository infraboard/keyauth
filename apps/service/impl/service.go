package impl

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/infraboard/keyauth/apps/service"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
)

func (i *impl) ValidateCredential(ctx context.Context, req *service.ValidateCredentialRequest) (
	*service.Service, error) {
	svr, err := i.DescribeService(ctx, service.NewDescribeServiceRequestByClientId(req.ClientId))
	if err != nil {
		return nil, err
	}

	if err := svr.Credential.Validate(req.ClientSecret); err != nil {
		return nil, err
	}

	return svr, nil
}

func (i *impl) CreateService(ctx context.Context, req *service.CreateServiceRequest) (
	*service.Service, error) {
	ins, err := service.NewService(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create book error, %s", err)
	}

	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) UpdateService(ctx context.Context, req *service.UpdateServiceRequest) (
	*service.Service, error) {
	ins, err := i.DescribeService(ctx, service.NewDescribeServiceRequest(req.Id))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		ins.Update(req)
	case request.UpdateMode_PATCH:
		err := ins.Patch(req)
		if err != nil {
			return nil, err
		}
	}

	// 校验更新后数据合法性
	if err := ins.Spec.Validate(); err != nil {
		return nil, err
	}

	if err := i.update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) QueryService(ctx context.Context, req *service.QueryServiceRequest) (
	*service.ServiceSet, error) {
	query := newQueryRequest(req)
	return i.query(ctx, query)
}

func (i *impl) DescribeService(ctx context.Context, req *service.DescribeServiceRequest) (
	*service.Service, error) {
	return i.get(ctx, req)
}

func (i *impl) DeleteService(ctx context.Context, req *service.DeleteServiceRequest) (
	*service.Service, error) {
	ins, err := i.DescribeService(ctx, service.NewDescribeServiceRequest(req.Id))
	if err != nil {
		return nil, err
	}

	if err := i.delete(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) RefreshCredential(ctx context.Context, req *service.DescribeServiceRequest) (
	*service.Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshCredential not implemented")
}
