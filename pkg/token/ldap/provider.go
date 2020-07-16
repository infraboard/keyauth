package ldap

import (
	"crypto/tls"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// OWASP recommends to escape some special characters.
// https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/LDAP_Injection_Prevention_Cheat_Sheet.md
const specialLDAPRunes = ",#+<>;\"="

// Provider LDAP provider
type Provider interface {
	CheckUserPassword(username string, password string) (bool, error)
	GetDetails(username string) (*UserDetails, error)
	UpdatePassword(username string, newPassword string) error
}

// NewProvider todo
func NewProvider(conf *Config) Provider {
	return &ldapProvider{
		conf: conf,
		log:  zap.L().Named("LDAP"),
	}
}

type ldapProvider struct {
	conf *Config
	log  logger.Logger
}

func (p *ldapProvider) dialTLS(network, addr string, config *tls.Config) (Connection, error) {
	conn, err := ldap.DialTLS(network, addr, config)
	if err != nil {
		return nil, err
	}

	return NewLDAPConnectionImpl(conn), nil
}

func (p *ldapProvider) dial(network, addr string) (Connection, error) {
	conn, err := ldap.Dial(network, addr)
	if err != nil {
		return nil, err
	}

	return NewLDAPConnectionImpl(conn), nil
}

func (p *ldapProvider) connect(userDN string, password string) (Connection, error) {
	var conn Connection

	url, err := url.Parse(p.conf.URL)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse URL to LDAP: %s", url)
	}

	if url.Scheme == "ldaps" {
		p.log.Debug("LDAP client starts a TLS session")
		tlsConn, err := p.dialTLS("tcp", url.Host, &tls.Config{
			InsecureSkipVerify: p.conf.SkipVerify,
		})
		if err != nil {
			return nil, err
		}

		conn = tlsConn
	} else {
		p.log.Debug("LDAP client starts a session over raw TCP")
		rawConn, err := p.dial("tcp", url.Host)
		if err != nil {
			return nil, err
		}
		conn = rawConn
	}

	if err := conn.Bind(userDN, password); err != nil {
		return nil, err
	}

	return conn, nil
}

// CheckUserPassword checks if provided password matches for the given user.
func (p *ldapProvider) CheckUserPassword(inputUsername string, password string) (bool, error) {
	adminClient, err := p.connect(p.conf.User, p.conf.Password)
	if err != nil {
		return false, err
	}
	defer adminClient.Close()

	profile, err := p.getUserProfile(adminClient, inputUsername)
	if err != nil {
		return false, err
	}

	conn, err := p.connect(profile.DN, password)
	if err != nil {
		return false, fmt.Errorf("Authentication of user %s failed. Cause: %s", inputUsername, err)
	}
	defer conn.Close()

	return true, nil
}

func (p *ldapProvider) ldapEscape(inputUsername string) string {
	inputUsername = ldap.EscapeFilter(inputUsername)
	for _, c := range specialLDAPRunes {
		inputUsername = strings.ReplaceAll(inputUsername, string(c), fmt.Sprintf("\\%c", c))
	}

	return inputUsername
}

func (p *ldapProvider) resolveUsersFilter(userFilter string, inputUsername string) string {
	inputUsername = p.ldapEscape(inputUsername)

	// We temporarily keep placeholder {0} for backward compatibility.
	userFilter = strings.ReplaceAll(userFilter, "{0}", inputUsername)

	// The {username} placeholder is equivalent to {0}, it's the new way, a named placeholder.
	userFilter = strings.ReplaceAll(userFilter, "{input}", inputUsername)

	// {username_attribute} and {mail_attribute} are replaced by the content of the attribute defined
	// in configuration.
	userFilter = strings.ReplaceAll(userFilter, "{username_attribute}", p.conf.UsernameAttribute)
	userFilter = strings.ReplaceAll(userFilter, "{mail_attribute}", p.conf.MailAttribute)

	return userFilter
}

func (p *ldapProvider) getUserProfile(conn Connection, inputUsername string) (*UserProfile, error) {
	userFilter := p.resolveUsersFilter(p.conf.UsersFilter, inputUsername)
	p.log.Debugf("Computed user filter is %s", userFilter)

	baseDN := p.conf.BaseDN
	if p.conf.AdditionalUsersDN != "" {
		baseDN = p.conf.AdditionalUsersDN + "," + baseDN
	}

	attributes := []string{"dn",
		p.conf.MailAttribute,
		p.conf.UsernameAttribute}

	// Search for the given username.
	searchRequest := ldap.NewSearchRequest(
		baseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,
		1, 0, false, userFilter, attributes, nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("Cannot find user DN of user %s. Cause: %s", inputUsername, err)
	}

	if len(sr.Entries) == 0 {
		return nil, exception.NewNotFound("user not found")
	}

	if len(sr.Entries) > 1 {
		return nil, fmt.Errorf("Multiple users %s found", inputUsername)
	}

	userProfile := UserProfile{
		DN: sr.Entries[0].DN,
	}

	for _, attr := range sr.Entries[0].Attributes {
		if attr.Name == p.conf.MailAttribute {
			userProfile.Emails = attr.Values
		}

		if attr.Name == p.conf.UsernameAttribute {
			if len(attr.Values) != 1 {
				return nil, fmt.Errorf("User %s cannot have multiple value for attribute %s",
					inputUsername, p.conf.UsernameAttribute)
			}

			userProfile.Username = attr.Values[0]
		}
	}

	if userProfile.DN == "" {
		return nil, fmt.Errorf("No DN has been found for user %s", inputUsername)
	}

	return &userProfile, nil
}

func (p *ldapProvider) resolveGroupsFilter(inputUsername string, profile *UserProfile) (string, error) { //nolint:unparam
	inputUsername = p.ldapEscape(inputUsername)

	// We temporarily keep placeholder {0} for backward compatibility.
	groupFilter := strings.ReplaceAll(p.conf.GroupsFilter, "{0}", inputUsername)
	groupFilter = strings.ReplaceAll(groupFilter, "{input}", inputUsername)

	if profile != nil {
		// We temporarily keep placeholder {1} for backward compatibility.
		groupFilter = strings.ReplaceAll(groupFilter, "{1}", ldap.EscapeFilter(profile.Username))
		groupFilter = strings.ReplaceAll(groupFilter, "{username}", ldap.EscapeFilter(profile.Username))
		groupFilter = strings.ReplaceAll(groupFilter, "{dn}", ldap.EscapeFilter(profile.DN))
	}

	return groupFilter, nil
}

// GetDetails retrieve the groups a user belongs to.
func (p *ldapProvider) GetDetails(inputUsername string) (*UserDetails, error) {
	conn, err := p.connect(p.conf.User, p.conf.Password)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	profile, err := p.getUserProfile(conn, inputUsername)
	if err != nil {
		return nil, err
	}

	groupsFilter, err := p.resolveGroupsFilter(inputUsername, profile)
	if err != nil {
		return nil, fmt.Errorf("Unable to create group filter for user %s. Cause: %s", inputUsername, err)
	}

	p.log.Debugf("Computed groups filter is %s", groupsFilter)

	groupBaseDN := p.conf.BaseDN
	if p.conf.AdditionalGroupsDN != "" {
		groupBaseDN = p.conf.AdditionalGroupsDN + "," + groupBaseDN
	}

	// Search for the given username.
	searchGroupRequest := ldap.NewSearchRequest(
		groupBaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,
		0, 0, false, groupsFilter, []string{p.conf.GroupNameAttribute}, nil,
	)

	sr, err := conn.Search(searchGroupRequest)

	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve groups of user %s. Cause: %s", inputUsername, err)
	}

	groups := make([]string, 0)

	for _, res := range sr.Entries {
		if len(res.Attributes) == 0 {
			p.log.Warnf("No groups retrieved from LDAP for user %s", inputUsername)
			break
		}
		// Append all values of the document. Normally there should be only one per document.
		groups = append(groups, res.Attributes[0].Values...)
	}

	return &UserDetails{
		Username: profile.Username,
		Emails:   profile.Emails,
		Groups:   groups,
	}, nil
}

// UpdatePassword update the password of the given user.
func (p *ldapProvider) UpdatePassword(inputUsername string, newPassword string) error {
	client, err := p.connect(p.conf.User, p.conf.Password)

	if err != nil {
		return fmt.Errorf("Unable to update password. Cause: %s", err)
	}

	profile, err := p.getUserProfile(client, inputUsername)

	if err != nil {
		return fmt.Errorf("Unable to update password. Cause: %s", err)
	}

	modifyRequest := ldap.NewModifyRequest(profile.DN, nil)

	modifyRequest.Replace("userPassword", []string{newPassword})

	err = client.Modify(modifyRequest)

	if err != nil {
		return fmt.Errorf("Unable to update password. Cause: %s", err)
	}

	return nil
}
