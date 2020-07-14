package ldap_test

import (
	"testing"

	"github.com/infraboard/keyauth/pkg/token/ldap"
)

var testConfig struct {
	Server       string
	Port         int
	TLSPort      int
	BindUPN      string
	BindPass     string
	BindSecurity ldap.SecurityType
	BaseDN       string
	PasswordUPN  string
}

func TestConfigDomain(t *testing.T) {
	tests := []string{
		"dc=example,dc=com",
		"ou=test,dc=example,dc=com",
		"dc=example, dc=com",
		"DC=example,DC=com",
		"OU=test,dc=example,DC=com",
	}
	for _, test := range tests {
		if domain, err := (&ldap.Config{BaseDN: test}).Domain(); domain != "example.com" {
			if err != nil {
				t.Error("Failed Test:", test, "\n\tError:", err)
			} else {
				t.Error("Failed Test:", test, "\n\tOutput:", domain, "Expected: example.com")
			}
		}
	}
	errorTests := []string{
		"",
		"com",
		"ou=test",
		"OU=test",
	}
	for _, test := range errorTests {
		if _, err := (&ldap.Config{BaseDN: test}).Domain(); err == nil {
			t.Error("Failed Test:", test, "\n\tError: err not nil")
		}
	}
}

func TestConfigUPN(t *testing.T) {
	dnTests := []string{
		"dc=example,dc=com",
		"ou=test,dc=example,dc=com",
		"dc=example, dc=com",
		"DC=example,DC=com",
		"OU=test,dc=example,DC=com",
	}
	userTests := []string{
		"example.user",
		"example.user@example.com",
	}
	for _, dn := range dnTests {
		for _, user := range userTests {
			if upn, err := (&ldap.Config{BaseDN: dn}).UPN(user); upn != "example.user@example.com" {
				if err != nil {
					t.Error("Failed Test:", dn, user, "\n\tError:", err)
				} else {
					t.Error("Failed Test:", dn, user, "\n\tOutput:", upn, "Expected: example.user@example.com")
				}
			}
		}
	}
	if _, err := (&ldap.Config{BaseDN: "Bad OU"}).UPN("Test"); err == nil {
		t.Error("Expected configuration error but got nil")
	}
}

func init() {
	testConfig.Server = "127.0.0.1"
	testConfig.Port = 389
	testConfig.BindUPN = "cn=admin,dc=example,dc=org"
	testConfig.BindPass = "admin"
	testConfig.BindSecurity = ldap.SecurityNone
	testConfig.BaseDN = "dc=example,dc=org"
}
