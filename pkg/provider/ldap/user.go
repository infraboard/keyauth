package ldap

// UserProfile todo
type UserProfile struct {
	DN       string
	Emails   []string
	Username string
}

// UserDetails represent the details retrieved for a given user.
type UserDetails struct {
	Username    string
	DisplayName string
	Emails      []string
	Groups      []string
}
