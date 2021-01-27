package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/common/session"
	"github.com/infraboard/keyauth/pkg/application"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/micro"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func (s *service) CreateService(ctx context.Context, req *micro.CreateMicroRequest) (
	*micro.Micro, error) {
	ins, err := micro.New(req)
	if err != nil {
		return nil, err
	}

	tk := session.GetTokenFromContext(ctx)
	if tk == nil {
		return nil, exception.NewPermissionDeny("token required")
	}

	user, pass := ins.Data.Name, xid.New().String()
	// 创建服务用户
	account, err := s.createServiceAccount(ctx, user, pass)
	if err != nil {
		return nil, exception.NewInternalServerError("create service account error, %s", err)
	}
	ins.Account = account.Data.Profile.Account

	// 使用用户创建服务访问Token
	svrTK, err := s.createServiceToken(tk.GetUserAgent(), tk.GetRemoteIP(), user, pass)
	if err != nil {
		return nil, exception.NewInternalServerError("create service token error, %s", err)
	}
	ins.AccessToken = svrTK.AccessToken
	ins.RefreshToken = svrTK.RefreshToken
	ins.Creater = svrTK.Account
	ins.Domain = svrTK.Domain

	// 为服务用户添加策略
	_, err = s.createPolicy(ctx, ins.Account, req.RoleId)
	if err != nil {
		s.log.Errorf("create service: %s policy error, %s", ins.Data.Name, err)
	}

	if _, err := s.scol.InsertOne(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a service document error, %s", err)
	}
	return ins, nil
}

func (s *service) createServiceAccount(ctx context.Context, name, pass string) (*user.User, error) {
	req := user.NewCreateUserRequest()
	req.Profile.Account = name
	req.Password = pass
	req.UserType = types.UserType_SERVICE
	return s.user.CreateAccount(ctx, req)
}

func (s *service) createServiceToken(userAgent, remoteIP, user, pass string) (*token.Token, error) {
	app, err := s.app.GetBuildInApplication(nil, application.NewGetBuildInAdminApplicationRequest())
	if err != nil {
		return nil, err
	}
	req := token.NewIssueTokenRequest()
	req.GrantType = token.GrantType_PASSWORD
	req.Username = user
	req.Password = pass
	req.ClientId = app.ClientId
	req.ClientSecret = app.ClientSecret
	req.WithRemoteIP(remoteIP)
	req.WithUserAgent(userAgent)
	return s.token.IssueToken(nil, req)
}

func (s *service) revolkServiceToken(accessToken string) error {
	app, err := s.app.GetBuildInApplication(nil, application.NewGetBuildInAdminApplicationRequest())
	if err != nil {
		return err
	}
	req := token.NewRevolkTokenRequest(app.ClientId, app.ClientSecret)
	req.AccessToken = accessToken
	_, err = s.token.RevolkToken(nil, req)
	return err
}

func (s *service) createPolicy(ctx context.Context, account, roleID string) (*policy.Policy, error) {
	if roleID == "" {
		descR := role.NewDescribeRoleRequestWithName(role.VisitorRoleName)
		adminR, err := s.role.DescribeRole(ctx, descR)
		if err != nil {
			return nil, err
		}
		roleID = adminR.Id
	}

	req := policy.NewCreatePolicyRequest()
	req.Account = account
	req.NamespaceId = "*"
	req.RoleId = roleID
	req.Type = policy.PolicyType_BUILD_IN
	return s.policy.CreatePolicy(ctx, req)
}

func (s *service) refreshServiceToken(at, rt string) (*token.Token, error) {
	app, err := s.app.GetBuildInApplication(nil, application.NewGetBuildInAdminApplicationRequest())
	if err != nil {
		return nil, err
	}
	req := token.NewIssueTokenRequest()
	req.GrantType = token.GrantType_REFRESH
	req.AccessToken = at
	req.RefreshToken = rt
	req.ClientId = app.ClientId
	req.ClientSecret = app.ClientSecret
	return s.token.IssueToken(nil, req)
}

func (s *service) QueryService(ctx context.Context, req *micro.QueryMicroRequest) (*micro.Set, error) {
	r := newPaggingQuery(req)
	resp, err := s.scol.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find service error, error is %s", err)
	}

	set := micro.NewMicroSet()
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

func (s *service) DescribeService(ctx context.Context, req *micro.DescribeMicroRequest) (
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

func (s *service) RefreshServiceToken(ctx context.Context, req *micro.DescribeMicroRequest) (
	*token.Token, error) {
	ins, err := s.DescribeService(ctx, req)
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

func (s *service) DeleteService(ctx context.Context, req *micro.DeleteMicroRequest) (*micro.Micro, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate delete service error, %s", err)
	}

	describeReq := micro.NewDescribeServiceRequest()
	describeReq.Id = req.Id
	svr, err := s.DescribeService(ctx, describeReq)
	if err != nil {
		return nil, err
	}

	if micro.Type.IsIn(micro.Type_BUILD_IN) {
		return nil, exception.NewBadRequest("service %s is system service, your can't delete", svr.Data.Name)
	}

	// 清除服务实体
	_, err = s.scol.DeleteOne(context.TODO(), bson.M{"_id": req.Id})
	if err != nil {
		return nil, exception.NewInternalServerError("delete service(%s) error, %s", req.Id, err)
	}

	// 删除服务默认策略
	dpReq := policy.NewDeletePolicyRequestWithAccount(svr.Account)
	_, err = s.policy.DeletePolicy(ctx, dpReq)
	if err != nil {
		s.log.Errorf("delete service policy error, %s", err)
	}

	// 删除服务注册的Endpoint
	deReq := endpoint.NewDeleteEndpointRequestWithServiceID(svr.Id)
	err = s.endpoint.DeleteEndpoint(deReq)
	if err != nil {
		s.log.Errorf("delete service endpoint error, %s", err)
	}

	// 删除服务的Token
	err = s.revolkServiceToken(svr.AccessToken)
	if err != nil {
		s.log.Errorf("revolk service token error, %s", err)
	}

	return svr, nil
}
