package geoip

import (
	"io"
	"net"
)

// Service todo
type Service interface {
	UploadDBFile(reader io.ReadCloser) error
	Lookup(ipAddress net.IP) (*Record, error)
}
