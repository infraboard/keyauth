package impl

import (
	"context"

	"github.com/infraboard/keyauth/apps/wxwork"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"

	authWXWork "github.com/infraboard/keyauth/apps/provider/auth/wxwork"
)

func (s *service) SaveConfig(req *wxwork.SaveConfRequest) (
	*wxwork.WechatWorkConfig, error) {
	ins, err := wxwork.NewWechatWorkConfig(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate error, %s", err)
	}

	p := authWXWork.NewProvider(ins.Config)
	if err := p.Check(); err != nil {
		return nil, exception.NewBadRequest("try Authentication error, %s", err)
	}

	descWxwork := wxwork.NewDescribeConfWithDomain(ins.Domain)
	old, err := s.DescribeConfig(descWxwork)
	if err != nil && !exception.IsNotFoundError(err) {
		return nil, err
	}

	// 保存入库
	if old == nil {
		err = s.save(ins)
	} else {
		err = s.update(ins)
	}
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) QueryConfig(req *wxwork.QueryConfigRequest) (*wxwork.WechatWorkSet, error) {
	r := newQueryConfigRequest(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find wechat work config error, error is %s", err)
	}

	set := wxwork.NewConfigSet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		ins := wxwork.NewDefaultWechatWorkConfig()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode wechat work config error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get wechat work config count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (s *service) DescribeConfig(req *wxwork.DescribeWechatWorkConf) (*wxwork.WechatWorkConfig, error) {
	r, err := newDescribeWechatWorkRequest(req)
	if err != nil {
		return nil, err
	}

	ins := wxwork.NewDefaultWechatWorkConfig()
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("wxwork %s not found", req)
		}

		return nil, exception.NewInternalServerError("find wxwork %s error, %s", req.Domain, err)
	}

	return ins, nil
}

func (s *service) DeleteConfig(req *wxwork.DescribeWechatWorkConf) error {
	r, err := newDescribeWechatWorkRequest(req)
	if err != nil {
		return err
	}

	_, err = s.col.DeleteOne(context.TODO(), r.FindFilter())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return exception.NewNotFound("wxwork %s not found", req)
		}

		return exception.NewInternalServerError("find wxwork %s error, %s", req.Domain, err)
	}

	return nil
}
