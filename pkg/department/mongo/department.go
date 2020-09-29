package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	common "github.com/infraboard/keyauth/common/types"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func (s *service) QueryDepartment(req *department.QueryDepartmentRequest) (
	*department.Set, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate query department error, %s", err)
	}

	query := newQueryDepartmentRequest(req)
	set := department.NewDepartmentSet(req.PageRequest)

	if !req.SkipItems {
		resp, err := s.col.Find(context.TODO(), query.FindFilter(), query.FindOptions())

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
				sc, err := s.querySubCount(ins.ID, req.GetToken())
				if err != nil {
					return nil, err
				}
				ins.SubCount = &sc
			}

			// 补充用户数量
			if req.WithUserCount {
				uc, err := s.queryUserCount(ins.ID, req.GetToken())
				if err != nil {
					return nil, err
				}
				ins.UserCount = &uc
			}

			set.Add(ins)
		}
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get department count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) DescribeDepartment(req *department.DescribeDeparmentRequest) (
	*department.Department, error) {
	r, err := newDescribeDepartment(req)
	if err != nil {
		return nil, err
	}

	ins := department.NewDefaultDepartment()
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("department %s not found", req)
		}

		return nil, exception.NewInternalServerError("find department %s error, %s", req.ID, err)
	}

	if req.WithSubCount {
		sc, err := s.querySubCount(ins.ID, req.GetToken())
		if err != nil {
			return nil, err
		}
		ins.SubCount = &sc
	}

	// 补充用户数量
	if req.WithUserCount {
		uc, err := s.queryUserCount(ins.ID, req.GetToken())
		if err != nil {
			return nil, err
		}
		ins.UserCount = &uc
	}

	return ins, nil
}

func (s *service) querySubCount(departmentID string, tk *token.Token) (int64, error) {
	query := department.NewQueryDepartmentRequest()
	query.WithToken(tk)
	query.ParentID = &departmentID
	query.SkipItems = true
	query.WithSubCount = true
	sc, err := s.QueryDepartment(query)
	if err != nil {
		return 0, exception.NewInternalServerError("query sub department count error, error is %s", err)
	}
	return sc.Total, nil
}

func (s *service) queryUserCount(departmentID string, tk *token.Token) (int64, error) {
	queryU := user.NewQueryAccountRequest()
	queryU.DepartmentID = departmentID
	queryU.SkipItems = true
	queryU.WithALLSub = true
	queryU.WithToken(tk)
	us, err := s.user.QueryAccount(types.SubAccount, queryU)
	if err != nil {
		return 0, exception.NewInternalServerError("query department user count error, error is %s", err)
	}

	return us.Total, nil
}

func (s *service) CreateDepartment(req *department.CreateDepartmentRequest) (
	*department.Department, error) {
	ins, err := department.NewDepartment(req, s, s.role, s.counter)
	if err != nil {
		return nil, err
	}

	if _, err := s.col.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted department(%s) document error, %s",
			ins.Name, err)
	}

	return ins, nil
}

func (s *service) DeleteDepartment(req *department.DeleteDepartmentRequest) error {
	if err := req.Validate(); err != nil {
		return exception.NewBadRequest("validate delete department request error, %s", err)
	}

	// 判断部门是否还有子部门
	desc := department.NewDescriptDepartmentRequest()
	desc.ID = req.ID
	desc.WithSubCount = true
	desc.WithTokenGetter(req)
	dep, err := s.DescribeDepartment(desc)
	if err != nil {
		return err
	}
	if dep.HasSubDepartment() {
		return exception.NewBadRequest("当前部门下还有%d个子部门, 请先删除", *dep.SubCount)
	}

	// 判断部门下是否还有用户
	userReq := user.NewQueryAccountRequest()
	userReq.SkipItems = true
	userReq.DepartmentID = req.ID
	userReq.WithTokenGetter(req)
	userSet, err := s.user.QueryAccount(types.SubAccount, userReq)

	if err != nil {
		return exception.NewBadRequest("quer department user error, %s", err)
	}
	if userSet.Total > 0 {
		return exception.NewBadRequest("当前部门下还有%d个用户, 请先迁移用户", userSet.Total)
	}

	result, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": req.ID})
	if err != nil {
		return exception.NewInternalServerError("delete department(%s) error, %s", req.ID, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("department %s not found", req.ID)
	}

	return nil
}

func (s *service) UpdateDepartment(req *department.UpdateDepartmentRequest) (*department.Department, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate update department error, %s", err)
	}

	dp, err := s.DescribeDepartment(department.NewDescriptDepartmentRequestWithID(req.ID))
	if err != nil {
		return nil, err
	}
	switch req.UpdateMode {
	case common.PutUpdateMode:
		*dp.CreateDepartmentRequest = *req.CreateDepartmentRequest
	case common.PatchUpdateMode:
		dp.CreateDepartmentRequest.Patch(req.CreateDepartmentRequest)
	default:
		return nil, exception.NewBadRequest("unknown update mode: %s", req.UpdateMode)
	}

	dp.UpdateAt = ftime.Now()
	_, err = s.col.UpdateOne(context.TODO(), bson.M{"_id": dp.ID}, bson.M{"$set": dp})
	if err != nil {
		return nil, exception.NewInternalServerError("update domain(%s) error, %s", dp.Name, err)
	}

	return dp, nil
}
