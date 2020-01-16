package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/domain"
)

func (s *service) CreateDomain(ownerID string, req *domain.CreateDomainRequst) (*domain.Domain, error) {
	d, err := domain.New(ownerID, req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	if _, err := s.col.InsertOne(context.TODO(), d); err != nil {
		return nil, exception.NewInternalServerError("inserted a domain document error, %s", err)
	}

	return d, nil
}

func (s *service) DescriptionDomain(req *domain.DescriptDomainRequest) (*domain.Domain, error) {
	d := new(domain.Domain)

	if err := s.col.FindOne(context.TODO(), bson.M{"_id": req.ID}).Decode(d); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("domain %s not found", req.ID)
		}

		return nil, exception.NewInternalServerError("find domain %s error, %s", req.ID, err)
	}

	return d, nil
}

func (s *service) QueryDomain(req *domain.QueryDomainRequest) (*domain.DomainSet, error) {
	r := newPaggingQuery(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find domain error, error is %s", err)
	}

	domainSet := domain.NewDomainSet(req.PageRequest)
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

func (s *service) UpdateDomain(d *domain.Domain) error {
	if err := d.CreateDomainRequst.Validate(); err != nil {
		return exception.NewBadRequest(err.Error())
	}

	_, err := s.col.UpdateOne(context.TODO(), bson.M{"_id": d.ID}, bson.M{"$set": d})
	if err != nil {
		return exception.NewInternalServerError("update domain(%s) error, %s", d.ID, err)
	}

	return nil
}

func (s *service) DeleteDomain(id string) error {
	_, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return exception.NewInternalServerError("delete domain(%s) error, %s", id, err)
	}
	return nil
}

func newPaggingQuery(req *domain.QueryDomainRequest) *request {
	return &request{req}
}

type request struct {
	*domain.QueryDomainRequest
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
