package mongo

import (
	"net"

	"github.com/infraboard/keyauth/pkg/geoip"
	"github.com/infraboard/mcube/exception"
)

func (s *service) UpdateDBFile(req *geoip.UpdateDBFileRequest) error {
	if err := req.Validate(); err != nil {
		return exception.NewBadRequest("validate update db file requrest error, %s", err)
	}

	return nil
}

func (s *service) LookupIP(ipAddress net.IP) (*geoip.Record, error) {
	return nil, nil
}
