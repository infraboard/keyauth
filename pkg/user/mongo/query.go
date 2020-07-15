package mongo

import (
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func newQueryUserRequest(req *user.QueryAccountRequest) (*queryUserRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &queryUserRequest{
		QueryAccountRequest: req,
	}, nil
}

type queryUserRequest struct {
	userType types.Type
	*user.QueryAccountRequest
}

func (r *queryUserRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{"create_at", -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryUserRequest) FindFilter() bson.M {
	tk := r.GetToken()
	filter := bson.M{
		"type":      r.userType,
		"domain_id": tk.DomainID,
	}

	if len(r.IDs) > 0 {
		filter["_id"] = bson.M{"$in": r.IDs}
	}

	return filter
}

func newDescribeRequest(req *user.DescriptAccountRequest) (*describeUserRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeUserRequest{req}, nil
}

type describeUserRequest struct {
	*user.DescriptAccountRequest
}

func (r *describeUserRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.ID != "" {
		filter["_id"] = r.ID
	}
	if r.Account != "" {
		filter["account"] = r.Account
	}

	return filter
}

func newUpdateUserRequest(u *user.User) (*updateUserRequest, error) {
	if err := u.ValidateUpdate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	return &updateUserRequest{u}, nil
}

type updateUserRequest struct {
	*user.User
}

func (r *updateUserRequest) updateData() bson.M {
	filter := bson.M{}
	filter["update_at"] = ftime.Now()

	if r.Mobile != "" {
		filter["mobile"] = r.Mobile
	}
	if r.Email != "" {
		filter["email"] = r.Email
	}
	if r.Phone != "" {
		filter["phone"] = r.Phone
	}
	if r.Address != "" {
		filter["address"] = r.Address
	}
	if r.RealName != "" {
		filter["real_name"] = r.RealName
	}
	if r.NickName != "" {
		filter["nick_name"] = r.NickName
	}
	if r.Gender != "" {
		filter["gender"] = r.Gender
	}
	if r.Avatar != "" {
		filter["avatar"] = r.Avatar
	}
	if r.Language != "" {
		filter["language"] = r.Language
	}
	if r.City != "" {
		filter["city"] = r.City
	}
	if r.Province != "" {
		filter["province"] = r.Province
	}

	return filter
}
