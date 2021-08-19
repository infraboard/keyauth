package grpc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/keyauth/conf/mock"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/domain/grpc"
)

func newSuit(t *testing.T) *suit {
	return &suit{
		t:     t,
		shoud: assert.New(t),
	}
}

type suit struct {
	t     *testing.T
	shoud *assert.Assertions

	service   domain.DomainServiceServer
	createReq *domain.CreateDomainRequest
}

func (s *suit) SetUp() {
	mock.Load()

	svr := grpc.Service
	err := svr.Config()
	if err != nil {
		panic(err)
	}

	s.service = svr

	s.createReq = &domain.CreateDomainRequest{
		Name: "test domain01",
		Profile: &domain.DomainProfile{
			DisplayName: "仅仅测试",
		},
	}
}

func (s *suit) TearDown() {

}

func (s *suit) CreateDomain() func(t *testing.T) {
	return func(t *testing.T) {
		_, err := s.service.CreateDomain(nil, s.createReq)
		s.shoud.NoError(err)
	}
}

func TestDomainSuit(t *testing.T) {
	suit := newSuit(t)
	suit.SetUp()
	defer suit.TearDown()

	t.Run("CreateOK", suit.CreateDomain())
}
