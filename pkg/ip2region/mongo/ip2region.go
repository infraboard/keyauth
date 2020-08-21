package mongo

import (
	"bytes"
	"fmt"
	"os"

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

	// 优先从本地文件加载DB文件
	if err := s.loadDBFileFromLocal(); err != nil {
		s.log.Infof("load ip2region db file from local error, %s, retry other load method ", err)
	} else {
		return s.dbReader, nil
	}

	if err := s.loadDBFileFromBucket(); err != nil {
		s.log.Info("load ip2region db file from bucket error, %s")
	} else {
		return s.dbReader, nil
	}

	return nil, fmt.Errorf("load ip2region db file error")
}

func (s *service) loadDBFileFromLocal() error {
	file, err := os.Open(s.dbFileName)
	if err != nil {
		return fmt.Errorf("open file error, %s", err)
	}

	reader, err := reader.New(file)
	if err != nil {
		return err
	}
	s.dbReader = reader
	return nil
}

func (s *service) loadDBFileFromBucket() error {
	buf := bytes.NewBuffer([]byte{})
	downloadReq := storage.NewDownloadFileRequest(s.bucketName, s.dbFileName, buf)
	if err := s.storage.Download(downloadReq); err != nil {
		return err
	}

	reader, err := reader.New(buf)
	if err != nil {
		return err
	}
	s.dbReader = reader

	return nil
}
