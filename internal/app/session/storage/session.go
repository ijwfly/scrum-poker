package storage

import (
	"context"
	"scrum-poker/internal/app/session"
)

type SessionStorage struct {
	storage map[int64]session.Session
}

func NewSessionStorage() *SessionStorage {
	return &SessionStorage{
		storage: make(map[int64]session.Session),
	}
}

func (s *SessionStorage) GetSession(ctx context.Context, sessionId int64) (session.Session, bool, error) {
	value, ok := s.storage[sessionId]
	return value, ok, nil
}
