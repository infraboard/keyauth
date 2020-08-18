package geoip

import (
	"io"
	"net"
	"net/http"

	"github.com/infraboard/keyauth/pkg/token"
)

// Service todo
type Service interface {
	UploadDBFile(*UploadFileRequest) error
	Lookup(ipAddress net.IP) (*Record, error)
}

// NewUploadFileRequestFromHTTP todo
func NewUploadFileRequestFromHTTP(r *http.Request) *UploadFileRequest {
	return &UploadFileRequest{
		reader:  r.Body,
		meta:    make(map[string]string),
		Session: token.NewSession(),
	}
}

// UploadFileRequest 上传文件请求
type UploadFileRequest struct {
	*token.Session
	reader io.ReadCloser
	meta   map[string]string
}

// SetMeta todo
func (req *UploadFileRequest) SetMeta(key, value string) {
	req.meta[key] = value
}

// Meta todo
func (req *UploadFileRequest) Meta() map[string]string {
	tk := req.GetToken()
	if tk != nil {
		req.SetMeta("account", tk.Account)
	}
	return req.meta
}

// ReadCloser todo
func (req *UploadFileRequest) ReadCloser() io.ReadCloser {
	return req.reader
}
