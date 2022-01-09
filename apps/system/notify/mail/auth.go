package mail

import (
	"errors"
	"net/smtp"
	"strings"
)

type loginAuth struct {
	username string
	password string
}

func newLoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

// Used for AUTH LOGIN. (Maybe password should be encrypted)
func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch strings.ToLower(string(fromServer)) {
		case "username:":
			return []byte(a.username), nil
		case "password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unexpected server challenge")
		}
	}
	return nil, nil
}
