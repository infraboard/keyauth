package ldap

import (
	"github.com/go-ldap/ldap/v3"
)

// Connection interface representing a connection to the ldap.
type Connection interface {
	Bind(username, password string) error
	Close()

	Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error)
	Modify(modifyRequest *ldap.ModifyRequest) error
}

// ConnectionImpl the production implementation of an ldap connection.
type ConnectionImpl struct {
	conn *ldap.Conn
}

// NewLDAPConnectionImpl create a new ldap connection.
func NewLDAPConnectionImpl(conn *ldap.Conn) *ConnectionImpl {
	return &ConnectionImpl{conn}
}

// Bind binds ldap connection to a username/password.
func (lc *ConnectionImpl) Bind(username, password string) error {
	return lc.conn.Bind(username, password)
}

// Close closes a ldap connection.
func (lc *ConnectionImpl) Close() {
	lc.conn.Close()
}

// Search searches a ldap server.
func (lc *ConnectionImpl) Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error) {
	return lc.conn.Search(searchRequest)
}

// Modify modifies an ldap object.
func (lc *ConnectionImpl) Modify(modifyRequest *ldap.ModifyRequest) error {
	return lc.conn.Modify(modifyRequest)
}
