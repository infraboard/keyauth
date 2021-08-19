package grpc

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/common/password"
	common "github.com/infraboard/keyauth/common/types"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func (s *service) QueryAccount(ctx context.Context, req *user.QueryAccountRequest) (*user.Set, error) {
	r, err := newQueryUserRequest(req)
	if err != nil {
		return nil, err
	}
	return s.queryAccount(ctx, r)
}

func (s *service) CreateAccount(ctx context.Context, req *user.CreateAccountRequest) (*user.User, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 非管理员, 主账号 可以创建子账号
	if !tk.UserType.IsIn(types.UserType_SUPPER, types.UserType_INTERNAL, types.UserType_PRIMARY) {
		return nil, exception.NewPermissionDeny("%s user can't create sub account", tk.UserType)
	}

	u, err := user.New(req)
	u.CreateType = user.CreateType_DOMAIN_CREATED
	if err != nil {
		return nil, err
	}

	if tk != nil {
		u.Domain = tk.Domain
	}

	// 如果是管理员创建的账号需要用户自己重置密码
	if u.CreateType.IsIn(user.CreateType_DOMAIN_CREATED) {
		u.HashedPassword.SetNeedReset("admin created user need reset when first login")
	}

	if err := s.saveAccount(u); err != nil {
		return nil, err
	}

	u.HashedPassword = nil
	return u, nil
}

func (s *service) UpdateAccountProfile(ctx context.Context, req *user.UpdateAccountRequest) (*user.User, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate update department error, %s", err)
	}

	s.log.Debugf("[%s] update %s profile", req.UpdateMode.String(), req.Account)
	u, err := s.DescribeAccount(ctx, user.NewDescriptAccountRequestWithAccount(req.Account))
	if err != nil {
		return nil, err
	}
	u.UpdateAt = ftime.Now().Timestamp()

	// 更新部门
	if req.DepartmentId != "" {
		if !tk.UserType.IsIn(types.UserType_SUPPER, types.UserType_INTERNAL, types.UserType_DOMAIN_ADMIN, types.UserType_ORG_ADMIN) {
			return nil, exception.NewBadRequest("组织管理员才能直接修改用户部门")
		}
		u.DepartmentId = req.DepartmentId
	}

	// 更新profile
	if req.Profile != nil {
		u.IsInitialized = true
		switch req.UpdateMode {
		case common.UpdateMode_PUT:
			*u.Profile = *req.Profile
		case common.UpdateMode_PATCH:
			u.Profile.Patch(req.Profile)
		default:
			return nil, exception.NewBadRequest("unknown update mode: %s", req.UpdateMode)
		}
	}

	_, err = s.col.UpdateOne(context.TODO(), bson.M{"_id": u.Account}, bson.M{"$set": u})
	if err != nil {
		return nil, exception.NewInternalServerError("update user(%s) error, %s", u.Account, err)
	}

	return u, nil
}

func (s *service) UpdateAccountPassword(ctx context.Context, req *user.UpdatePasswordRequest) (*user.Password, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("check update pass request error, %s", err)
	}
	return s.changePass(ctx, req.Account, req.OldPass, req.NewPass)
}

func (s *service) changePass(ctx context.Context, account, old, new string) (*user.Password, error) {
	descReq := user.NewDescriptAccountRequest()
	descReq.Account = account
	s.log.Debugf("query user account ...")
	u, err := s.DescribeAccount(ctx, descReq)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query domain security setting ...")
	// 根据域设置的规则检测密码策略
	descDom := domain.NewDescribeDomainRequestWithName(u.Domain)
	dom, err := s.domain.DescribeDomain(ctx, descDom)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("check password  strength ...")
	// 检测密码强度
	if err := dom.SecuritySetting.PasswordSecurity.Check(new); err != nil {
		return nil, err
	}

	s.log.Debugf("check password  is history ...")
	// 判断是不是历史密码
	if u.HashedPassword.IsHistory(new) {
		return nil, exception.NewBadRequest("password not last %d used", dom.SecuritySetting.PasswordSecurity.RepeateLimite)
	}

	// 非本人重置密码, 需要用户下次登录时重置
	var isReset bool
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	if !tk.IsOwner(account) {
		isReset = true
	}

	s.log.Debugf("change password ...")
	if err := u.ChangePassword(old, new, dom.SecuritySetting.GetPasswordRepeateLimite(), isReset); err != nil {
		return nil, exception.NewBadRequest("change password error, %s", err)
	}

	s.log.Debugf("save password to db ...")
	_, err = s.col.UpdateOne(context.TODO(), bson.M{"_id": u.Account}, bson.M{"$set": bson.M{
		"password": u.HashedPassword,
	}})

	if err != nil {
		return nil, exception.NewInternalServerError("update user(%s) password error, %s", u.Account, err)
	}

	u.Desensitize()
	return u.HashedPassword, nil
}

func (s *service) DescribeAccount(ctx context.Context, req *user.DescribeAccountRequest) (*user.User, error) {
	r, err := newDescribeRequest(req)
	if err != nil {
		return nil, err
	}

	ins := user.NewDefaultUser()
	if err := s.col.FindOne(ctx, r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("user %s not found", req)
		}

		return nil, exception.NewInternalServerError("find user %s error, %s", req, err)
	}

	dom, err := s.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequestWithName(ins.Domain))
	if err != nil {
		return nil, err
	}

	dom.SecuritySetting.PasswordSecurity.SetPasswordNeedReset(ins.HashedPassword)
	return ins, nil
}

func (s *service) BlockAccount(ctx context.Context, req *user.BlockAccountRequest) (*user.User, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	desc := user.NewDescriptAccountRequestWithAccount(req.Account)
	user, err := s.DescribeAccount(ctx, desc)
	if err != nil {
		return nil, fmt.Errorf("describe user error, %s", err)
	}

	user.Block(req.Reason)
	_, err = s.col.UpdateOne(context.TODO(), bson.M{"_id": user.Account}, bson.M{"$set": bson.M{
		"status": user.Status,
	}})
	if err != nil {
		return nil, fmt.Errorf("update user status error, %s", err)
	}

	return user, nil
}

func (s *service) UnBlockAccount(ctx context.Context, req *user.UnBlockAccountRequest) (*user.User, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	desc := user.NewDescriptAccountRequestWithAccount(req.Account)
	user, err := s.DescribeAccount(ctx, desc)
	if err != nil {
		return nil, fmt.Errorf("describe user error, %s", err)
	}

	err = user.UnBlock()
	if err != nil {
		return nil, err
	}

	_, err = s.col.UpdateOne(context.TODO(), bson.M{"_id": user.Account}, bson.M{"$set": bson.M{
		"status": user.Status,
	}})
	if err != nil {
		return nil, fmt.Errorf("update user status error, %s", err)
	}

	return user, nil
}

func (s *service) DeleteAccount(ctx context.Context, req *user.DeleteAccountRequest) (*user.User, error) {
	_, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": req.Account})
	if err != nil {
		return nil, exception.NewInternalServerError("delete user(%s) error, %s", req.Account, err)
	}

	// 清除账号的关联的所有策略
	if _, err := s.policy.DeletePolicy(ctx, policy.NewDeletePolicyRequestWithAccount(req.Account)); err != nil {
		s.log.Errorf("delete account policy error, %s", err)
	}

	return user.NewDefaultUser(), nil
}

func (s *service) GeneratePassword(ctx context.Context, req *user.GeneratePasswordRequest) (*user.GeneratePasswordResponse, error) {
	tk, err := pkg.GetTokenFromGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query domain security setting ...")
	// 根据域设置的规则检测密码策略
	descDom := domain.NewDescribeDomainRequestWithName(tk.Domain)
	dom, err := s.domain.DescribeDomain(ctx, descDom)
	if err != nil {
		return nil, err
	}

	genConf := dom.SecuritySetting.PasswordSecurity.GenRandomPasswordConfig()
	ranPass, err := password.New(&genConf).Generate()
	if err != nil {
		return nil, fmt.Errorf("generate random password error, %s", err)
	}
	return user.NewGeneratePasswordResponse(*ranPass), nil
}
