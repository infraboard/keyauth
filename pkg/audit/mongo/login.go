package mongo

import (
	"context"

	"github.com/infraboard/keyauth/pkg/audit"
	"github.com/infraboard/mcube/exception"
)

func (s *service) SaveLoginRecord(req *audit.LoginLogData) error {
	if err := req.Validate(); err != nil {
		return err
	}

	record := audit.NewLoginLog(req)
	record.ParseUserAgent()
	if err := record.ParseLoginIP(s.ip); err != nil {
		s.log.Errorf("parse login ip error, %s", err)
	}

	if _, err := s.login.InsertOne(context.TODO(), record); err != nil {
		return exception.NewInternalServerError("inserted login document error, %s", err)
	}

	return nil
}
