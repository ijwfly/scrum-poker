package poker

import "math/rand"

type EstimateOption string
type OnlineState string

const (
	Online  OnlineState = "online"
	Offline OnlineState = "offline"
)

const TOKEN_LENGTH = 32

type User struct {
	Name  string `json:"name"`
	Token string `json:"token,omitempty"`
}

func GenerateToken() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345678")
	b := make([]rune, TOKEN_LENGTH)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewUser(name string) User {
	var user User
	user.Name = name
	user.Token = GenerateToken()
	return user
}

type UserState struct {
	User
	Estimate *EstimateOption `json:"estimate"`
	Presence OnlineState     `json:"presence"`
}

type Session struct {
	SessionId              string               `json:"session_id"`
	SessionEstimateOptions []EstimateOption     `json:"session_estimate_options"`
	ShowEstimates          bool                 `json:"show_estimates"`
	Users                  map[string]UserState `json:"users"`
}

func defaultEstimateOptions() []EstimateOption {
	return []EstimateOption{
		"?",
		"0.5",
		"1",
		"2",
		"3",
		"5",
		"8",
		"13",
	}
}

func NewSession(sessionId string) Session {
	var session Session
	session.SessionId = sessionId
	session.SessionEstimateOptions = defaultEstimateOptions()
	session.Users = make(map[string]UserState)
	session.ShowEstimates = false
	return session
}
