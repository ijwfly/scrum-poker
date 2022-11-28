package storagecontroller

import (
	"scrum-poker/internal/app/poker"

	"sync"
)

type SessionWrapper struct {
	mu sync.Mutex
	poker.Session
}

func NewSessionWrapper(sessionId string) *SessionWrapper {
	var sessionMutex SessionWrapper
	sessionMutex.Session = poker.NewSession(sessionId)
	return &sessionMutex
}

type SessionStorageController struct {
	mu      sync.Mutex
	storage map[string]*SessionWrapper
}

func NewSessionStorageController() *SessionStorageController {
	var sessionStorageController SessionStorageController
	sessionStorageController.storage = make(map[string]*SessionWrapper)
	return &sessionStorageController
}

func (s *SessionStorageController) getOrCreateSessionWrapper(sessionId string) *SessionWrapper {
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

func (s *SessionStorageController) GetOrCreateSession(sessionId string) poker.Session {
	return s.getOrCreateSessionWrapper(sessionId).Session
}

func (s *SessionStorageController) UserJoinSession(user poker.User, sessionId string) (poker.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj, _ := poker.UserJoinSession(sessionWrapper.Session, user)
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}

func (s *SessionStorageController) UserLeaveSession(user poker.User, sessionId string) (poker.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj := poker.UserLeaveSession(sessionWrapper.Session, user)
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}

func (s *SessionStorageController) UserSetEstimate(user poker.User, sessionId string, estimate poker.EstimateOption) (poker.Session, error) {
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

func (s *SessionStorageController) SessionResetEstimates(sessionId string) (poker.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj := poker.ResetEstimates(sessionWrapper.Session)
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}

func (s *SessionStorageController) SessionShowEstimatesToggle(sessionId string) (poker.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj := poker.ShowEstimatesToggle(sessionWrapper.Session)
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}
