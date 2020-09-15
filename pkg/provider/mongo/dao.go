package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/infraboard/keyauth/pkg/provider"
)

func (s *service) update(ins *provider.LDAPConfig) error {
	ins.UpdateAt = ftime.Now()
	_, err := s.col.UpdateOne(context.TODO(), bson.M{"_id": ins.Domain}, bson.M{"$set": ins})
	if err != nil {
		return exception.NewInternalServerError("update domain(%s) error, %s", ins.Domain, err)
	}

	return nil
}

func (s *service) save(ins *provider.LDAPConfig) error {
	if _, err := s.col.InsertOne(context.TODO(), ins); err != nil {
		return exception.NewInternalServerError("inserted ldap(%s) document error, %s",
			ins.BaseDN, err)
	}
	return nil
}
