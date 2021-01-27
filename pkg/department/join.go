package department

import (
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/infraboard/keyauth/pkg/token"
)

// NewApplicationForm todo
func NewApplicationForm(tk *token.Token, req *JoinDepartmentRequest) (*ApplicationForm, error) {
	ins := &ApplicationForm{
		Id:       xid.New().String(),
		Domain:   tk.Domain,
		CreateAt: ftime.Now().Timestamp(),
		UpdateAt: ftime.Now().Timestamp(),
		Creater:  tk.Account,
		Data:     req,
		Status:   ApplicationFormStatus_PENDDING,
	}

	return ins, nil
}

// NewDeafultApplicationForm todo
func NewDeafultApplicationForm() *ApplicationForm {
	return &ApplicationForm{
		Data: NewJoinDepartmentRequest(),
	}
}

// NewDApplicationFormSet 实例化
func NewDApplicationFormSet() *ApplicationFormSet {
	return &ApplicationFormSet{
		Items: []*ApplicationForm{},
	}
}

// Length 个数
func (s *ApplicationFormSet) Length() int {
	return len(s.Items)
}

// Add 添加
func (s *ApplicationFormSet) Add(e *ApplicationForm) {
	s.Items = append(s.Items, e)
}
