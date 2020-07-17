package ldap

import "fmt"

// NewDefaultConfig represents the default LDAP config.
func NewDefaultConfig() *Config {
	return &Config{
		MailAttribute:        "mail",
		DisplayNameAttribute: "displayname",
		GroupNameAttribute:   "cn",
		UsersFilter:          "(objectclass=simpleSecurityObject)",
		UsernameAttribute:    "uid",
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

// Validate todo
func (c *Config) Validate() error {
	if c.URL == "" || c.BaseDN == "" {
		return fmt.Errorf("url and base_dn required")
	}

	if c.User == "" || c.Password == "" {
		return fmt.Errorf("user and password required")
	}

	return nil
}
