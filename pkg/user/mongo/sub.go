package mongo

import "github.com/infraboard/keyauth/pkg/user"

func (s *service) CreateSubAccount(domainID string, req *user.CreateUserRequest) (*user.User, error) {
	return nil, nil
}

func (s *service) DeleteSubAccount(userID string) error {
	return nil
}

func (s *service) QuerySubAccount(req *user.QueryAccountRequest) (*user.Set, error) {
	return nil, nil
}
