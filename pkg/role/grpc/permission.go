package grpc

import (
	"context"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func insertDocs(perms []*role.Permission) []interface{} {
	docs := []interface{}{}
	for i := range perms {
		docs = append(docs, perms[i])
	}
	return docs
}

func (s *service) QueryPermission(ctx context.Context, req *role.QueryPermissionRequest) (*role.PermissionSet, error) {
	query, err := newQueryPermissionRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query permission filter: %s", query.FindFilter())
	resp, err := s.perm.Find(context.TODO(), query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find permissionn error, error is %s", err)
	}

	// 循环
	set := role.NewPermissionSet()
	if !req.SkipItmes {
		for resp.Next(context.TODO()) {
			ins := role.NewDeaultPermission()
			if err := resp.Decode(ins); err != nil {
				return nil, exception.NewInternalServerError("decode permission error, error is %s", err)
			}
			set.Add(ins)
		}
	}

	// count
	count, err := s.perm.CountDocuments(context.TODO(), query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get permission count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) DescribePermission(ctx context.Context, req *role.DescribePermissionRequest) (*role.Permission, error) {
	query, err := newDescribePermissionRequest(req)
	if err != nil {
		return nil, err
	}

	ins := role.NewDeaultPermission()
	if err := s.perm.FindOne(context.TODO(), query.FindFilter(), query.FindOptions()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("permission %s not found", req)
		}

		return nil, exception.NewInternalServerError("permission %s error, %s", req, err)
	}

	return ins, nil
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

	// 查询角色条目数是否超标
	queryPerm := role.NewQueryPermissionRequest(request.NewPageRequest(role.RoleMaxPermission, 1))
	queryPerm.SkipItmes = true
	ps, err := s.QueryPermission(ctx, queryPerm)
	if err != nil {
		return nil, err
	}
	if ps.Total+int64(req.Length()) > role.RoleMaxPermission {
		return nil, exception.NewBadRequest("一个角色最多可以添加%d权限条目, 当前条目数: %d, 新增条目数: %d",
			role.RoleMaxPermission, ps.Total, req.Length())
	}

	perms := role.NewPermission(ins.Id, tk.Account, req.Permissions)
	if _, err := s.perm.InsertMany(context.TODO(), insertDocs(perms)); err != nil {
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

	delReq, err := newDeletePermissionRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.perm.DeleteMany(context.TODO(), delReq.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("delete permission(%s) error, %s", req.PermissionId, err)
	}

	if resp.DeletedCount == 0 {
		return nil, exception.NewNotFound("permission(%s) not found", req.PermissionId)
	}

	set := role.NewPermissionSet()
	return set, nil
}

func (s *service) UpdatePermission(ctx context.Context, req *role.UpdatePermissionRequest) (*role.Permission, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate remove permission error, %s", err)
	}

	ins, err := s.DescribePermission(ctx, role.NewDescribePermissionRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	ins.LabelKey = req.LabelKey
	ins.MatchAll = req.MatchAll
	ins.LabelValues = req.LabelValues

	_, err = s.perm.UpdateOne(context.TODO(), bson.M{"_id": ins.Id}, bson.M{"$set": ins})
	if err != nil {
		return nil, exception.NewInternalServerError("update permission(%s) error, %s", ins.Id, err)
	}
	return ins, nil
}
