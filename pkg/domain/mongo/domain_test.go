package mongo_test

import (
	"testing"
	"time"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/conf/mock"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/domain/mongo"
	"github.com/stretchr/testify/require"
)

func newSuit(t *testing.T) *suit {
	return &suit{
		t:     t,
		shoud: require.New(t),
	}
}

type suit struct {
	t     *testing.T
	shoud *require.Assertions

	service domain.Service
	d1      *domain.Domain
}

func (s *suit) SetUp() {
	mock.Load()

	db := conf.C().Mongo.GetDB()
	s.service = mongo.NewService(db)

	s.d1 = &domain.Domain{
		ID:          "test01",
		CreateAt:    time.Now().Unix(),
		Type:        domain.Personal,
		Name:        "test domain01",
		DisplayName: "仅仅测试",
	}
}

func (s *suit) TearDown() {

}

func (s *suit) CreateDomain() func(t *testing.T) {
	return func(t *testing.T) {
		err := s.service.CreateDomain(s.d1)
		s.shoud.NoError(err)
	}
}

func TestDomainSuit(t *testing.T) {
	suit := newSuit(t)
	suit.SetUp()
	defer suit.TearDown()

	t.Run("CreateOK", suit.CreateDomain())
}
