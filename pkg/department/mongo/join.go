package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"

	"github.com/infraboard/keyauth/pkg/department"
)

func (s *service) JoinDepartment(req *department.JoinDepartmentRequest) (*department.ApplicationForm, error) {
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
	// dp.UpdateAt = ftime.Now()
	// _, err = s.dc.UpdateOne(context.TODO(), bson.M{"_id": dp.ID}, bson.M{"$set": dp})
	// if err != nil {
	// 	return nil, exception.NewInternalServerError("update domain(%s) error, %s", dp.Name, err)
	// }
	return nil, nil
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
