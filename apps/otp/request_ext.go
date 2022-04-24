package otp

import "github.com/go-playground/validator/v10"

// CreateOTPAuth(context.Context, *CreateOTPAuthRequest) (*OTPAuth, error)
// DescribeOTPAuth(context.Context, *DescribeOTPAuthRequest) (*OTPAuth, error)
// DeleteOTPAuth(context.Context, *DeleteOTPAuthRequest) (*OTPAuth, error)
// PatchOTPAuth(context.Context, *PatchOTPAuthRequest) (*OTPAuth, error)

var (
	validate = validator.New()
)

func NewCreateOTPAuthRequestWithName(accountname string) *CreateOTPAuthRequest {
	return &CreateOTPAuthRequest{
		Account: accountname,
	}
}

func NewCreateOTPAuthRequest() *CreateOTPAuthRequest {
	return &CreateOTPAuthRequest{}
}

func (req *CreateOTPAuthRequest) Validate() error {
	return validate.Struct(req)
}

func NewDescribeOTPAuthRequestWithName(accountname string) *DescribeOTPAuthRequest {
	return &DescribeOTPAuthRequest{
		Account: accountname,
	}
}

func (req *DescribeOTPAuthRequest) Validate() error {
	return validate.Struct(req)
}
func NewDeleteOTPAuthRequestWithName(accountname string) *DeleteOTPAuthRequest {
	return &DeleteOTPAuthRequest{
		Account: accountname,
	}
}

func (req *DeleteOTPAuthRequest) Validate() error {
	return validate.Struct(req)
}

// func NewEnableOTPAuthRequest() *EnableOTPAuthRequest {
// 	return &EnableOTPAuthRequest{}
// }

// func (req *EnableOTPAuthRequest) Validate() error {
// 	return validate.Struct(req)
// }

// func NewDisableOTPAuthRequest() *DisableOTPAuthRequest {
// 	return &DisableOTPAuthRequest{}
// }

// func (req *DisableOTPAuthRequest) Validate() error {
// 	return validate.Struct(req)
// }
func NewUpdateOTPStatusRequest() *UpdateOTPStatusRequest {
	return &UpdateOTPStatusRequest{}
}

func (req *UpdateOTPStatusRequest) Validate() error {
	return validate.Struct(req)
}
