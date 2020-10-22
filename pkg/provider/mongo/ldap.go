package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/provider"
	"github.com/infraboard/keyauth/pkg/provider/ldap"
)

func (s *service) SaveConfig(req *provider.SaveLDAPConfigRequest) (
	*provider.LDAPConfig, error) {
	ins, err := provider.NewLDAPConfig(req)
	if err != nil {
		return nil, err
	}

	// 创建或者更新
	descLDAP := provider.NewDescribeLDAPConfigWithDomain(ins.Domain)
	_, err = s.DescribeConfig(descLDAP)
	if exception.IsNotFoundError(err) {
		err = s.save(ins)
	} else {
		err = s.update(ins)
	}
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) QueryConfig(req *provider.QueryLDAPConfigRequest) (*provider.LDAPSet, error) {
	r := newQueryLDAPRequest(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find ldap error, error is %s", err)
	}

	set := provider.NewLDAPSet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		ins := provider.NewDefaultLDAPConfig()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode ldap error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get ldap count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (s *service) DescribeConfig(req *provider.DescribeLDAPConfig) (*provider.LDAPConfig, error) {
	r, err := newDescribeLDAPRequest(req)
	if err != nil {
		return nil, err
	}

	ins := provider.NewDefaultLDAPConfig()
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("ldap %s not found", req)
		}

		return nil, exception.NewInternalServerError("find ldap %s error, %s", req.Domain, err)
	}

	return ins, nil
}

func (s *service) DeleteConfig(req *provider.DeleteLDAPConfig) error {
	return nil
}

func (s *service) CheckConnect(req *provider.DescribeLDAPConfig) error {
	cfg, err := s.DescribeConfig(req)
	if err != nil {
		return err
	}

	p := ldap.NewProvider(cfg.Config)
	return p.CheckConnect()
}
