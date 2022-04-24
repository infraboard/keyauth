package impl

import (
	"github.com/infraboard/keyauth/apps/otp"
	"go.mongodb.org/mongo-driver/bson"
)

func newdescribeOTPAuthRequest(req *otp.DescribeOTPAuthRequest) *describeOTPAuthRequest {
	return &describeOTPAuthRequest{
		DescribeOTPAuthRequest: req,
	}
}

type describeOTPAuthRequest struct {
	*otp.DescribeOTPAuthRequest
}

func (req *describeOTPAuthRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.Account != "" {
		filter["account"] = req.Account
	}
	return filter
}
func newdeleteOTPAuthRequest(req *otp.DeleteOTPAuthRequest) *deleteOTPAuthRequest {
	return &deleteOTPAuthRequest{
		DeleteOTPAuthRequest: req,
	}
}

type deleteOTPAuthRequest struct {
	*otp.DeleteOTPAuthRequest
}

func (req *deleteOTPAuthRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.Account != "" {
		filter["account"] = req.Account
	}
	return filter
}
