package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/keyauth/apps/otp"
	"github.com/infraboard/keyauth/apps/user"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateOTPAuth(ctx context.Context, req *otp.CreateOTPAuthRequest) (*otp.OTPAuth, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := otp.NewOTPAuth()
	ins.Account = req.Account
	ins.GenSecret()
	ins.GenOtpCode(ins.Account, ins.SecretKey)
	ins.GenQrcodeUrl(ins.Account, ins.SecretKey)

	_, err := s.col.InsertOne(context.Background(), ins)
	if err != nil {
		return nil, exception.NewInternalServerError("insert document(%s) error, %s", ins, err)
	}
	return ins, nil
}

func (s *service) DescribeOTPAuth(ctx context.Context, req *otp.DescribeOTPAuthRequest) (*otp.OTPAuth, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	r := newdescribeOTPAuthRequest(req)
	ins := otp.NewOTPAuth()

	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("otp for %s not found", req.Account)
		}
		return nil, exception.NewInternalServerError("find otp %s error, %s", req.Account, err)
	}

	return ins, nil
}

func (s *service) DeleteOTPAuth(ctx context.Context, req *otp.DeleteOTPAuthRequest) (*otp.OTPAuth, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	r := newdeleteOTPAuthRequest(req)

	ins, err := s.DescribeOTPAuth(ctx, otp.NewDescribeOTPAuthRequestWithName(req.Account))
	if err != nil {
		return nil, err
	}

	_, err = s.col.DeleteOne(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("delete document(%s) otp error, %s", r.Account, err)
	}
	return ins, nil
}

func (s *service) UpdateOTPAuthStatus(ctx context.Context, req *otp.UpdateOTPStatusRequest) (*otp.OTPAuth, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	u, err := s.user.DescribeAccount(ctx, user.NewDescriptAccountRequestWithAccount(req.Account))
	if err != nil {
		return nil, err
	}

	if u.OtpStatus == req.OtpStatus {
		return nil, fmt.Errorf("do not need to update OTP status")
	}

	u.OtpStatus = req.OtpStatus
	updateOTPReq := user.NewUpdateOTPStatusRequest()
	updateOTPReq.Account = req.Account
	updateOTPReq.OtpStatus = req.OtpStatus
	_, err = s.user.UpdateOTPStatus(ctx, updateOTPReq)
	if err != nil {
		return nil, err
	}

	switch req.OtpStatus {
	case otp.OTPStatus_DISABLED:
		ins, err := s.DeleteOTPAuth(ctx, otp.NewDeleteOTPAuthRequestWithName(req.Account))
		if err != nil {
			return nil, err
		}
		return ins, nil
	case otp.OTPStatus_ENABLED:
		ins, err := s.CreateOTPAuth(ctx, otp.NewCreateOTPAuthRequestWithName(req.Account))
		if err != nil {
			return nil, err
		}
		return ins, nil
	default:
		return nil, fmt.Errorf("unknown OTPStatus type")
	}
}
