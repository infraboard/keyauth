package mongo

import (
	"github.com/infraboard/keyauth/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewService 基于MongoDB存储实现的service
func NewService(db *mongo.Database) domain.Service {
	return &service{
		dc: db.Collection("domain"),
	}
}

type service struct {
	dc            *mongo.Collection
	enableCache   bool
	notifyCachPre string
}

func (s *service) CreateDomain(domain *domain.Domain) error {
	return nil
}

func (s *service) UpdateDomain(domain *domain.Domain) error {
	return nil
}

func (s *service) DeleteDomainByID(id string) error {
	return nil
}

func (s *service) GetDomainByID(domainID string) (*domain.Domain, error) {
	return nil, nil
}

func (s *service) ListDomain(req *domain.Request) (domains []*domain.Domain, totalPage int64, err error) {
	return nil, 0, nil
}
