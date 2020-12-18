package notify

import "strings"

// SMSSender 短信投递
type SMSSender interface {
	Send(*SendSMSRequest) error
}

// NewSendSMSRequest todo
func NewSendSMSRequest() *SendSMSRequest {
	return &SendSMSRequest{}
}

// SendSMSRequest todo
type SendSMSRequest struct {
	TemplateID     string   `json:"template_id"`
	ParamSet       []string `json:"param_set"`
	PhoneNumberSet []string `json:"phone_number_set" validate:"required"`
}

// AddParams todo
func (req *SendSMSRequest) AddParams(params ...string) {
	req.ParamSet = append(req.ParamSet, params...)
}

// AddPhone todo
func (req *SendSMSRequest) AddPhone(params ...string) {
	req.PhoneNumberSet = append(req.PhoneNumberSet, params...)
}

// Validate todo
func (req *SendSMSRequest) Validate() error {
	return validate.Struct(req)
}

// InjectDefaultIsoCode todo
func (req *SendSMSRequest) InjectDefaultIsoCode() {
	for i, number := range req.PhoneNumberSet {
		if !strings.HasPrefix(number, "+") {
			req.PhoneNumberSet[i] = "+86" + number
		}
	}
}
