package mongo

import (
	"bufio"
	"io"
	"net"

	"github.com/infraboard/keyauth/pkg/geoip"
	"github.com/infraboard/mcube/exception"
)

func (s *service) UpdateDBFile(req *geoip.UpdateDBFileRequest) error {
	if err := req.Validate(); err != nil {
		return exception.NewBadRequest("validate update db file requrest error, %s", err)
	}

	stream := req.ReadCloser()
	defer stream.Close()

	var lineCount uint
	reader := bufio.NewReader(stream)

	var (
		ls *geoip.LocationSet
		is *geoip.IPv4Set
	)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return exception.NewInternalServerError("read line data error, %s", err)
		}

		if err == io.EOF {
			if err := s.saveLocationSet(ls); err != nil {
				return exception.NewInternalServerError("save location set error, %s", err)
			}
			if err := s.saveIPv4Set(is); err != nil {
				return exception.NewInternalServerError("save ipv4 set error, %s", err)
			}
			s.log.Infof("read line complete, total line: %d", lineCount)
			break
		}
		lineCount++

		// 由于第一行是表头, 需要跳过
		if lineCount == 1 {
			continue
		}

		switch req.ContentType {
		case geoip.IPv4Content:
			ipv4, err := geoip.ParseIPv4FromCsvLine(line)
			if err != nil {
				return exception.NewBadRequest("parse line %d csv data error, %s", lineCount, err)
			}
			if is == nil {
				is = geoip.NewIPv4Set(1024)
			}
			is.Add(ipv4)
			if is.IsFull() {
				if err := s.saveIPv4Set(is); err != nil {
					return exception.NewInternalServerError("bulk save error, %s", err)
				}
				is.Reset()
				s.log.Infof("lineCount: %d, ", lineCount)
			}
		case geoip.LocationContent:
			location, err := geoip.ParseLocationFromCsvLine(line)
			if err != nil {
				return exception.NewBadRequest("parse line %d csv data error, %s", lineCount, err)
			}
			if ls == nil {
				ls = geoip.NewLocationSet(256)
			}
			ls.Add(location)
			if ls.IsFull() {
				if err := s.saveLocationSet(ls); err != nil {
					return exception.NewInternalServerError("bulk save error, %s", err)
				}
				ls.Reset()
				s.log.Infof("lineCount: %d, ", lineCount)
			}
		default:
			return exception.NewBadRequest("unknown content type, %s", req.ContentType)
		}
	}

	return nil
}

func (s *service) LookupIP(ip net.IP) (*geoip.Record, error) {
	if ip.To4() == nil {
		return nil, exception.NewBadRequest("%v is not an IPv4 address", ip)
	}

	req := newGetIPv4RequestFromIP(ip)
	ipv4, err := s.findIPv4One(req)
	if err != nil {
		return nil, err
	}

	locationReq := newGetLocationRequestFromID(ipv4.GeonameID)
	location, err := s.findLocationOne(locationReq)
	if err != nil {
		s.log.Errorf("find geoip location error, %s", err)
	} else {

	}

	return geoip.NewRecord(ipv4, location), nil
}
