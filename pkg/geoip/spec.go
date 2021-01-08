package geoip

import (
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/infraboard/keyauth/pkg/token/session"
)

// Service todo
type Service interface {
	UpdateDBFile(*UpdateDBFileRequest) error
	LookupIP(ipAddress net.IP) (*Record, error)
}

// NewUploadFileRequestFromHTTP todo
func NewUploadFileRequestFromHTTP(r *http.Request) (*UpdateDBFileRequest, error) {
	qs := r.URL.Query()

	req := &UpdateDBFileRequest{
		reader:  r.Body,
		Session: session.NewSession(),
	}

	ctStr := qs.Get("content_type")
	if ctStr != "" {
		ct, err := ParseDBFileContentType(ctStr)
		if err != nil {
			return nil, err
		}

		req.ContentType = ct
	}

	return req, nil
}

// UpdateDBFileRequest 上传文件请求
type UpdateDBFileRequest struct {
	*session.Session
	reader io.ReadCloser

	ContentType DBFileContentType
}

// Validate 校验参数
func (req *UpdateDBFileRequest) Validate() error {
	if req.reader == nil {
		return fmt.Errorf("file reader is nil")
	}

	return nil
}

// ReadCloser todo
func (req *UpdateDBFileRequest) ReadCloser() io.ReadCloser {
	return req.reader
}
