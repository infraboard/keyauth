package department

import (
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
)

// NewApplicationForm todo
func NewApplicationForm(req *JoinDepartmentRequest) (*ApplicationForm, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk := req.GetToken()

	ins := &ApplicationForm{
		CreateAt:              ftime.Now(),
		UpdateAt:              ftime.Now(),
		Creater:               tk.Account,
		JoinDepartmentRequest: req,
		Status:                Pending,
	}

	return ins, nil
}

// NewDeafultApplicationForm todo
func NewDeafultApplicationForm() *ApplicationForm {
	return &ApplicationForm{
		JoinDepartmentRequest: NewJoinDepartmentRequest(),
	}
}

// ApplicationForm todo
type ApplicationForm struct {
	Creater  string                `bson:"creater" json:"creater"`     // 申请人
	CreateAt ftime.Time            `bson:"create_at" json:"create_at"` // 创建时间
	UpdateAt ftime.Time            `bson:"update_at" json:"update_at"` // 更新时间
	Status   ApplicationFormStatus `bson:"status" json:"status"`       // 状态
	*JoinDepartmentRequest
}

// NewDApplicationFormSet 实例化
func NewDApplicationFormSet(req *request.PageRequest) *ApplicationFormSet {
	return &ApplicationFormSet{
		PageRequest: req,
		Items:       []*ApplicationForm{},
	}
}

// ApplicationFormSet todo
type ApplicationFormSet struct {
	*request.PageRequest

	Total int64              `json:"total"`
	Items []*ApplicationForm `json:"items"`
}

// Add 添加
func (s *ApplicationFormSet) Add(e *ApplicationForm) {
	s.Items = append(s.Items, e)
}
