package geoip

import (
	"io"
	"net"
	"net/http"

	"github.com/infraboard/keyauth/pkg/token"
)

// Service todo
type Service interface {
	UpdateDBFile(*UpdateDBFileRequest) error
	LookupIP(ipAddress net.IP) (*Record, error)
}

// NewUploadFileRequestFromHTTP todo
func NewUploadFileRequestFromHTTP(r *http.Request) *UpdateDBFileRequest {
	return &UpdateDBFileRequest{
		reader:  r.Body,
		Session: token.NewSession(),
	}
}

// UpdateDBFileRequest 上传文件请求
type UpdateDBFileRequest struct {
	*token.Session
	reader io.ReadCloser
}

// ReadCloser todo
func (req *UpdateDBFileRequest) ReadCloser() io.ReadCloser {
	return req.reader
}
