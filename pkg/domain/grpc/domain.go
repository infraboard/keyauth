package grpc

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/common/types"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
)

func (s *service) CreateDomain(ctx context.Context, req *domain.CreateDomainRequest) (*domain.Domain, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	d, err := domain.New(tk.Account, req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	if _, err := s.col.InsertOne(context.TODO(), d); err != nil {
		return nil, exception.NewInternalServerError("inserted a domain document error, %s", err)
	}

	return d, nil
}

func (s *service) DescribeDomain(ctx context.Context, req *domain.DescribeDomainRequest) (*domain.Domain, error) {
	r, err := newDescDomainRequest(req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	d := domain.NewDefault()
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(d); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("domain %s not found", req)
		}

		return nil, exception.NewInternalServerError("find domain %s error, %s", req.Name, err)
	}

	return d, nil
}

func (s *service) QueryDomain(ctx context.Context, req *domain.QueryDomainRequest) (*domain.Set, error) {
	r := newQueryDomainRequest(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find domain error, error is %s", err)
	}

	domainSet := domain.NewDomainSet()
	// 循环
	for resp.Next(context.TODO()) {
		d := new(domain.Domain)
		if err := resp.Decode(d); err != nil {
			return nil, exception.NewInternalServerError("decode domain error, error is %s", err)
		}

		domainSet.Add(d)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	domainSet.Total = count

	return domainSet, nil
}

func (s *service) UpdateDomain(ctx context.Context, req *domain.UpdateDomainInfoRequest) (*domain.Domain, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	d, err := s.DescribeDomain(ctx, domain.NewDescribeDomainRequestWithName(req.Name))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case types.UpdateMode_PUT:
		*d.Profile = *req.Profile
	case types.UpdateMode_PATCH:
		d.Profile.Patch(req.Profile)
	default:
		return nil, exception.NewBadRequest("unknown update mode: %s", req.UpdateMode)
	}

	d.UpdateAt = ftime.Now().Timestamp()
	_, err = s.col.UpdateOne(context.TODO(), bson.M{"_id": d.Name}, bson.M{"$set": d})
	if err != nil {
		return nil, exception.NewInternalServerError("update domain(%s) error, %s", d.Name, err)
	}

	return d, nil
}

func (s *service) UpdateDomainSecurity(ctx context.Context, req *domain.UpdateDomainSecurityRequest) (*domain.SecuritySetting, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	d, err := s.DescribeDomain(ctx, domain.NewDescribeDomainRequestWithName(req.Name))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case types.UpdateMode_PUT:
		*d.SecuritySetting = *req.SecuritySetting
	case types.UpdateMode_PATCH:
		d.SecuritySetting.Patch(req.SecuritySetting)
	default:
		return nil, exception.NewBadRequest("unknown update mode: %s", req.UpdateMode)
	}

	d.UpdateAt = ftime.Now().Timestamp()
	_, err = s.col.UpdateOne(context.TODO(), bson.M{"_id": d.Name}, bson.M{"$set": d})
	if err != nil {
		return nil, exception.NewInternalServerError("update domain(%s) error, %s", d.Name, err)
	}

	return d.SecuritySetting, nil
}

func (s *service) DeleteDomain(ctx context.Context, req *domain.DeleteDomainRequest) (*domain.Domain, error) {
	result, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": req.Name})
	if err != nil {
		return nil, exception.NewInternalServerError("delete domain(%s) error, %s", req.Name, err)
	}

	if result.DeletedCount == 0 {
		return nil, exception.NewNotFound("domain %s not found", req.Name)
	}

	return nil, nil
}
