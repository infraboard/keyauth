package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	dc            *mongo.Collection
	enableCache   bool
	notifyCachPre string
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	s.dc = db.Collection("domain")
	return nil
}

func (s *service) CreateDomain(domain *domain.Domain) error {
	_, err := s.dc.InsertOne(context.TODO(), domain)
	if err != nil {
		return fmt.Errorf("inserted a domain document error, %s", err)
	}

	return nil
}

func (s *service) GetDomainByID(domainID string) (*domain.Domain, error) {
	d := new(domain.Domain)

	if err := s.dc.FindOne(context.TODO(), bson.M{"_id": domainID}).Decode(d); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("domain %s not found", domainID)
		}

		return nil, exception.NewInternalServerError("find domain %s error, %s", domainID, err)
	}

	return d, nil
}

func (s *service) ListDomain(req *domain.Request) (domains []*domain.Domain, totalPage int64, err error) {
	return nil, 0, nil
}

func (s *service) UpdateDomain(domain *domain.Domain) error {
	return nil
}

func (s *service) DeleteDomainByID(id string) error {
	return nil
}

func init() {
	pkg.RegistryService("mongo", Service)
}
