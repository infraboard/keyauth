package department

import (
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

// NewApplicationForm todo
func NewApplicationForm(req *JoinDepartmentRequest) (*ApplicationForm, error) {
	tk := req.GetToken()

	ins := &ApplicationForm{
		ID:                    xid.New().String(),
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
	ID       string                `bson:"_id" json:"id"`              // 申请单ID
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

// Length 个数
func (s *ApplicationFormSet) Length() int {
	return len(s.Items)
}

// Add 添加
func (s *ApplicationFormSet) Add(e *ApplicationForm) {
	s.Items = append(s.Items, e)
}
