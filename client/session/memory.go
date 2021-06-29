package session

import (
	"sync"

	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		store: make(map[string]*session),
		l:     zap.L().Named("Memory Session"),
	}
}

type MemoryStore struct {
	store map[string]*session
	m     sync.Mutex
	l     logger.Logger
}

func (s *MemoryStore) Debug(l logger.Logger) {
	s.l = l
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

	s.l.Debugf("set token: %s", tk.AccessToken[:8])
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
		s.l.Debugf("lease token: %s %d", token[:8], v.refCount)
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
		s.l.Debugf("return token: %s %d", tk.AccessToken[:8], v.refCount)
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
	return s.refCount
}

func (s *session) Dec() uint {
	s.refCount--
	return s.refCount
}
