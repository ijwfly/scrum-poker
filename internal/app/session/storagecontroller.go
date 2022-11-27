package session

type StorageController interface {
	GetOrCreateSession(sessionId string) Session
	UserJoinSession(user User, sessionId string) (Session, error)
	UserLeaveSession(user User, sessionId string) (Session, error)
	UserSetEstimate(user User, sessionId string, estimate EstimateOption) (Session, error)
	SessionResetEstimates(sessionId string) (Session, error)
	SessionShowEstimatesToggle(sessionId string) (Session, error)
}
