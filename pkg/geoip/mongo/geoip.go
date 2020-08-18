package mongo

import (
	"io"
	"net"

	"github.com/infraboard/keyauth/pkg/geoip"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) UploadDBFile(req *geoip.UploadFileRequest) error {
	opts := options.GridFSUpload()
	opts.Metadata = req.Meta()

	// 清除已有文件
	s.bucket.Delete(s.dbFileName)

	// 上传新文件
	uploadStream, err := s.bucket.OpenUploadStreamWithID(s.dbFileName, s.dbFileName, opts)
	if err != nil {
		return err
	}
	defer uploadStream.Close()

	fileSize, err := io.Copy(uploadStream, req.ReadCloser())
	if err != nil {
		return err
	}

	s.log.Debugf("Write file %s to DB was successful. File size: %d M", s.dbFileName, fileSize/1024/1024)
	return nil
}

func (s *service) Lookup(ipAddress net.IP) (*geoip.Record, error) {
	return nil, nil
}
