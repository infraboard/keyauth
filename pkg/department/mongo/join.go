package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	pending = department.ApplicationFormStatus_PENDDING
)

func (s *service) JoinDepartment(ctx context.Context, req *department.JoinDepartmentRequest) (*department.ApplicationForm, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk := session.GetTokenFromContext(ctx)
	// 检测部署是否存在
	_, err := s.DescribeDepartment(ctx, department.NewDescribeDepartmentRequestWithID(req.DepartmentId))
	if err != nil {
		return nil, err
	}

	// 一个用户只能申请加入一个部门
	query := department.NewQueryApplicationFormRequet()
	query.SkipItems = true
	query.Account = tk.Account
	query.Status = pending

	as, err := s.QueryApplicationForm(ctx, query)
	if err != nil {
		return nil, err
	}
	if as.Total > 1 {
		return nil, exception.NewBadRequest("your has an join_apply pendding")
	}

	ins, err := department.NewApplicationForm(tk, req)
	if err != nil {
		return nil, err
	}

	if _, err := s.ac.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted join_apply(%s) document error, %s",
			ins.Creater, err)
	}

	return ins, nil
}

func (s *service) DealApplicationForm(ctx context.Context, req *department.DealApplicationFormRequest) (*department.ApplicationForm, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate deal application form request error, %s", err)
	}

	tk := session.GetTokenFromContext(ctx)
	descReq := department.NewDescribeApplicationFormRequetWithID(req.Id)
	af, err := s.DescribeApplicationForm(ctx, descReq)
	if err != nil {
		return nil, err
	}

	if !af.Status.Equal(department.ApplicationFormStatus_PENDDING) {
		return nil, exception.NewBadRequest("application form has deal")
	}

	// 判断用户申请的部门是否还存在
	dp, err := s.DescribeDepartment(ctx, department.NewDescribeDepartmentRequestWithID(af.Data.DepartmentId))
	if err != nil {
		return nil, err
	}

	// 只有部门管理员才能出来成员加入申请
	if dp.Data.Manager != tk.Account {
		return nil, exception.NewPermissionDeny("only department manger can deal join apply")
	}

	// 修改用户的归属部门
	u, err := s.user.DescribeAccount(ctx, user.NewDescriptAccountRequestWithAccount(af.Data.Account))
	if err != nil {
		return nil, err
	}

	if u.HasDepartment() {
		return nil, exception.NewBadRequest("user has deparment can't join other")
	}

	u.Data.Profile.DepartmentId = af.Data.DepartmentId
	patchReq := user.NewPutAccountRequest()
	patchReq.Profile = u.Data.Profile

	_, err = s.user.UpdateAccountProfile(ctx, patchReq)
	if err != nil {
		return nil, err
	}

	// 持久化数据
	af.UpdateAt = ftime.Now().Timestamp()
	af.Status = req.Status
	af.Data.Message = req.Message
	_, err = s.ac.UpdateOne(context.TODO(), bson.M{"_id": af.Id}, bson.M{"$set": af})
	if err != nil {
		return nil, exception.NewInternalServerError("update id(%s) application form  error, %s", af.Data.Account, err)
	}

	return af, nil
}

func (s *service) QueryApplicationForm(ctx context.Context, req *department.QueryApplicationFormRequet) (*department.ApplicationFormSet, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate query application form error, %s", err)
	}
	tk := session.GetTokenFromContext(ctx)

	query := newQueryApplicationFormRequest(tk, req)
	set := department.NewDApplicationFormSet()

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

func (s *service) DescribeApplicationForm(ctx context.Context, req *department.DescribeApplicationFormRequet) (
	*department.ApplicationForm, error) {
	tk := session.GetTokenFromContext(ctx)
	r := newDescribeApplicationForm(tk, req)

	ins := department.NewDeafultApplicationForm()
	if err := s.ac.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("application form %s not found", req)
		}

		return nil, exception.NewInternalServerError("find application form %s error, %s", req.Id, err)
	}

	return ins, nil
}
