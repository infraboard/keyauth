package mongo

import "github.com/infraboard/keyauth/pkg/audit"

func (s *service) SaveOperateRecord(*audit.OperateLogData) {

}

func (s *service) QueryOperateRecord(*audit.QueryOperateRecordRequest) (*audit.OperateRecordSet, error) {
	return nil, nil
}
