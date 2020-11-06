package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/infraboard/keyauth/pkg/department"
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

	ins, err := department.NewApplicationForm(req)
	if err != nil {
		return nil, err
	}

	if _, err := s.dc.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted join_apply(%s) document error, %s",
			ins.Creater, err)
	}

	return ins, nil
}

func (s *service) DealApplicationForm(req *department.DealApplicationFormRequest) (*department.ApplicationForm, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate deal application form request error, %s", err)
	}

	query := department.NewQueryApplicationFormRequet()
	query.WithTokenGetter(req)
	query.Account = req.Account

	as, err := s.QueryApplicationForm(query)
	if err != nil {
		return nil, err
	}
	if as.Length() == 0 {
		return nil, exception.NewBadRequest("application form not found")
	}

	if as.Length() > 1 {
		s.log.Errorf("there has more than one application form to account, only deal first one")
	}

	// 判断用户申请的部门是否还存在
	af := as.Items[0]
	dp, err := s.DescribeDepartment(department.NewDescriptDepartmentRequestWithID(af.DepartmentID))
	if err != nil {
		return nil, err
	}

	// 只有部门管理员才能出来成员加入申请
	if dp.Manager != req.GetAccount() {
		return nil, exception.NewPermissionDeny("only department manger can deal join apply")
	}

	af.UpdateAt = ftime.Now()
	af.Status = req.Status
	af.Message = req.Message
	_, err = s.dc.UpdateOne(context.TODO(), bson.M{"_id": af.Account}, bson.M{"$set": af})
	if err != nil {
		return nil, exception.NewInternalServerError("update account(%s) application form  error, %s", af.Account, err)
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
		resp, err := s.dc.Find(context.TODO(), query.FindFilter(), query.FindOptions())

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
	count, err := s.dc.CountDocuments(context.TODO(), query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get application form count error, error is %s", err)
	}
	set.Total = count
	return nil, nil
}
