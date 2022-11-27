package session

import "context"

type Storage interface {
	GetSession(ctx context.Context, sessionId int64) (Session, bool, error)
}
