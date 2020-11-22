package tls

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

// Config configures the options for TLS connections.
type Config struct {
	CAFile             string // The CA cert to use for the targets.
	CertFile           string // The client cert file for the targets.
	KeyFile            string // The client key file for the targets.
	ServerName         string // Used to verify the hostname for the targets.
	InsecureSkipVerify bool   // Disable target certificate validation.
}

// NewTLSConfig creates a new tls.Config from the given
func (c *Config) NewTLSConfig() (*tls.Config, error) {
	tlsConfig := &tls.Config{InsecureSkipVerify: c.InsecureSkipVerify}

	// If a CA cert is provided then let's read it in so we can validate the
	// scrape target's certificate properly.
	if len(c.CAFile) > 0 {
		b, err := readCAFile(c.CAFile)
		if err != nil {
			return nil, err
		}
		if !updateRootCA(tlsConfig, b) {
			return nil, fmt.Errorf("unable to use specified CA cert %s", c.CAFile)
		}
	}

	if len(c.ServerName) > 0 {
		tlsConfig.ServerName = c.ServerName
	}
	// If a client cert & key is provided then configure TLS config accordingly.
	if len(c.CertFile) > 0 && len(c.KeyFile) == 0 {
		return nil, fmt.Errorf("client cert file %q specified without client key file", c.CertFile)
	} else if len(c.KeyFile) > 0 && len(c.CertFile) == 0 {
		return nil, fmt.Errorf("client key file %q specified without client cert file", c.KeyFile)
	} else if len(c.CertFile) > 0 && len(c.KeyFile) > 0 {
		// Verify that client cert and key are valid.
		if _, err := c.getClientCertificate(nil); err != nil {
			return nil, err
		}
		tlsConfig.GetClientCertificate = c.getClientCertificate
	}

	return tlsConfig, nil
}

// Connect todo
func (c *Config) Connect(host string) (*tls.Conn, error) {
	conf, err := c.NewTLSConfig()
	if err != nil {
		return nil, err
	}
	return tls.Dial("tcp", host, conf)
}

// getClientCertificate reads the pair of client cert and key from disk and returns a tls.Certificate.
func (c *Config) getClientCertificate(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("unable to use specified client cert (%s) & key (%s): %s", c.CertFile, c.KeyFile, err)
	}
	return &cert, nil
}

// readCAFile reads the CA cert file from disk.
func readCAFile(f string) ([]byte, error) {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("unable to load specified CA cert %s: %s", f, err)
	}
	return data, nil
}

// updateRootCA parses the given byte slice as a series of PEM encoded certificates and updates tls.Config.RootCAs.
func updateRootCA(cfg *tls.Config, b []byte) bool {
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(b) {
		return false
	}
	cfg.RootCAs = caCertPool
	return true
}
