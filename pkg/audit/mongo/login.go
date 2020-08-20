package mongo

import (
	"context"

	"github.com/infraboard/keyauth/pkg/audit"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) SaveLoginRecord(req *audit.LoginLogData) {
	if err := req.Validate(); err != nil {
		s.log.Errorf("validate login record error, %s", err)
		return
	}

	record := audit.NewLoginLog(req)
	record.ParseUserAgent()
	if err := record.ParseLoginIP(s.ip); err != nil {
		s.log.Errorf("parse login ip(%s) error, %s", record.LoginIP, err)
	}

	switch req.ActionType() {
	case audit.LoginAction:
		if _, err := s.login.InsertOne(context.TODO(), record); err != nil {
			s.log.Errorf("inserted login document error, %s", err)
		}
	case audit.LogoutAction:
		queryReq := audit.NewQueryLoginRecordRequestFromData(req)
		queryReq.WithToken(req.GetToken())
		set, err := s.QueryLoginRecord(queryReq)
		if err != nil {
			s.log.Errorf("login record query error, %s", err)
			return
		}

		if set.IsEmpty() {
			s.log.Errorf("login record ont found")
			return
		}

		rc := set.Items[0]
		rc.LogoutAt = req.LogoutAt
		if err := s.updateLoginRecord(rc); err != nil {
			s.log.Errorf("update record error, %s", err)
		}
	}

	return
}

func (s *service) QueryLoginRecord(req *audit.QueryLoginRecordRequest) (*audit.LoginRecordSet, error) {
	r, err := newQueryLoginLogRequest(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate query login record request error, %s", err)
	}

	resp, err := s.login.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find login record error, error is %s", err)
	}

	set := audit.NewLoginRecordSet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		ins := audit.NewDefaultLoginLog()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode login record error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.login.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get login record count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (s *service) updateLoginRecord(rd *audit.LoginLog) error {
	_, err := s.login.UpdateOne(context.TODO(), bson.M{"_id": rd.ID}, bson.M{"$set": rd})
	if err != nil {
		return exception.NewInternalServerError("update login record(%s) error, %s", rd.ID, err)
	}

	return nil
}
