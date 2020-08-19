package mongo

import "github.com/infraboard/keyauth/pkg/ip2region"

func (s *service) UpdateDBFile(*ip2region.UpdateDBFileRequest) error {
	return nil
}

func (s *service) LookupIP(ip string) (*ip2region.IPInfo, error) {
	return nil, nil
}
