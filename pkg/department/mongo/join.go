package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	pending = department.Pending
)

func (s *service) JoinDepartment(req *department.JoinDepartmentRequest) (*department.ApplicationForm, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 检测部署是否存在
	_, err := s.DescribeDepartment(department.NewDescriptDepartmentRequestWithID(req.DepartmentID))
	if err != nil {
		return nil, err
	}

	// 一个用户只能申请加入一个部门
	query := department.NewQueryApplicationFormRequet()
	query.SkipItems = true
	query.Account = req.Account
	query.Status = &pending
	query.WithTokenGetter(req)
	as, err := s.QueryApplicationForm(query)
	if err != nil {
		return nil, err
	}
	if as.Total > 1 {
		return nil, exception.NewBadRequest("your has an join_apply pendding")
	}

	ins, err := department.NewApplicationForm(req)
	if err != nil {
		return nil, err
	}

	if _, err := s.ac.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted join_apply(%s) document error, %s",
			ins.Creater, err)
	}

	return ins, nil
}

func (s *service) DealApplicationForm(req *department.DealApplicationFormRequest) (*department.ApplicationForm, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate deal application form request error, %s", err)
	}

	descReq := department.NewDescribeApplicationFormRequetWithID(req.ID)
	descReq.WithTokenGetter(req)

	af, err := s.DescribeApplicationForm(descReq)
	if err != nil {
		return nil, err
	}

	if !af.Status.Is(department.Pending) {
		return nil, exception.NewBadRequest("application form has deal")
	}

	// 判断用户申请的部门是否还存在
	dp, err := s.DescribeDepartment(department.NewDescriptDepartmentRequestWithID(af.DepartmentID))
	if err != nil {
		return nil, err
	}

	// 只有部门管理员才能出来成员加入申请
	if dp.Manager != req.GetAccount() {
		return nil, exception.NewPermissionDeny("only department manger can deal join apply")
	}

	// 修改用户的归属部门
	u, err := s.user.DescribeAccount(user.NewDescriptAccountRequestWithAccount(af.Account))
	if err != nil {
		return nil, err
	}

	if u.HasDepartment() {
		return nil, exception.NewBadRequest("user has deparment can't join other")
	}

	u.DepartmentID = af.DepartmentID
	patchReq := user.NewPutAccountRequest()
	patchReq.Profile = u.Profile
	patchReq.WithTokenGetter(req)

	_, err = s.user.UpdateAccountProfile(patchReq)
	if err != nil {
		return nil, err
	}

	// 持久化数据
	af.UpdateAt = ftime.Now()
	af.Status = req.Status
	af.Message = req.Message
	_, err = s.ac.UpdateOne(context.TODO(), bson.M{"_id": af.ID}, bson.M{"$set": af})
	if err != nil {
		return nil, exception.NewInternalServerError("update id(%s) application form  error, %s", af.Account, err)
	}

	return af, nil
}

func (s *service) QueryApplicationForm(req *department.QueryApplicationFormRequet) (*department.ApplicationFormSet, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate query application form error, %s", err)
	}

	query := newQueryApplicationFormRequest(req)
	set := department.NewDApplicationFormSet(req.PageRequest)

	if !req.SkipItems {
		resp, err := s.ac.Find(context.TODO(), query.FindFilter(), query.FindOptions())

		if err != nil {
			return nil, exception.NewInternalServerError("find application form error, error is %s", err)
		}

		// 循环
		for resp.Next(context.TODO()) {
			ins := department.NewDeafultApplicationForm()
			if err := resp.Decode(ins); err != nil {
				return nil, exception.NewInternalServerError("decode application form error, error is %s", err)
			}

			set.Add(ins)
		}
	}

	// count
	count, err := s.ac.CountDocuments(context.TODO(), query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get application form count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (s *service) DescribeApplicationForm(req *department.DescribeApplicationFormRequet) (
	*department.ApplicationForm, error) {
	r := newDescribeApplicationForm(req)

	ins := department.NewDeafultApplicationForm()
	if err := s.ac.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("application form %s not found", req)
		}

		return nil, exception.NewInternalServerError("find application form %s error, %s", req.ID, err)
	}

	return ins, nil
}
