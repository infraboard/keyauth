package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
)

func (s *service) CreateNamespace(ctx context.Context, req *namespace.CreateNamespaceRequest) (
	*namespace.Namespace, error) {
	ins, err := s.newNamespace(ctx, req)
	if err != nil {
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted namespace(%s) document error, %s",
			ins.Name, err)
	}

	if err := s.updateNamespacePolicy(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) updateNamespacePolicy(ctx context.Context, ns *namespace.Namespace) error {
	descR := role.NewDescribeRoleRequestWithName(role.AdminRoleName)
	r, err := s.role.DescribeRole(ctx, descR)
	if err != nil {
		return err
	}
	pReq := policy.NewCreatePolicyRequest()
	pReq.NamespaceId = ns.Id
	pReq.RoleId = r.Id
	pReq.Account = ns.Owner
	pReq.Type = policy.PolicyType_BUILD_IN
	_, err = s.policy.CreatePolicy(ctx, pReq)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) QueryNamespace(ctx context.Context, req *namespace.QueryNamespaceRequest) (
	*namespace.Set, error) {
	r := newPaggingQuery(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find namespace error, error is %s", err)
	}

	set := namespace.NewNamespaceSet()
	// 循环
	for resp.Next(context.TODO()) {
		ins := namespace.NewDefaultNamespace()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode namespace error, error is %s", err)
		}

		// 补充用户的部门信息
		if req.WithDepartment && ins.DepartmentId != "" {
			depart, err := s.depart.DescribeDepartment(ctx, department.NewDescribeDepartmentRequestWithID(ins.DepartmentId))
			if err != nil {
				s.log.Errorf("get user department error, %s", err)
			} else {
				ins.Department = depart
			}
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get namespace count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) DescribeNamespace(ctx context.Context, req *namespace.DescriptNamespaceRequest) (
	*namespace.Namespace, error) {
	r, err := newDescribeQuery(req)
	if err != nil {
		return nil, err
	}

	ins := namespace.NewDefaultNamespace()
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("namespace %s not found", req)
		}

		return nil, exception.NewInternalServerError("find namespace %s error, %s", req.Id, err)
	}

	// 补充用户的部门信息
	if req.WithDepartment && ins.DepartmentId != "" {
		depart, err := s.depart.DescribeDepartment(ctx, department.NewDescribeDepartmentRequestWithID(ins.DepartmentId))
		if err != nil {
			s.log.Errorf("get user department error, %s", err)
		} else {
			ins.Department = depart
		}
	}

	return ins, nil
}

func (s *service) DeleteNamespace(ctx context.Context, req *namespace.DeleteNamespaceRequest) (*namespace.Namespace, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	r, err := newDeleteRequest(tk, req)
	if err != nil {
		return nil, err
	}

	result, err := s.col.DeleteOne(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("delete namespace(%s) error, %s", req.Id, err)
	}

	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("namespace %s not found", req.Id)
	}

	// 清除空间管理的所有策略
	_, err = s.policy.DeletePolicy(ctx, policy.NewDeletePolicyRequestWithNamespaceID(req.Id))
	if err != nil {
		s.log.Errorf("delete namespace policy error, %s", err)
	}

	return nil, nil
}

// NewNamespace todo
func (s *service) newNamespace(ctx context.Context, req *namespace.CreateNamespaceRequest) (*namespace.Namespace, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	ins := &namespace.Namespace{
		Id:           xid.New().String(),
		Domain:       tk.Domain,
		Creater:      tk.Account,
		CreateAt:     ftime.Now().Timestamp(),
		UpdateAt:     ftime.Now().Timestamp(),
		DepartmentId: req.DepartmentId,
		Name:         req.Name,
		Picture:      req.Picture,
		Owner:        req.Owner,
		Description:  req.Description,
	}

	descD := department.NewDescribeDepartmentRequest()
	descD.Id = req.DepartmentId
	d, err := s.depart.DescribeDepartment(ctx, descD)
	if err != nil {
		return nil, err
	}
	// 部门负责人就是空间负责人
	ins.Owner = d.Manager

	return ins, nil
}
