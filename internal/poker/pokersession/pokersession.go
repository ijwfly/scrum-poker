package pokersession

import (
	"scrum-poker/internal/poker"

	"sync"
)

type PokerSessionWrapper struct {
	mu sync.Mutex
	poker.Session
}

func NewSessionWrapper(sessionId string) *PokerSessionWrapper {
	var sessionMutex PokerSessionWrapper
	sessionMutex.Session = poker.NewSession(sessionId)
	return &sessionMutex
}

type PokerSession struct {
	mu      sync.Mutex
	storage map[string]*PokerSessionWrapper
}

func NewMemPokerSession() *PokerSession {
	var sessionStorageController PokerSession
	sessionStorageController.storage = make(map[string]*PokerSessionWrapper)
	return &sessionStorageController
}

func (s *PokerSession) getOrCreateSessionWrapper(sessionId string) *PokerSessionWrapper {
	if sessionWrapper, ok := s.storage[sessionId]; ok {
		return sessionWrapper
	} else {
		newSessionWrapper := NewSessionWrapper(sessionId)
		s.mu.Lock()
		s.storage[sessionId] = newSessionWrapper
		s.mu.Unlock()
		return newSessionWrapper
	}
}

func (s *PokerSession) GetOrCreateSession(sessionId string) poker.Session {
	return s.getOrCreateSessionWrapper(sessionId).Session
}

func (s *PokerSession) UserJoinSession(user poker.User, sessionId string) (poker.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj, _ := poker.UserJoinSession(sessionWrapper.Session, user)
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}

func (s *PokerSession) UserLeaveSession(user poker.User, sessionId string) (poker.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj := poker.UserLeaveSession(sessionWrapper.Session, user)
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}

func (s *PokerSession) UserSetEstimate(user poker.User, sessionId string, estimate poker.EstimateOption) (poker.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj, err := poker.UserSetEstimate(sessionWrapper.Session, user, estimate)
	if err != nil {
		return sessionObj, err
	}
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}

func (s *PokerSession) ResetEstimates(sessionId string) (poker.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj := poker.ResetEstimates(sessionWrapper.Session)
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}

func (s *PokerSession) ShowEstimatesToggle(sessionId string) (poker.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj := poker.ShowEstimatesToggle(sessionWrapper.Session)
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}
