package storage

import (
	"fmt"
	"io"
	"net/http"

	"github.com/infraboard/keyauth/pkg/token/session"
)

// Service 存储服务
type Service interface {
	UploadFile(*UploadFileRequest) error
	Download(*DownloadFileRequest) error
}

// NewUploadFileRequestFromHTTP todo
func NewUploadFileRequestFromHTTP(r *http.Request) *UploadFileRequest {
	return &UploadFileRequest{
		reader:  r.Body,
		meta:    make(map[string]string),
		Session: session.NewSession(),
	}
}

// NewUploadFileRequest todo
func NewUploadFileRequest(bucketName, fileName string, file io.ReadCloser) *UploadFileRequest {
	return &UploadFileRequest{
		BucketName: bucketName,
		FileName:   fileName,
		reader:     file,
		meta:       make(map[string]string),
		Session:    session.NewSession(),
	}
}

// UploadFileRequest 上传文件请求
type UploadFileRequest struct {
	BucketName string
	FileName   string

	*session.Session
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

// NewDownloadFileRequest todo
func NewDownloadFileRequest(bucketName, fileID string, writer io.Writer) *DownloadFileRequest {
	return &DownloadFileRequest{
		BucketName: bucketName,
		FileID:     fileID,
		writer:     writer,
		Session:    session.NewSession(),
	}
}

// DownloadFileRequest 上传文件请求
type DownloadFileRequest struct {
	BucketName string
	FileID     string

	*session.Session
	writer io.Writer
}

// Validate 输入参数校验
func (req *DownloadFileRequest) Validate() error {
	if req.BucketName == "" || req.FileID == "" {
		return fmt.Errorf("bucket name or file name is \"\"")
	}

	if req.writer == nil {
		return fmt.Errorf("object file reader is nil")
	}

	return nil
}

// Writer todo
func (req *DownloadFileRequest) Writer() io.Writer {
	return req.writer
}
