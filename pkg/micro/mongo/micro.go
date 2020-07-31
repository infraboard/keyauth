package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func (s *service) CreateService(req *micro.CreateMicroRequest) (
	*micro.Micro, error) {
	ins, err := micro.New(req)
	if err != nil {
		return nil, err
	}

	user, pass := ins.Name, xid.New().String()
	// 创建服务用户
	account, err := s.createServiceAccount(req.GetToken(), user, pass)
	if err != nil {
		return nil, exception.NewInternalServerError("create service account error, %s", err)
	}
	ins.Account = account.Account

	// 使用用户创建服务访问Token
	tk, err := s.createServiceToken(user, pass)
	if err != nil {
		return nil, exception.NewInternalServerError("create service token error, %s", err)
	}
	ins.AccessToken = tk.AccessToken
	ins.RefreshToken = tk.RefreshToken
	ins.Creater = tk.Account

	if _, err := s.scol.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a service document error, %s", err)
	}
	return ins, nil
}

func (s *service) createServiceAccount(tk *token.Token, name, pass string) (*user.User, error) {
	req := user.NewCreateUserRequest()
	req.WithToken(tk)
	req.Account = name
	req.Password = pass
	return s.user.CreateAccount(types.ServiceAccount, req)
}

func (s *service) createServiceToken(user, pass string) (*token.Token, error) {
	app, err := s.app.GetBuildInApplication(application.AdminServiceApplicationName)
	if err != nil {
		return nil, err
	}
	req := token.NewIssueTokenRequest()
	req.GrantType = token.PASSWORD
	req.Username = user
	req.Password = pass
	req.ClientID = app.ClientID
	req.ClientSecret = app.ClientSecret
	return s.token.IssueToken(req)
}

func (s *service) QueryService(req *micro.QueryMicroRequest) (*micro.Set, error) {
	r := newPaggingQuery(req)
	resp, err := s.scol.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find service error, error is %s", err)
	}

	set := micro.NewMicroSet(req.PageRequest)
	// 循环
	for resp.Next(context.TODO()) {
		ins := new(micro.Micro)
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode service error, error is %s", err)
		}

		set.Add(ins)
	}

	// count
	count, err := s.scol.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) DescribeService(req *micro.DescribeMicroRequest) (
	*micro.Micro, error) {
	r, err := newDescribeQuery(req)
	if err != nil {
		return nil, err
	}

	ins := new(micro.Micro)
	if err := s.scol.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("service %s not found", req)
		}

		return nil, exception.NewInternalServerError("find service %s error, %s", req, err)
	}
	return ins, nil
}

func (s *service) DeleteService(id string) error {
	describeReq := micro.NewDescriptServiceRequest()
	describeReq.ID = id
	if _, err := s.DescribeService(describeReq); err != nil {
		return err
	}

	_, err := s.scol.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return exception.NewInternalServerError("delete service(%s) error, %s", id, err)
	}
	return nil
}
