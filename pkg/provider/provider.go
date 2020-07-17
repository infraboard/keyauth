package provider

// NewDefaultLDAPConfig represents the default LDAP config.
func NewDefaultLDAPConfig() *LDAPConfig {
	return &LDAPConfig{
		URL:                  "ldap://127.0.0.1:389",
		MailAttribute:        "mail",
		DisplayNameAttribute: "displayname",
		GroupNameAttribute:   "cn",
		User:                 "cn=admin,dc=example,dc=org",
		Password:             "admin",
		BaseDN:               "dc=example,dc=org",
		UsersFilter:          "(objectclass=simpleSecurityObject)",
		UsernameAttribute:    "uid",
	}
}

// LDAPConfig represents the configuration related to LDAP server.
type LDAPConfig struct {
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
