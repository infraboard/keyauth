package provider

// UserProvider LDAP provider
type UserProvider interface {
	CheckUserPassword(username string, password string) (bool, error)
	GetDetails(username string) (*UserDetails, error)
	UpdatePassword(username string, newPassword string) error
}
