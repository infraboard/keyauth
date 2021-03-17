package mongo

import (
	"context"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
)

func insertDocs(perms []*role.Permission) []interface{} {
	docs := []interface{}{}
	for i := range perms {
		docs = append(docs, perms[i])
	}
	return docs
}

func (s *service) QueryPermission(ctx context.Context, req *role.QueryPermissionRequest) (*role.PermissionSet, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	query, err := newQueryPermissionRequest(tk, req)
	if err != nil {
		return nil, err
	}

	resp, err := s.col.Find(context.TODO(), query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find permissionn error, error is %s", err)
	}

	set := role.NewPermissionSet()
	// 循环
	for resp.Next(context.TODO()) {
		ins := role.NewDeaultPermission()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode permission error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get permission count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (s *service) AddPermissionToRole(ctx context.Context, req *role.AddPermissionToRoleRequest) (*role.PermissionSet, error) {
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
	if _, err := s.col.InsertMany(context.TODO(), insertDocs(perms)); err != nil {
		return nil, exception.NewInternalServerError("inserted permission(%s) document error, %s",
			perms, err)
	}

	set := role.NewPermissionSet()
	set.Items = perms
	return set, nil
}

func (s *service) RemovePermissionFromRole(ctx context.Context, req *role.RemovePermissionFromRoleRequest) (*role.PermissionSet, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate remove permission error, %s", err)
	}

	r, err := s.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(req.RoleId))
	if err != nil {
		return nil, err
	}

	resp, err := s.col.DeleteMany(context.TODO(), bson.M{"role_id": r.Id, "_id": bson.M{"$in": req.PermissionId}})
	if err != nil {
		return nil, exception.NewInternalServerError("delete permission(%s) error, %s", req.PermissionId, err)
	}

	if resp.DeletedCount == 0 {
		return nil, exception.NewNotFound("permission(%s) not found", req.PermissionId)
	}

	set := role.NewPermissionSet()
	return set, nil
}
