package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	common "github.com/infraboard/keyauth/common/types"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func (s *service) QueryAccount(t types.Type, req *user.QueryAccountRequest) (*user.Set, error) {
	r, err := newQueryUserRequest(req)
	if err != nil {
		return nil, err
	}

	r.userType = t
	return s.queryAccount(r)
}

func (s *service) CreateAccount(t types.Type, req *user.CreateAccountRequest) (*user.User, error) {
	u, err := user.New(req)
	if err != nil {
		return nil, err
	}

	tk := req.GetToken()
	if tk != nil {
		u.Domain = tk.Domain
	}

	u.Type = t
	if err := s.saveAccount(u); err != nil {
		return nil, err
	}

	u.HashedPassword = nil
	return u, nil
}

func (s *service) UpdateAccountProfile(req *user.UpdateAccountRequest) (*user.User, error) {
	u, err := s.DescribeAccount(user.NewDescriptAccountRequestWithAccount(req.Account))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case common.PutUpdateMode:
		*u.Profile = *req.Profile
	case common.PatchUpdateMode:
		u.Profile.Patch(req.Profile)
	default:
		return nil, exception.NewBadRequest("unknown update mode: %s", req.UpdateMode)
	}

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate update department error, %s", err)
	}

	u.UpdateAt = ftime.Now()

	_, err = s.col.UpdateOne(context.TODO(), bson.M{"_id": u.Account}, bson.M{"$set": u})
	if err != nil {
		return nil, exception.NewInternalServerError("update user(%s) error, %s", u.Account, err)
	}

	return u, nil
}

func (s *service) UpdateAccountPassword(req *user.UpdatePasswordRequest) (*user.Password, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("check update pass request error, %s", err)
	}
	return s.changePass(req.Account, req.OldPass, req.NewPass, req.IsReset())
}

func (s *service) changePass(account, old, new string, isReset bool) (*user.Password, error) {
	descReq := user.NewDescriptAccountRequest()
	descReq.Account = account
	u, err := s.DescribeAccount(descReq)
	if err != nil {
		return nil, err
	}

	// 根据域设置的规则检测密码策略
	descDom := domain.NewDescriptDomainRequestWithName(u.Domain)
	dom, err := s.domain.DescriptionDomain(descDom)
	if err != nil {
		return nil, err
	}

	// 检测密码强度
	if err := dom.SecuritySetting.PasswordSecurity.Check(new); err != nil {
		return nil, err
	}

	// 判断是不是历史密码
	if u.HashedPassword.IsHistory(new) {
		return nil, exception.NewBadRequest("password must not last %d", dom.SecuritySetting.PasswordSecurity.RepeateLimite)
	}

	if err := u.ChangePassword(old, new, dom.SecuritySetting.GetPasswordRepeateLimite(), isReset); err != nil {
		return nil, exception.NewBadRequest("change password error, %s", err)
	}

	_, err = s.col.UpdateOne(context.TODO(), bson.M{"_id": u.Account}, bson.M{"$set": bson.M{
		"password": u.HashedPassword,
	}})

	if err != nil {
		return nil, exception.NewInternalServerError("update user(%s) password error, %s", u.Account, err)
	}

	return u.HashedPassword, nil
}

func (s *service) DescribeAccount(req *user.DescriptAccountRequest) (*user.User, error) {
	r, err := newDescribeRequest(req)
	if err != nil {
		return nil, err
	}

	ins := user.NewDefaultUser()
	if err := s.col.FindOne(context.TODO(), r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("user %s not found", req)
		}

		return nil, exception.NewInternalServerError("find user %s error, %s", req, err)
	}

	dom, err := s.domain.DescriptionDomain(domain.NewDescriptDomainRequestWithName(ins.Domain))
	if err != nil {
		return nil, err
	}

	dom.SecuritySetting.PasswordSecurity.SetPasswordNeedReset(ins.HashedPassword)
	return ins, nil
}

func (s *service) BlockAccount(account, reason string) error {
	desc := user.NewDescriptAccountRequestWithAccount(account)
	user, err := s.DescribeAccount(desc)
	if err != nil {
		return fmt.Errorf("describe user error, %s", err)
	}

	user.Block(reason)
	return s.saveAccount(user)
}

func (s *service) DeleteAccount(account string) error {
	_, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": account})
	if err != nil {
		return exception.NewInternalServerError("delete user(%s) error, %s", account, err)
	}

	// 清除账号的关联的所有策略
	if err := s.policy.DeletePolicy(policy.NewDeletePolicyRequestWithAccount(account)); err != nil {
		s.log.Errorf("delete account policy error, %s", err)
	}

	return nil
}
