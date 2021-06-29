package session_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/infraboard/keyauth/client/session"
	"github.com/infraboard/keyauth/pkg/token"
)

func TestMemStore(t *testing.T) {
	s := session.NewMemoryStore()
	tk := &token.Token{
		AccessToken: "abc",
	}

	s.SetToken(tk)
	go func() {
		tk := s.LeaseToken("abc")
		s.ReturnToken(tk)
	}()
	go func() {
		tk := s.LeaseToken("abc")
		s.ReturnToken(tk)
	}()

	time.Sleep(2 * time.Second)
	s.ReturnToken(tk)
	fmt.Println(s.GetToken("abc"))
}
