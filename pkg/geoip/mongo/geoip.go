package mongo

import (
	"io"
	"net"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/geoip"
)

func (s *service) UpdateDBFile(req *geoip.UpdateDBFileRequest) error {
	opts := options.GridFSUpload()

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

func (s *service) LookupIP(ipAddress net.IP) (*geoip.Record, error) {
	return nil, nil
}
