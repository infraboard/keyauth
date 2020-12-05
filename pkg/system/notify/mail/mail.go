package mail

import (
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/infraboard/keyauth/pkg/system/notify"
)

// NewSender todo
func NewSender(conf *Config) (notify.MailSender, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	return &sender{Config: conf}, nil
}

type sender struct {
	*Config
}

// Send 发送邮件
func (s *sender) Send(req *notify.SendMailRequest) error {
	c, err := s.client()
	if err != nil {
		return err
	}

	// 设置发信人
	from, err := mail.ParseAddress(s.From)
	if err != nil {
		return fmt.Errorf("parsing from addresses: %s", err)
	}
	if err := c.Mail(from.Address); err != nil {
		return fmt.Errorf("sending mail from: %s", err)
	}

	// 设置收信人
	toAddrs, err := mail.ParseAddressList(req.To)
	if err != nil {
		return fmt.Errorf("parsing to addresses: %s", err)
	}
	for _, addr := range toAddrs {
		if err := c.Rcpt(addr.Address); err != nil {
			return fmt.Errorf("sending rcpt to: %s", err)
		}
	}

	msg, err := req.PrepareBody()
	if err != nil {
		return fmt.Errorf("parpare body error")
	}

	// 设置内容
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func (s *sender) auth(mechs string) (smtp.Auth, error) {
	var (
		errStr []string
	)

	for _, mech := range strings.Split(mechs, " ") {
		switch mech {
		case "CRAM-MD5":
			if s.AuthSecret == "" {
				errStr = append(errStr, "missing secret for CRAM-MD5 auth mechanism")
				continue
			}
			return smtp.CRAMMD5Auth(s.AuthUserName, s.AuthSecret), nil
		case "PLAIN":
			if s.AuthPassword == "" {
				errStr = append(errStr, "missing password for PLAIN auth mechanism")
				continue
			}
			identity := s.AuthIdentity

			// We need to know the hostname for both auth and TLS.
			host, _, err := net.SplitHostPort(s.Host)
			if err != nil {
				return nil, fmt.Errorf("invalid address: %s", err)
			}
			return smtp.PlainAuth(identity, s.AuthUserName, s.AuthPassword, host), nil
		case "LOGIN":
			if s.AuthPassword == "" {
				errStr = append(errStr, "missing password for LOGIN auth mechanism")
				continue
			}
			return newLoginAuth(s.AuthUserName, s.AuthPassword), nil
		case "NTLM":
			errStr = append(errStr, "NTLM auth not impliment")
		}

	}

	if len(errStr) == 0 {
		errStr = append(errStr, "unknown auth mechanism: "+mechs)
	}

	return nil, errors.New(strings.Join(errStr, ","))
}

func (s *sender) client() (*smtp.Client, error) {
	// We need to know the hostname for both auth and TLS.
	var c *smtp.Client

	host, port, err := net.SplitHostPort(s.Host)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %s", err)
	}

	if port == "465" {
		tlsConfig, err := s.TLSConfig.NewTLSConfig()
		if err != nil {
			return nil, err
		}
		if tlsConfig.ServerName == "" {
			tlsConfig.ServerName = host
		}

		conn, err := s.TLSConfig.Connect(s.Host)
		if err != nil {
			return nil, err
		}
		c, err = smtp.NewClient(conn, host)
		if err != nil {
			return nil, err
		}
	} else {
		// Connect to the SMTP smarthost.
		c, err = smtp.Dial(s.Host)
		if err != nil {
			return nil, err
		}
	}

	// 发送自定义Hello数据
	if s.Hello != "" {
		err := c.Hello(s.Hello)
		if err != nil {
			return nil, err
		}
	}

	// Global Config guarantees RequireTLS is not nil.
	if s.RequireTLS {
		if ok, _ := c.Extension("STARTTLS"); !ok {
			return nil, fmt.Errorf("require_tls: true (default), but %q does not advertise the STARTTLS extension", s.Host)
		}

		tlsConf, err := s.TLSConfig.NewTLSConfig()
		if err != nil {
			return nil, err
		}

		if tlsConf.ServerName == "" {
			tlsConf.ServerName = host
		}

		if err := c.StartTLS(tlsConf); err != nil {
			return nil, fmt.Errorf("starttls failed: %s", err)
		}
	}

	// 根据服务器推荐的方式选择认证方式
	if !s.SkipAuth {
		if ok, mech := c.Extension("AUTH"); ok {
			auth, err := s.auth(mech)
			if err != nil {
				return nil, err
			}
			if auth != nil {
				if err := c.Auth(auth); err != nil {
					return nil, fmt.Errorf("%T failed: %s", auth, err)
				}
			}
		}
	}

	return c, nil
}
