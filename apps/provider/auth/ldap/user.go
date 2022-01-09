package ldap

// UserProfile todo
type UserProfile struct {
	DN          string
	Emails      []string
	Username    string
	DisplayName string
	Groups      []string
}
