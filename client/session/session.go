package session

import (
	"github.com/infraboard/keyauth/pkg/token"
)

var (
	sess Store
)

func S() Getter {
	return sess
}

func SetStroe(s Store) {
	sess = s
}

type Store interface {
	Getter
	Setter
}

type Getter interface {
	GetToken(token string) *token.Token
}

type Setter interface {
	SetToken(*token.Token) error
	LeaseToken(token string) *token.Token
	ReturnToken(*token.Token)
}
