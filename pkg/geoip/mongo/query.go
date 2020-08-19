package mongo

import (
	"net"

	"github.com/infraboard/keyauth/pkg/geoip"
	"go.mongodb.org/mongo-driver/bson"
)

// 只处理 IPv4
func newGetIPv4RequestFromIP(ip net.IP) *getIPv4Request {
	return &getIPv4Request{
		ip:     ip.To4(),
		filter: bson.A{},
	}
}

type getIPv4Request struct {
	ip     net.IP
	filter bson.A
}

func (req *getIPv4Request) addCond(cond bson.M) {
	req.filter = append(req.filter, cond)
}

func (req *getIPv4Request) Filter() bson.M {
	if req.ip != nil {
		ipInt, _ := geoip.IPToInt(req.ip)
		req.addCond(bson.M{"start": bson.M{"$lte": ipInt.Uint64()}})
		req.addCond(bson.M{"end": bson.M{"$gte": ipInt.Uint64()}})
	}

	if len(req.filter) == 0 {
		return bson.M{}
	}

	return bson.M{"$and": req.filter}
}

// 只处理 IPv4
func newGetLocationRequestFromID(geonameID string) *getLocationRequest {
	return &getLocationRequest{
		GeonameID: geonameID,
		filter:    bson.M{},
	}
}

type getLocationRequest struct {
	GeonameID string
	filter    bson.M
}

func (req *getLocationRequest) Filter() bson.M {
	if req.GeonameID != "" {
		req.filter["_id"] = req.GeonameID
	}

	return req.filter
}
