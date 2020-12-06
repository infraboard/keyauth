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
	TemplateID     string   `json:"template_id"`
	ParamSet       []string `json:"param_set"`
	PhoneNumberSet []string `json:"phone_number_set" validate:"required"`
}

// Validate todo
func (req *SendSMSRequest) Validate() error {
	return validate.Struct(req)
}
