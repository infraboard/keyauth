package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	uc            *mongo.Collection
	enableCache   bool
	notifyCachPre string
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	s.uc = db.Collection("user")
	return nil
}

func (s *service) CreatePrimayAccount(req *user.CreateUserRequest) (*user.User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	user := user.NewUser(req)
	user.Primary = true
	_, err := s.uc.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, fmt.Errorf("inserted a user document error, %s", err)
	}

	return user, nil
}

func (s *service) DescribeAccount(req *user.DescriptAccountRequest) (*user.User, error) {
	user := user.NewDescribeUser()

	if err := s.uc.FindOne(context.TODO(), bson.M{"_id": req.ID}).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("user %s not found", req.ID)
		}

		return nil, exception.NewInternalServerError("find user %s error, %s", req.ID, err)
	}

	return user, nil
}

func (s *service) QueryDomain(req *domain.QueryDomainRequest) (domains []*domain.Domain, totalPage int64, err error) {
	r := request{QueryDomainRequest: req}
	resp, err := s.uc.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, 0, exception.NewInternalServerError("find domain error, error is %s", err)
	}

	// 循环
	for resp.Next(context.TODO()) {
		d := new(domain.Domain)
		if err := resp.Decode(d); err != nil {
			return nil, 0, exception.NewInternalServerError("decode domain error, error is %s", err)
		}

		domains = append(domains, d)
	}

	// count
	count, err := s.uc.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, 0, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	totalPage = count

	return domains, totalPage, nil
}

func (s *service) UpdateDomain(d *domain.Domain) error {
	if err := d.CreateDomainRequst.Validate(); err != nil {
		return exception.NewBadRequest(err.Error())
	}

	_, err := s.uc.UpdateOne(context.TODO(), bson.M{"_id": d.ID}, bson.M{"$set": d})
	if err != nil {
		return exception.NewInternalServerError("update domain(%s) error, %s", d.ID, err)
	}

	return nil
}

func (s *service) DeleteDomain(id string) error {
	_, err := s.uc.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return exception.NewInternalServerError("delete domain(%s) error, %s", id, err)
	}
	return nil
}

type request struct {
	*domain.QueryDomainRequest

	opt *options.FindOptions
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

func init() {
	pkg.RegistryService("mongo", Service)
}
