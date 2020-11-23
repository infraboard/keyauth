package notify

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
	TemplateID     string
	ParamSet       []string
	PhoneNumberSet []string
}
