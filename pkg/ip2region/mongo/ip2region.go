package mongo

import (
	"bytes"

	"github.com/infraboard/mcube/exception"

	"github.com/infraboard/keyauth/pkg/ip2region"
	"github.com/infraboard/keyauth/pkg/ip2region/reader"
	"github.com/infraboard/keyauth/pkg/storage"
)

func (s *service) UpdateDBFile(req *ip2region.UpdateDBFileRequest) error {
	if err := req.Validate(); err != nil {
		return exception.NewBadRequest("validate update db file requrest error, %s", err)
	}

	reader := req.ReadCloser()
	defer reader.Close()

	uploadReq := storage.NewUploadFileRequest(s.bucketName, s.dbFileName, reader)
	uploadReq.WithToken(req.GetToken())
	return s.storage.UploadFile(uploadReq)
}

func (s *service) LookupIP(ip string) (*ip2region.IPInfo, error) {
	dbReader, err := s.getDBReader()
	if err != nil {
		return nil, err
	}
	return dbReader.MemorySearch(ip)
}

func (s *service) getDBReader() (*reader.IPReader, error) {
	s.Lock()
	defer s.Unlock()

	if s.dbReader != nil {
		return s.dbReader, nil
	}

	buf := bytes.NewBuffer([]byte{})
	downloadReq := storage.NewDownloadFileRequest(s.bucketName, s.dbFileName, buf)
	if err := s.storage.Download(downloadReq); err != nil {
		return nil, err
	}
	reader, err := reader.New(buf)
	if err != nil {
		return nil, err
	}
	s.dbReader = reader

	return s.dbReader, nil
}
