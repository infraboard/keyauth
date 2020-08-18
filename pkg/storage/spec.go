package storage

import (
	"fmt"
	"io"
	"net/http"

	"github.com/infraboard/keyauth/pkg/token"
)

// Service 存储服务
type Service interface {
	UploadFile(*UploadFileRequest) error
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
	BucketName string
	FileName   string

	*token.Session
	reader io.ReadCloser
	meta   map[string]string
}

// Validate 输入参数校验
func (req *UploadFileRequest) Validate() error {
	if req.BucketName == "" || req.FileName == "" {
		return fmt.Errorf("bucket name or file name is \"\"")
	}

	if req.reader == nil {
		return fmt.Errorf("object file reader is nil")
	}

	return nil
}

// Meta 文件meta
func (req *UploadFileRequest) Meta() map[string]string {
	return req.meta
}

// ReadCloser todo
func (req *UploadFileRequest) ReadCloser() io.ReadCloser {
	return req.reader
}
