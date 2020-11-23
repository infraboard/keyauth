package sms

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"

	"github.com/infraboard/keyauth/pkg/system/notify"
)

// NewTenCentSMSSender todo
func NewTenCentSMSSender(conf *TenCentSMS) (notify.SMSSender, error) {
	s := &tencent{TenCentSMS: conf}
	if err := s.init(); err != nil {
		return nil, err
	}
	return s, nil
}

type tencent struct {
	*TenCentSMS

	sms *sms.Client
}

func (s *tencent) init() error {
	credential := common.NewCredential(
		s.SecretID,
		s.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = s.GetEndpoint()
	client, err := sms.NewClient(credential, "", cpf)
	if err != nil {
		return err
	}
	s.sms = client
	return nil
}

// Send todo
func (s *tencent) Send(req *notify.SendSMSRequest) error {
	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs(req.PhoneNumberSet)
	request.TemplateParamSet = common.StringPtrs(req.ParamSet)
	request.TemplateID = common.StringPtr(req.TemplateID)
	request.SmsSdkAppid = common.StringPtr(s.AppID)
	request.Sign = common.StringPtr(s.Sign)

	response, err := s.sms.SendSms(request)
	if err != nil {
		return err
	}
	fmt.Printf("%s", response.ToJsonString())
	return nil
}
