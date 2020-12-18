package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (s *service) IssueCode(req *verifycode.IssueCodeRequest) (*verifycode.Code, error) {
	code, err := verifycode.NewCode(req)
	if err != nil {
		return nil, err
	}

	// 根据系统配置设置校验码过期时间
	cf, err := s.system.GetConfig()
	if err != nil {
		s.log.Errorf("get system config error, %s", err)
	} else {
		code.ExpiredMinite = cf.VerifyCode.ExpireMinutes
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
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate check code request error, %s", err)
	}

	code := verifycode.NewDefaultCode()
	if err := s.col.FindOne(context.TODO(), bson.M{"_id": req.HashID()}).Decode(code); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("code: %s  not found", req.Number)
		}

		return nil, exception.NewInternalServerError("find system config %s error, %s", req.Number, err)
	}

	// 校验Token是否过期
	if code.IsExpired() {
		return nil, exception.NewPermissionDeny("code is expired")
	}

	return code, nil
}
