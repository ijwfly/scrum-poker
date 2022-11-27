package storagecontroller

import (
	"scrum-poker/internal/app/session"

	"sync"
)

type SessionWrapper struct {
	mu sync.Mutex
	session.Session
}

func NewSessionWrapper(sessionId string) *SessionWrapper {
	var sessionMutex SessionWrapper
	sessionMutex.Session = session.NewSession(sessionId)
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

func (s *SessionStorageController) GetOrCreateSession(sessionId string) session.Session {
	return s.getOrCreateSessionWrapper(sessionId).Session
}

func (s *SessionStorageController) UserJoinSession(user session.User, sessionId string) (session.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj, _ := session.UserJoinSession(sessionWrapper.Session, user)
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}

func (s *SessionStorageController) UserLeaveSession(user session.User, sessionId string) (session.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj := session.UserLeaveSession(sessionWrapper.Session, user)
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}

func (s *SessionStorageController) UserSetEstimate(user session.User, sessionId string, estimate session.EstimateOption) (session.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj, err := session.UserSetEstimate(sessionWrapper.Session, user, estimate)
	if err != nil {
		return sessionObj, err
	}
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}

func (s *SessionStorageController) SessionResetEstimates(sessionId string) (session.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj := session.ResetEstimates(sessionWrapper.Session)
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}

func (s *SessionStorageController) SessionShowEstimatesToggle(sessionId string) (session.Session, error) {
	sessionWrapper := s.getOrCreateSessionWrapper(sessionId)

	sessionWrapper.mu.Lock()
	defer sessionWrapper.mu.Unlock()

	sessionObj := session.ShowEstimatesToggle(sessionWrapper.Session)
	sessionWrapper.Session = sessionObj

	return sessionObj, nil
}
