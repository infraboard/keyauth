package mongo

import "github.com/infraboard/keyauth/pkg/user"

func (s *service) CreateServiceAccount(req *user.CreateUserRequest) (*user.User, error) {
	u, err := user.New(req)
	if err != nil {
		return nil, err
	}

	u.Type = user.ServiceAccount
	if err := s.saveAccount(u); err != nil {
		return nil, err
	}

	u.HashedPassword = nil
	return u, nil
}
