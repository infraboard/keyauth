package provider

import (
	"fmt"

	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg/provider/ldap"
	"github.com/infraboard/keyauth/pkg/token"
)

// LDAP todo
type LDAP interface {
	SaveConfig(*SaveLDAPConfigRequest) (*LDAPConfig, error)
	QueryConfig(*QueryLDAPConfigRequest) (*LDAPSet, error)
	DescribeConfig(*DescribeLDAPConfig) (*LDAPConfig, error)
	DeleteConfig(*DeleteLDAPConfig) error
}

// NewSaveLDAPConfigRequest todo
func NewSaveLDAPConfigRequest() *SaveLDAPConfigRequest {
	return &SaveLDAPConfigRequest{
		Session: token.NewSession(),
		Enabled: true,
		Config:  ldap.NewDefaultConfig(),
	}
}

// SaveLDAPConfigRequest todo
type SaveLDAPConfigRequest struct {
	Enabled        bool `bson:"enabled" json:"enabled"`
	*ldap.Config   `bson:",inline"`
	*token.Session `bson"-" json:"-"`
}

// Validate todo
func (req *SaveLDAPConfigRequest) Validate() error {
	return nil
}

// NewQueryLDAPConfigRequest todo
func NewQueryLDAPConfigRequest(pageReq *request.PageRequest) *QueryLDAPConfigRequest {
	return &QueryLDAPConfigRequest{
		Session:     token.NewSession(),
		PageRequest: pageReq,
	}
}

// QueryLDAPConfigRequest 查询LDAP配置
type QueryLDAPConfigRequest struct {
	*token.Session
	*request.PageRequest
}

// NewDescribeLDAPConfigWithBaseDN todo
func NewDescribeLDAPConfigWithBaseDN(baseDN string) *DescribeLDAPConfig {
	return &DescribeLDAPConfig{
		BaseDN: baseDN,
	}
}

// NewDescribeLDAPConfigWithID todo
func NewDescribeLDAPConfigWithID(id string) *DescribeLDAPConfig {
	return &DescribeLDAPConfig{
		ID: id,
	}
}

// DescribeLDAPConfig 描述配置
type DescribeLDAPConfig struct {
	ID     string
	BaseDN string
}

// Validate todo
func (req *DescribeLDAPConfig) Validate() error {
	if req.ID == "" && req.BaseDN == "" {
		return fmt.Errorf("id or base_dn required")
	}

	return nil
}

// DeleteLDAPConfig todo
type DeleteLDAPConfig struct {
	ID string
}
