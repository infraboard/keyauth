package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
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
	tk, err := s.createServiceToken(req.GetUserAgent(), req.GetRemoteIP(), user, pass)
	if err != nil {
		return nil, exception.NewInternalServerError("create service token error, %s", err)
	}
	ins.AccessToken = tk.AccessToken
	ins.RefreshToken = tk.RefreshToken
	ins.Creater = tk.Account
	ins.Domain = tk.Domain

	// 为服务用户添加策略
	_, err = s.createPolicy(tk, ins.Account, req.RoleID)
	if err != nil {
		s.log.Errorf("create service: %s policy error, %s", ins.Name, err)
	}

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

func (s *service) createServiceToken(userAgent, remoteIP, user, pass string) (*token.Token, error) {
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
	req.WithRemoteIP(remoteIP)
	req.WithUserAgent(userAgent)
	return s.token.IssueToken(req)
}

func (s *service) createPolicy(tk *token.Token, account, roleID string) (*policy.Policy, error) {
	if roleID == "" {
		descR := role.NewDescribeRoleRequestWithName(role.VisitorRoleName)
		descR.WithToken(tk)
		adminR, err := s.role.DescribeRole(descR)
		if err != nil {
			return nil, err
		}
		roleID = adminR.ID
	}

	req := policy.NewCreatePolicyRequest()
	req.WithToken(tk)
	req.Account = account
	req.NamespaceID = "*"
	req.RoleID = roleID
	return s.policy.CreatePolicy(policy.BuildInPolicy, req)
}

func (s *service) refreshServiceToken(at, rt string) (*token.Token, error) {
	app, err := s.app.GetBuildInApplication(application.AdminServiceApplicationName)
	if err != nil {
		return nil, err
	}
	req := token.NewIssueTokenRequest()
	req.GrantType = token.REFRESH
	req.AccessToken = at
	req.RefreshToken = rt
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

func (s *service) RefreshServicToken(req *micro.DescribeMicroRequest) (
	*token.Token, error) {
	ins, err := s.DescribeService(req)
	if err != nil {
		return nil, err
	}

	tk, err := s.refreshServiceToken(ins.AccessToken, ins.RefreshToken)
	if err != nil {
		return nil, err
	}

	ins.AccessToken = tk.AccessToken
	ins.RefreshToken = tk.RefreshToken
	tk.Desensitize()

	if err := s.update(ins); err != nil {
		return nil, err
	}

	return tk, nil
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
