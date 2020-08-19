package ip2region

import (
	"fmt"
	"io"
	"net/http"

	"github.com/infraboard/keyauth/pkg/token"
)

// Service todo
type Service interface {
	UpdateDBFile(*UpdateDBFileRequest) error
	LookupIP(ip string) (*IPInfo, error)
}

// IPInfo todo
type IPInfo struct {
	CityID   int64  `json:"city_id,omitempty"`
	Country  string `json:"country,omitempty"`
	Region   string `json:"region,omitempty"`
	Province string `json:"province,omitempty"`
	City     string `json:"city,omitempty"`
	ISP      string `json:"isp,omitempty"`
}

// NewUploadFileRequestFromHTTP todo
func NewUploadFileRequestFromHTTP(r *http.Request) (*UpdateDBFileRequest, error) {
	req := &UpdateDBFileRequest{
		reader:  r.Body,
		Session: token.NewSession(),
	}
	return req, nil
}

// UpdateDBFileRequest 上传文件请求
type UpdateDBFileRequest struct {
	*token.Session
	reader io.ReadCloser
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
