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

func TestConfigConnect(t *testing.T) {
	if _, err := (&ldap.Config{Server: "127.0.0.1", Port: 1, Security: ldap.SecurityNone}).Connect(); err == nil {
		t.Error("SecurityNone: Expected connect error but got nil")
	}
	if _, err := (&ldap.Config{Server: "127.0.0.1", Port: 1, Security: ldap.SecurityTLS}).Connect(); err == nil {
		t.Error("SecurityTLS: Expected connect error but got nil")
	}
	if _, err := (&ldap.Config{Server: "127.0.0.1", Port: 1, Security: ldap.SecurityStartTLS}).Connect(); err == nil {
		t.Error("SecurityStartTLS: Expected connect error but got nil")
	}
	if _, err := (&ldap.Config{Server: "127.0.0.1", Port: 1, Security: ldap.SecurityInsecureTLS}).Connect(); err == nil {
		t.Error("SecurityInsecureTLS: Expected connect error but got nil")
	}
	if _, err := (&ldap.Config{Server: "127.0.0.1", Port: 1, Security: ldap.SecurityInsecureStartTLS}).Connect(); err == nil {
		t.Error("SecurityInsecureStartTLS: Expected connect error but got nil")
	}

	if _, err := (&ldap.Config{Server: "127.0.0.1", Port: 1, Security: ldap.SecurityType(100)}).Connect(); err == nil {
		t.Error("Invalid Security: Expected configuration error but got nil")
	}

	if testConfig.Server == "" {
		t.Skip("ADTEST_SERVER not set")
		return
	}

	if _, err := (&ldap.Config{Server: testConfig.Server, Port: testConfig.Port, Security: ldap.SecurityNone}).Connect(); err != nil {
		t.Error("SecurityNone: Expected connect error to be nil but got:", err)
	}
	if _, err := (&ldap.Config{Server: testConfig.Server, Port: testConfig.TLSPort, Security: ldap.SecurityTLS}).Connect(); err != nil {
		t.Error("SecurityTLS: Expected connect error to be nil but got:", err)
	}
	if _, err := (&ldap.Config{Server: testConfig.Server, Port: testConfig.Port, Security: ldap.SecurityStartTLS}).Connect(); err != nil {
		t.Error("SecurityStartTLS: Expected connect error to be nil but got:", err)
	}
	if _, err := (&ldap.Config{Server: testConfig.Server, Port: testConfig.TLSPort, Security: ldap.SecurityInsecureTLS}).Connect(); err != nil {
		t.Error("SecurityInsecureTLS: Expected connect error to be nil but got:", err)
	}
	if _, err := (&ldap.Config{Server: testConfig.Server, Port: testConfig.Port, Security: ldap.SecurityInsecureStartTLS}).Connect(); err != nil {
		t.Error("SecurityInsecureStartTLS: Expected connect error to be nil but got:", err)
	}
}

func TestConnBind(t *testing.T) {
	if testConfig.Server == "" {
		t.Skip("ADTEST_SERVER not set")
		return
	}

	config := &ldap.Config{Server: testConfig.Server, Port: testConfig.Port, Security: testConfig.BindSecurity}
	conn, err := config.Connect()
	if err != nil {
		t.Fatal("Error connecting to server:", err)
	}
	defer conn.Conn.Close()

	if status, _ := conn.Bind("test", ""); status {
		t.Error("Empty password: Expected authentication status to be false")
	}

	if status, _ := conn.Bind("go-ad-auth", "invalid_password"); status {
		t.Error("Invalid credentials: Expected authentication status to be false")
	}

	if testConfig.BindUPN == "" || testConfig.BindPass == "" {
		t.Skip("ADTEST_BIND_UPN or ADTEST_BIND_PASS not set")
		return
	}

	if status, _ := conn.Bind(testConfig.BindUPN, testConfig.BindPass); !status {
		t.Error("Valid credentials: Expected authentication status to be true")
	}
}
