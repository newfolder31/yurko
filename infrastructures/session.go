package infrastructures

import (
	"crypto/rand"
	"fmt"
)

const (
	COOKIE_NAME = "sessionId"
)

var inMemorySession *Session

type sessionData struct {
	Email string
}

type Session struct {
	data map[string]*sessionData
}

func NewSession() *Session {
	s := new(Session)

	s.data = make(map[string]*sessionData)

	return s
}

func (s *Session) Init(email string) string {
	sessionId := generateSessionId()

	data := &sessionData{Email: email}
	s.data[sessionId] = data

	return sessionId
}

func (s *Session) Get(sessionId string) string {
	data := s.data[sessionId]

	if data == nil {
		return ""
	}

	return data.Email
}

func (s *Session) Delete(sessionId string) {
	delete(s.data, sessionId)
}

func generateSessionId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
