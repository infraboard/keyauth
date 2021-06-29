package session

import (
	"fmt"
	"sync"

	"github.com/infraboard/keyauth/pkg/token"
)

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		store: make(map[string]*session),
	}
}

type MemoryStore struct {
	store map[string]*session
	m     sync.Mutex
}

func (s *MemoryStore) GetToken(token string) *token.Token {
	if v, ok := s.store[token]; ok {
		return v.tk
	}
	return nil
}

func (s *MemoryStore) SetToken(tk *token.Token) error {
	if tk == nil {
		return nil
	}

	s.m.Lock()
	defer s.m.Unlock()

	if v, ok := s.store[tk.AccessToken]; ok {
		v.Inc()
	} else {
		s.store[tk.AccessToken] = newSession(tk)
	}
	return nil
}

// 租把这个token租出去
func (s *MemoryStore) LeaseToken(token string) *token.Token {
	s.m.Lock()
	defer s.m.Unlock()

	if v, ok := s.store[token]; ok {
		v.Inc()
		return s.store[token].tk
	}

	return nil
}

// 还回来
func (s *MemoryStore) ReturnToken(tk *token.Token) {
	if tk == nil {
		return
	}

	s.m.Lock()
	defer s.m.Unlock()

	if v, ok := s.store[tk.AccessToken]; ok {
		if v.Dec() == 0 {
			delete(s.store, tk.AccessToken)
		}
	}
}

func newSession(tk *token.Token) *session {
	return &session{
		tk:       tk,
		refCount: 1,
	}
}

type session struct {
	refCount uint
	tk       *token.Token
}

func (s *session) Inc() uint {
	s.refCount++
	fmt.Println("inc", s.refCount)
	return s.refCount
}

func (s *session) Dec() uint {
	s.refCount--
	fmt.Println("dec", s.refCount)
	return s.refCount
}
