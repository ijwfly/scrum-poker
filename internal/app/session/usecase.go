package session

import (
	"scrum-poker/internal/app/room"
	"scrum-poker/internal/app/user"
)

func NewSession(sessionId int64, room room.Room) Session {
	var session Session
	session.SessionId = sessionId
	session.RoomId = room.RoomId
	session.SessionEstimateOptions = room.EstimateOptions
	session.Users = make(map[int64]UserState)
	session.ShowEstimates = false
	return session
}

func UserJoinSession(session Session, user user.User) Session {
	if value, ok := session.Users[user.UserId]; !ok {
		var userState UserState
		userState.UserName = user.UserName
		userState.Estimate = nil
		userState.Presence = Online
		session.Users[user.UserId] = userState
	} else {
		value.Presence = Online
		session.Users[user.UserId] = value
	}

	return session
}

func UserLeaveSession(session Session, user user.User) Session {
	if value, ok := session.Users[user.UserId]; ok {
		value.Presence = Offline
		session.Users[user.UserId] = value
	}

	return session
}

func UserSetEstimate(session Session, user user.User, estimate room.EstimateOption) Session {
	for _, option := range session.SessionEstimateOptions {
		if option == estimate {
			if value, ok := session.Users[user.UserId]; ok {
				value.Estimate = &estimate
				session.Users[user.UserId] = value
			}
		}
	}

	return session
}

func SessionResetEstimates(session Session) Session {
	for k, userObj := range session.Users {
		userObj.Estimate = nil
		session.Users[k] = userObj
	}

	session.ShowEstimates = false

	return session
}

func SessionShowEstimatesToggle(session Session) Session {
	session.ShowEstimates = !session.ShowEstimates
	return session
}
