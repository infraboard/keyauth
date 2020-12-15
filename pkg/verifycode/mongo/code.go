package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (s *service) IssueCode(req *verifycode.IssueCodeRequest) (*verifycode.Code, error) {
	code, err := verifycode.NewCode(req)
	if err != nil {
		return nil, err
	}

	// 如果是issue by pass, 这要检测
	if req.IssueType.Is(verifycode.IssueTypePass) {
		_, err = s.issuer.IssueToken(token.NewIssueTokenByPassword(
			req.ClientID,
			req.ClientSecret,
			req.Account(),
			req.Password),
		)
		if err != nil {
			return nil, err
		}
	}

	if _, err := s.col.InsertOne(context.TODO(), code); err != nil {
		return nil, exception.NewInternalServerError("inserted verify code(%s) document error, %s",
			code, err)
	}

	return code, nil
}

func (s *service) CheckCode(req *verifycode.CheckCodeRequest) (*verifycode.Code, error) {
	return nil, nil
}
