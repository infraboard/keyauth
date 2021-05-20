package provider

import (
	"fmt"

	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/common/dryrun"
	"github.com/infraboard/keyauth/pkg/provider/auth/ldap"
	"github.com/infraboard/keyauth/pkg/token/session"
)

// LDAP todo
type LDAP interface {
	SaveConfig(*SaveLDAPConfigRequest) (*LDAPConfig, error)
	QueryConfig(*QueryLDAPConfigRequest) (*LDAPSet, error)
	DescribeConfig(*DescribeLDAPConfig) (*LDAPConfig, error)
	DeleteConfig(*DeleteLDAPConfig) error
	CheckConnect(*DescribeLDAPConfig) error
}

// NewSaveLDAPConfigRequest todo
func NewSaveLDAPConfigRequest() *SaveLDAPConfigRequest {
	return &SaveLDAPConfigRequest{
		Session: session.NewSession(),
		Enabled: true,
		Config:  ldap.NewDefaultConfig(),
		DryRun:  dryrun.NewDryRun(),
	}
}

// SaveLDAPConfigRequest todo
type SaveLDAPConfigRequest struct {
	Enabled          bool `bson:"enabled" json:"enabled"`
	*ldap.Config     `bson:",inline"`
	*session.Session `bson:"-" json:"-"`
	*dryrun.DryRun   `bson:"-" json:"-"`
}

// Validate todo
func (req *SaveLDAPConfigRequest) Validate() error {
	return nil
}

// NewQueryLDAPConfigRequest todo
func NewQueryLDAPConfigRequest(pageReq *request.PageRequest) *QueryLDAPConfigRequest {
	return &QueryLDAPConfigRequest{
		Session:     session.NewSession(),
		PageRequest: pageReq,
	}
}

// QueryLDAPConfigRequest 查询LDAP配置
type QueryLDAPConfigRequest struct {
	*session.Session
	*request.PageRequest
}

// NewDescribeLDAPConfigWithBaseDN todo
func NewDescribeLDAPConfigWithBaseDN(baseDN string) *DescribeLDAPConfig {
	return &DescribeLDAPConfig{
		BaseDN: baseDN,
	}
}

// NewDescribeLDAPConfigWithDomain todo
func NewDescribeLDAPConfigWithDomain(domain string) *DescribeLDAPConfig {
	return &DescribeLDAPConfig{
		Domain: domain,
	}
}

// DescribeLDAPConfig 描述配置
type DescribeLDAPConfig struct {
	Domain string
	BaseDN string
}

// Validate todo
func (req *DescribeLDAPConfig) Validate() error {
	if req.Domain == "" && req.BaseDN == "" {
		return fmt.Errorf("domain or base_dn required")
	}

	return nil
}

// DeleteLDAPConfig todo
type DeleteLDAPConfig struct {
	ID string
}
