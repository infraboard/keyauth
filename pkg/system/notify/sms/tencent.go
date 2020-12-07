package sms

import (
	"fmt"
	"strings"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"

	"github.com/infraboard/keyauth/pkg/system/notify"
)

// newTenCentSMSSender todo
func newTenCentSMSSender(conf *TenCentConfig) (notify.SMSSender, error) {
	if err := conf.Validate(); err != nil {
		return nil, fmt.Errorf("validate tencent sms config error, %s", err)
	}

	s := &tencent{
		TenCentConfig: conf,
		log:           zap.L().Named("TenCent SMS"),
	}
	if err := s.init(); err != nil {
		return nil, err
	}
	return s, nil
}

type tencent struct {
	*TenCentConfig

	sms *sms.Client
	log logger.Logger
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
	// 补充默认模板ID
	if req.TemplateID == "" {
		req.TemplateID = s.TemplateID
	}

	// 补充默认+86
	req.InjectDefaultIsoCode()

	if err := req.Validate(); err != nil {
		return fmt.Errorf("validate send sms request error, %s", err)
	}

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

	for i := range response.Response.SendStatusSet {
		if strings.ToUpper(*(response.Response.SendStatusSet[i].Code)) != "OK" {
			return fmt.Errorf("send sms error, response is %s", response.ToJsonString())
		}
	}

	s.log.Debugf("send sms response success: %s", response.ToJsonString())
	return nil
}
