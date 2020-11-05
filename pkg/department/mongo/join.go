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
