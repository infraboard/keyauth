package grpc

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/role"
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

func newQueryRoleRequest(req *role.QueryRoleRequest) (*queryRoleRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryRoleRequest{
		QueryRoleRequest: req}, nil
}

type queryRoleRequest struct {
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
			bson.M{"type": role.RoleType_CUSTOM, "domain": r.Domain},
		}
	}

	return filter
}

func newQueryPermissionRequest(req *role.QueryPermissionRequest) (*queryPermissionRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryPermissionRequest{
		QueryPermissionRequest: req}, nil
}

type queryPermissionRequest struct {
	*role.QueryPermissionRequest
}

func (r *queryPermissionRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryPermissionRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.RoleId != "" {
		filter["role_id"] = r.RoleId
	}

	return filter
}

func newDeletePermissionRequest(req *role.RemovePermissionFromRoleRequest) (*deletePermissionRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &deletePermissionRequest{
		RemovePermissionFromRoleRequest: req}, nil
}

type deletePermissionRequest struct {
	*role.RemovePermissionFromRoleRequest
}

func (r *deletePermissionRequest) FindFilter() bson.M {
	filter := bson.M{}

	filter["role_id"] = r.RoleId
	if !r.RemoveAll {
		filter["_id"] = bson.M{"$in": r.PermissionId}
	}

	return filter
}

func newDescribePermissionRequest(req *role.DescribePermissionRequest) (*describePermissionRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &describePermissionRequest{req}, nil
}

type describePermissionRequest struct {
	*role.DescribePermissionRequest
}

func (req *describePermissionRequest) String() string {
	return fmt.Sprintf("permission: %s", req.Id)
}

func (req *describePermissionRequest) FindFilter() bson.M {
	filter := bson.M{}

	if req.Id != "" {
		filter["_id"] = req.Id
	}

	return filter
}

// FindOptions todo
func (req *describePermissionRequest) FindOptions() *options.FindOneOptions {
	opt := &options.FindOneOptions{}

	return opt
}
