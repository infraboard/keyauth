package mongo

import "github.com/infraboard/keyauth/pkg/user"

func (s *service) CreateRAMAccount(domainID string, req *user.CreateUserRequest) (*user.User, error) {
	return nil, nil
}

func (s *service) DeleteRAMAccount(userID string) error {
	return nil
}

func (s *service) QueryRAMAccount(req *user.QueryRAMAccountRequest) ([]*user.User, error) {
	return nil, nil
}
