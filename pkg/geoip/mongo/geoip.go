package mongo

import (
	"net"

	"github.com/infraboard/keyauth/pkg/geoip"
)

func (s *service) UpdateDBFile(req *geoip.UpdateDBFileRequest) error {
	return nil
}

func (s *service) LookupIP(ipAddress net.IP) (*geoip.Record, error) {
	return nil, nil
}
