package ldap

import (
	"fmt"
	"strings"
)

// NewDefaultConfig represents the default LDAP config.
func NewDefaultConfig() *Config {
	return &Config{
		MailAttribute:        "mail",
		DisplayNameAttribute: "displayName",
		GroupNameAttribute:   "cn",
		UsernameAttribute:    "uid",
		UsersFilter:          "(uid={input})",
		GroupsFilter:         "(|(member={dn})(uid={username})(uid={input}))",
	}
}

// Config represents the configuration related to LDAP server.
type Config struct {
	URL                  string `bson:"url" json:"url"`
	SkipVerify           bool   `bson:"skip_verify" json:"skip_verify"`
	BaseDN               string `bson:"base_dn" json:"base_dn"`
	AdditionalUsersDN    string `bson:"additional_users_dn" json:"additional_users_dn"`
	UsersFilter          string `bson:"users_filter" json:"users_filter"`
	AdditionalGroupsDN   string `bson:"additional_groups_dn" json:"additional_groups_dn"`
	GroupsFilter         string `bson:"groups_filter" json:"groups_filter"`
	GroupNameAttribute   string `bson:"group_name_attribute" json:"group_name_attribute"`
	UsernameAttribute    string `bson:"username_attribute" json:"username_attribute"`
	MailAttribute        string `bson:"mail_attribute" json:"mail_attribute"`
	DisplayNameAttribute string `bson:"display_name_attribute" json:"display_name_attribute"`
	User                 string `bson:"user" json:"user"`
	Password             string `bson:"password" json:"password"`
}

// GetBaseDNFromUser 从用户中获取BaseDN
func (c *Config) GetBaseDNFromUser() string {
	baseDN := []string{}
	for _, item := range strings.Split(c.User, ",") {
		if !strings.HasPrefix(item, "cn=") {
			baseDN = append(baseDN, item)
		}
	}

	return strings.Join(baseDN, ",")
}

// Validate todo
func (c *Config) Validate() error {
	if c.URL == "" {
		return fmt.Errorf("url required")
	}

	if c.User == "" || c.Password == "" {
		return fmt.Errorf("user and password required")
	}

	return nil
}

// Desensitize todo
func (c *Config) Desensitize() {
	c.Password = ""
}
