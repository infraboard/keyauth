package grpc

import (
	"context"
	"fmt"
	"strings"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	common "github.com/infraboard/keyauth/common/types"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func (s *service) QueryDepartment(ctx context.Context, req *department.QueryDepartmentRequest) (
	*department.Set, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate query department error, %s", err)
	}

	query := newQueryDepartmentRequest(req)
	set := department.NewDepartmentSet()

	if !req.SkipItems {
		s.log.Debugf("query filter: %s", query.FindFilter())
		resp, err := s.dc.Find(context.TODO(), query.FindFilter(), query.FindOptions())

		if err != nil {
			return nil, exception.NewInternalServerError("find department error, error is %s", err)
		}

		// 循环
		for resp.Next(context.TODO()) {
			ins := department.NewDefaultDepartment()
			if err := resp.Decode(ins); err != nil {
				return nil, exception.NewInternalServerError("decode department error, error is %s", err)
			}

			if req.WithSubCount {
				sc, err := s.querySubCount(ctx, ins.Id)
				if err != nil {
					return nil, err
				}
				ins.SubCount = sc
			}

			// 补充用户数量
			if req.WithUserCount {
				uc, err := s.queryUserCount(ctx, ins.Id)
				if err != nil {
					return nil, err
				}
				ins.UserCount = uc
			}

			if req.WithRole && ins.DefaultRoleId != "" {
				rIns, err := s.role.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(ins.DefaultRoleId))
				if err != nil {
					return nil, err
				}
				ins.DefaultRole = rIns
			}

			set.Add(ins)
		}
	}

	// count
	count, err := s.dc.CountDocuments(context.TODO(), query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get department count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) DescribeDepartment(ctx context.Context, req *department.DescribeDeparmentRequest) (
	*department.Department, error) {
	r, err := newDescribeDepartment(req)
	if err != nil {
		return nil, err
	}

	ins := department.NewDefaultDepartment()
	if err := s.dc.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("department %s not found", req)
		}

		return nil, exception.NewInternalServerError("find department %s error, %s", req.Id, err)
	}

	if req.WithSubCount {
		sc, err := s.querySubCount(ctx, ins.Id)
		if err != nil {
			return nil, err
		}
		ins.SubCount = sc
	}

	// 补充用户数量
	if req.WithUserCount {
		uc, err := s.queryUserCount(ctx, ins.Id)
		if err != nil {
			return nil, err
		}
		ins.UserCount = uc
	}

	if req.WithRole && ins.DefaultRoleId != "" {
		rIns, err := s.role.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(ins.DefaultRoleId))
		if err != nil {
			return nil, err
		}
		ins.DefaultRole = rIns
	}

	return ins, nil
}

func (s *service) querySubCount(ctx context.Context, departmentID string) (int64, error) {
	query := department.NewQueryDepartmentRequest()
	query.ParentId = departmentID
	query.SkipItems = true
	query.WithSubCount = true
	sc, err := s.QueryDepartment(ctx, query)
	if err != nil {
		return 0, exception.NewInternalServerError("query sub department count error, error is %s", err)
	}
	return sc.Total, nil
}

func (s *service) queryUserCount(ctx context.Context, departmentID string) (int64, error) {
	queryU := user.NewQueryAccountRequest()
	queryU.DepartmentId = departmentID
	queryU.SkipItems = true
	queryU.WithAllSub = true
	queryU.UserType = types.UserType_SUB
	us, err := s.user.QueryAccount(ctx, queryU)
	if err != nil {
		return 0, exception.NewInternalServerError("query department user count error, error is %s", err)
	}

	return us.Total, nil
}

func (s *service) CreateDepartment(ctx context.Context, req *department.CreateDepartmentRequest) (
	*department.Department, error) {
	ins, err := s.newDepartment(ctx, req)
	if err != nil {
		return nil, err
	}

	if _, err := s.dc.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted department(%s) document error, %s",
			ins.Name, err)
	}

	return ins, nil
}

func (s *service) DeleteDepartment(ctx context.Context, req *department.DeleteDepartmentRequest) (*department.Department, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate delete department request error, %s", err)
	}

	// 判断部门是否还有子部门
	desc := department.NewDescribeDepartmentRequest()
	desc.Id = req.Id
	desc.WithSubCount = true
	dep, err := s.DescribeDepartment(ctx, desc)
	if err != nil {
		return nil, err
	}
	if dep.HasSubDepartment() {
		return nil, exception.NewBadRequest("当前部门下还有%d个子部门, 请先删除", dep.SubCount)
	}

	// 判断部门下是否还有用户
	userReq := user.NewQueryAccountRequest()
	userReq.SkipItems = true
	userReq.DepartmentId = req.Id
	userReq.UserType = types.UserType_SUB
	userSet, err := s.user.QueryAccount(ctx, userReq)

	if err != nil {
		return nil, exception.NewBadRequest("quer department user error, %s", err)
	}
	if userSet.Total > 0 {
		return nil, exception.NewBadRequest("当前部门下还有%d个用户, 请先迁移用户", userSet.Total)
	}

	result, err := s.dc.DeleteOne(context.TODO(), bson.M{"_id": req.Id})
	if err != nil {
		return nil, exception.NewInternalServerError("delete department(%s) error, %s", req.Id, err)
	}

	if result.DeletedCount == 0 {
		return nil, exception.NewNotFound("department %s not found", req.Id)
	}

	return dep, nil
}

func (s *service) UpdateDepartment(ctx context.Context, req *department.UpdateDepartmentRequest) (*department.Department, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate update department error, %s", err)
	}

	dp, err := s.DescribeDepartment(ctx, department.NewDescribeDepartmentRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case common.UpdateMode_PUT:
		dp.Update(req.Data)
	case common.UpdateMode_PATCH:
		dp.Patch(req.Data)
	default:
		return nil, exception.NewBadRequest("unknown update mode: %s", req.UpdateMode)
	}

	dp.UpdateAt = ftime.Now().Timestamp()
	_, err = s.dc.UpdateOne(context.TODO(), bson.M{"_id": dp.Id}, bson.M{"$set": dp})
	if err != nil {
		return nil, exception.NewInternalServerError("update domain(%s) error, %s", dp.Name, err)
	}

	return dp, nil
}

func (s *service) newDepartment(ctx context.Context, req *department.CreateDepartmentRequest) (*department.Department, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := &department.Department{
		CreateAt:      ftime.Now().Timestamp(),
		UpdateAt:      ftime.Now().Timestamp(),
		Grade:         1,
		Name:          req.Name,
		DisplayName:   req.DisplayName,
		ParentId:      req.ParentId,
		Manager:       req.Manager,
		DefaultRoleId: req.DefaultRoleId,
	}

	if req.ParentId != "" {
		pd, err := s.DescribeDepartment(ctx, department.NewDescribeDepartmentRequestWithID(req.ParentId))
		if err != nil {
			return nil, err
		}
		ins.ParentPath = pd.Path()
		ins.Grade = int32(len(strings.Split(pd.Path(), ".")))
	}

	if req.Manager == "" {
		ins.Manager = req.CreateBy
	}

	// 检查Role是否存在
	var err error
	if req.DefaultRoleId != "" {
		ins.DefaultRole, err = s.role.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(req.DefaultRoleId))
		if err != nil {
			return nil, err
		}
	} else {
		// 默认补充访客角色
		ins.DefaultRole, err = s.role.DescribeRole(ctx, role.NewDescribeRoleRequestWithName(role.VisitorRoleName))
		if err != nil {
			return nil, err
		}
		ins.DefaultRoleId = ins.DefaultRole.Id
	}

	// 计算ID
	count, err := s.counter.GetNextSequenceValue(ins.CounterKey())
	if err != nil {
		return nil, err
	}
	ins.Number = count.Value
	ins.Id = fmt.Sprintf("%s.%d", ins.ParentPath, ins.Number)

	return ins, nil
}
