package session

import (
	"scrum-poker/internal/app/room"
)

type OnlineState string

const (
	Online  OnlineState = "online"
	Offline OnlineState = "offline"
)

type UserState struct {
	UserName string
	Estimate *room.EstimateOption
	Presence OnlineState
}

type Session struct {
	SessionId              int64
	RoomId                 int64
	SessionEstimateOptions []room.EstimateOption
	ShowEstimates          bool
	Users                  map[int64]UserState
}
