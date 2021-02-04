package department

import (
	"fmt"
	"hash/fnv"

	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/pkg/token"
)

// NewApplicationForm todo
func NewApplicationForm(tk *token.Token, req *JoinDepartmentRequest) (*ApplicationForm, error) {
	// 计算申请单Hash
	hashedStr := fmt.Sprintf("%s-%s-%s", tk.Domain, req.Account, req.DepartmentId)
	h := fnv.New32a()
	h.Write([]byte(hashedStr))
	hashID := fmt.Sprintf("%x", h.Sum32())

	ins := &ApplicationForm{
		Id:           hashID,
		Domain:       tk.Domain,
		CreateAt:     ftime.Now().Timestamp(),
		UpdateAt:     ftime.Now().Timestamp(),
		Creater:      tk.Account,
		Account:      req.Account,
		DepartmentId: req.DepartmentId,
		Message:      req.Message,
		Status:       ApplicationFormStatus_PENDDING,
	}

	return ins, nil
}

// NewDeafultApplicationForm todo
func NewDeafultApplicationForm() *ApplicationForm {
	return &ApplicationForm{}
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
