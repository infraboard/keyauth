package mongo

import (
	"context"

	"github.com/infraboard/keyauth/pkg/geoip"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) saveLocationSet(ls *geoip.LocationSet) error {
	if ls == nil {
		return nil
	}

	if ls.Length() == 0 {
		return exception.NewBadRequest("no location need to save")
	}

	docs := make([]interface{}, 0, ls.Length())
	items := ls.Items()
	for i := range items {
		docs = append(docs, items[i])
	}

	rest, err := s.location.InsertMany(context.TODO(), docs)
	if err != nil {
		return err
	}

	s.log.Infof("bulk save locations %d", len(rest.InsertedIDs))
	return nil
}

func (s *service) saveIPv4Set(is *geoip.IPv4Set) error {
	if is == nil {
		return nil
	}

	if is.Length() == 0 {
		return exception.NewBadRequest("no location need to save")
	}

	docs := make([]interface{}, 0, is.Length())
	items := is.Items()
	for i := range items {
		docs = append(docs, items[i])
	}

	rest, err := s.ip.InsertMany(context.TODO(), docs)
	if err != nil {
		return err
	}

	s.log.Infof("bulk save locations %d", len(rest.InsertedIDs))
	return nil
}

func (s *service) findIPv4One(req *getIPv4Request) (*geoip.IPv4, error) {
	ins := geoip.NewDefaultIPv4()
	if err := s.ip.FindOne(context.TODO(), req.Filter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("geoip ipv4 %s not found", req)
		}

		return nil, exception.NewInternalServerError("find geoip ipv4 %s error, %s", req.ip, err)
	}

	return ins, nil
}

func (s *service) findLocationOne(req *getLocationRequest) (*geoip.Location, error) {

	ins := geoip.NewDefaultLocation()
	if err := s.location.FindOne(context.TODO(), req.Filter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("geoip location %s not found", req)
		}

		return nil, exception.NewInternalServerError("find geoip location %s error, %s", req.GeonameID, err)
	}

	return ins, nil
}
