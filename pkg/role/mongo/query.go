package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/token"
)

func newDescribeRoleRequest(req *role.DescribeRoleRequest) (*describeRoleRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &describeRoleRequest{req}, nil
}

type describeRoleRequest struct {
	*role.DescribeRoleRequest
}

func (req *describeRoleRequest) String() string {
	return fmt.Sprintf("role: %s", req.Name)
}

func (req *describeRoleRequest) FindFilter() bson.M {
	filter := bson.M{}

	if req.Id != "" {
		filter["_id"] = req.Id
	}

	if req.Name != "" {
		filter["name"] = req.Name
	}

	return filter
}

// FindOptions todo
func (req *describeRoleRequest) FindOptions() *options.FindOneOptions {
	opt := &options.FindOneOptions{}

	return opt
}

func newQueryRoleRequest(tk *token.Token, req *role.QueryRoleRequest) (*queryRoleRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryRoleRequest{
		tk:               tk,
		QueryRoleRequest: req}, nil
}

type queryRoleRequest struct {
	tk *token.Token
	*role.QueryRoleRequest
}

func (r *queryRoleRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	if !r.WithPermissions {
		opt.Projection = bson.M{"permissions": 0}
	}

	return opt
}

func (r *queryRoleRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Type != role.RoleType_NULL {
		filter["type"] = r.Type
	} else {
		// 获取内建和全局的角色以及域内自己创建的角色
		filter["$or"] = bson.A{
			bson.M{"type": role.RoleType_BUILDIN},
			bson.M{"type": role.RoleType_GLOBAL},
			bson.M{"type": role.RoleType_CUSTOM, "domain": r.tk.Domain},
		}
	}

	return filter
}
