package grpc

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/system/notify"
	"github.com/infraboard/keyauth/pkg/system/notify/mail"
	"github.com/infraboard/keyauth/pkg/system/notify/sms"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
	"github.com/infraboard/keyauth/pkg/verifycode"
)

func (s *service) IssueCode(ctx context.Context, req *verifycode.IssueCodeRequest) (
	*verifycode.IssueCodeResponse, error) {
	code, err := verifycode.NewCode(req)
	if err != nil {
		return nil, err
	}

	// 根据系统配置设置校验码过期时间
	cf, err := s.system.GetConfig()
	if err != nil {
		s.log.Errorf("get system config error, %s", err)
	} else {
		code.ExpiredMinite = uint32(cf.VerifyCode.ExpireMinutes)
	}

	// 如果是issue by pass, 这要检测
	var tk *token.Token
	switch req.IssueType {
	case verifycode.IssueType_PASS:
		tk, err = s.issuer.IssueToken(ctx, token.NewIssueTokenByPassword(
			req.ClientId,
			req.ClientSecret,
			req.Username,
			req.Password),
		)
		if err != nil {
			return nil, err
		}
	case verifycode.IssueType_TOKEN:
		tk, err = pkg.GetTokenFromGrpcInCtx(ctx)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown issue_type %s", req.IssueType)
	}
	fmt.Println(tk)
	if _, err := s.col.InsertOne(context.TODO(), code); err != nil {
		return nil, exception.NewInternalServerError("inserted verify code(%s) document error, %s",
			code, err)
	}

	inctx := pkg.NewGrpcInCtx()
	inctx.SetIsInternalCall(tk.Account, tk.Domain)

	msg, err := s.sendCode(inctx.Context(), code)
	if err != nil {
		return nil, exception.NewInternalServerError("send verify code error, %s", err)
	}

	return verifycode.NewIssueCodeResponse(msg), nil
}

func (s *service) sendCode(ctx context.Context, code *verifycode.Code) (string, error) {
	system, err := s.system.GetConfig()
	if err != nil {
		return "", fmt.Errorf("query system config error, %s", err)
	}

	u, err := s.user.DescribeAccount(ctx, user.NewDescriptAccountRequestWithAccount(code.Username))
	if err != nil {
		return "", fmt.Errorf("get user error, %s", err)
	}

	var message string
	vc := system.VerifyCode
	switch vc.NotifyType {
	case verifycode.NotifyType_MAIL:
		sender, err := mail.NewSender(system.Email)
		if err != nil {
			return "", fmt.Errorf("new sms sender error, %s", err)
		}
		req := notify.NewSendMailRequest()
		req.To = u.Profile.Email
		req.Subject = "验证码"
		req.Content = vc.RenderMailTemplate(code.Number, code.ExpiredMiniteString())
		if err := sender.Send(req); err != nil {
			return "", fmt.Errorf("send verify code by mail error, %s", err)
		}
		message = fmt.Sprintf("验证码已通过邮件发送到你的邮箱: %s, 请及时查收", u.Profile.Email)
		s.log.Debugf("send verify code to user: %s by mail ok", code.Username)
	case verifycode.NotifyType_SMS:
		sender, err := sms.NewSender(system.SMS)
		if err != nil {
			return "", fmt.Errorf("new sms sender error, %s", err)
		}
		req := notify.NewSendSMSRequest()
		req.AddPhone(u.Profile.Phone)
		req.TemplateID = vc.SmsTemplateID
		req.AddParams(code.Number, code.ExpiredMiniteString())
		if err := sender.Send(req); err != nil {
			return "", fmt.Errorf("send verify code by sms error, %s", err)
		}
		message = fmt.Sprintf("验证码已通过短信发送到你的手机: %s, 请及时查收", u.Profile.Phone)
		s.log.Debugf("send verify code to user: %s by sms ok", code.Username)
	default:
		return "", fmt.Errorf("unknown notify type %s", vc.NotifyType)
	}

	return message, nil
}

func (s *service) CheckCode(ctx context.Context, req *verifycode.CheckCodeRequest) (*verifycode.Code, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate check code request error, %s", err)
	}

	code := verifycode.NewDefaultCode()
	if err := s.col.FindOne(context.TODO(), bson.M{"_id": req.HashID()}).Decode(code); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("verify code: %s  not found", req.Number)
		}

		return nil, exception.NewInternalServerError("find system config %s error, %s", req.Number, err)
	}

	// 校验Token是否过期
	if code.IsExpired() {
		return nil, exception.NewPermissionDeny("verify code is expired")
	}

	// 没过去验证成功, 删除
	if err := s.delete(code); err != nil {
		s.log.Errorf("delete check ok verify code error, %s", err)
	}

	return code, nil
}
