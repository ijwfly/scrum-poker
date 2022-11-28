package poker

type PokerSessionController interface {
	GetOrCreateSession(sessionId string) Session
	UserJoinSession(user User, sessionId string) (Session, error)
	UserLeaveSession(user User, sessionId string) (Session, error)
	UserSetEstimate(user User, sessionId string, estimate EstimateOption) (Session, error)
	ResetEstimates(sessionId string) (Session, error)
	ShowEstimatesToggle(sessionId string) (Session, error)
}
