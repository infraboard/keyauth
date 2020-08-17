package mongo

import (
	"io"
	"net"

	"github.com/infraboard/keyauth/pkg/geoip"
)

func (s *service) UploadDBFile(reader io.ReadCloser) error {
	return nil
}

func (s *service) Lookup(ipAddress net.IP) (*geoip.Record, error) {
	return nil, nil
}
