package mongo

import (
	"fmt"
	"io"

	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/storage"
	"github.com/infraboard/mcube/exception"
)

func (s *service) UploadFile(req *storage.UploadFileRequest) error {
	if err := req.Validate(); err != nil {
		return exception.NewBadRequest("valiate upload file request error, %s", err)
	}

	bucket, err := s.getBucket(req.BucketName)

	opts := options.GridFSUpload()
	opts.Metadata = req.Meta()

	// 清除已有文件
	bucket.Delete(req.FileName)

	// 上传新文件
	uploadStream, err := bucket.OpenUploadStreamWithID(req.FileName, req.FileName, opts)
	if err != nil {
		return err
	}
	defer uploadStream.Close()

	fileSize, err := io.Copy(uploadStream, req.ReadCloser())
	if err != nil {
		return err
	}

	s.log.Debugf("Write file %s to DB was successful. File size: %d M", req.FileName, fileSize/1024/1024)
	return nil
}

func (s *service) Download(req *storage.DownloadFileRequest) error {
	if err := req.Validate(); err != nil {
		return exception.NewBadRequest("valiate upload file request error, %s", err)
	}

	bucket, err := s.getBucket(req.BucketName)

	s.log.Debugf("start down file: %s ...", req.FileID)
	// 下载文件
	size, err := bucket.DownloadToStream(req.FileID, req.Writer())
	if err != nil {
		return err
	}

	s.log.Debugf("down file: %s complete, size: %d", req.FileID, size)

	return nil
}

func (s *service) getBucket(name string) (*gridfs.Bucket, error) {
	opts := options.GridFSBucket()
	opts.SetName(name)

	bucket, err := gridfs.NewBucket(s.db, opts)
	if err != nil {
		return nil, fmt.Errorf("new bucket error, %s", err)
	}

	return bucket, nil
}
