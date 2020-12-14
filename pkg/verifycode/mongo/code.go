package mongo

import "github.com/infraboard/keyauth/pkg/verifycode"

func (s *service) IssueCode(*verifycode.IssueCodeRequest) (*verifycode.Code, error) {
	return nil, nil
}

func (s *service) CheckCode(*verifycode.CheckCodeRequest) (*verifycode.Code, error) {
	return nil, nil
}
