package mongo

import (
	"context"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) AddPermissionToRole(ctx context.Context, req *role.AddPermissionToRoleRequest) (*role.Role, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate add permission error, %s", err)
	}

	ins, err := s.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(req.RoleId))
	if err != nil {
		return nil, err
	}

	perms := role.NewPermission(ins.Id, tk.Account, req.Permissions)
	ins.UpdateAt = ftime.Now().Timestamp()
	_, err = s.col.UpdateOne(context.TODO(),
		bson.M{"_id": ins.Id},
		bson.M{"$addToSet": bson.M{"permissions": bson.M{"$each": perms}}},
	)
	if err != nil {
		return nil, exception.NewInternalServerError("add role permission(%s) error, %s", ins.Name, err)
	}

	ins.Permissions = append(ins.Permissions, perms...)
	return ins, nil
}

func (s *service) RemovePermissionFromRole(ctx context.Context, req *role.RemovePermissionFromRoleRequest) (*role.Role, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate remove permission error, %s", err)
	}

	ins, err := s.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(req.RoleId))
	if err != nil {
		return nil, err
	}

	ins.UpdateAt = ftime.Now().Timestamp()
	result, err := s.col.UpdateOne(context.TODO(),
		bson.M{"_id": ins.Id},
		bson.M{"$pull": bson.M{"permissions": bson.M{"_id": bson.M{"$in": req.PermissionId}}}},
	)
	if err != nil {
		return nil, exception.NewInternalServerError("remove role permission(%s) error, %s", ins.Name, err)
	}

	if result.ModifiedCount == 0 {
		return nil, exception.NewBadRequest("permission not found")
	}

	return ins, nil
}
