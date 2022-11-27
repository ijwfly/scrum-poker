package session

import (
	"errors"
)

func UserJoinSession(session Session, user User) (Session, UserState) {
	var userState UserState
	if value, ok := session.Users[user.Token]; !ok {
		userState.Name = user.Name
		userState.Estimate = nil
		userState.Presence = Online
		session.Users[user.Token] = userState
	} else {
		value.Presence = Online
		session.Users[user.Token] = value
		userState = value
	}

	return session, userState
}

func UserLeaveSession(session Session, user User) Session {
	if value, ok := session.Users[user.Token]; ok {
		value.Presence = Offline
		session.Users[user.Token] = value
	}

	return session
}

func UserSetEstimate(session Session, user User, estimate EstimateOption) (Session, error) {
	estimateFound := false
	for _, option := range session.SessionEstimateOptions {
		if option == estimate {
			estimateFound = true
			break
		}
	}
	if !estimateFound {
		return session, errors.New("estimate option not found")
	}

	if value, ok := session.Users[user.Token]; ok {
		if value.Token != user.Token {

		}
		value.Estimate = &estimate
		session.Users[user.Token] = value
	}

	return session, nil
}

func ResetEstimates(session Session) Session {
	for k, userObj := range session.Users {
		if userObj.Presence == Online {
			userObj.Estimate = nil
			session.Users[k] = userObj
		} else {
			delete(session.Users, k)
		}
	}

	session.ShowEstimates = false

	return session
}

func ShowEstimatesToggle(session Session) Session {
	session.ShowEstimates = !session.ShowEstimates
	return session
}
