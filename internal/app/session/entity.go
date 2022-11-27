package session

type OnlineState string

const (
	Online  OnlineState = "online"
	Offline OnlineState = "offline"
)

type UserState struct {
	UserName string
	Estimate *EstimateOption
	Presence OnlineState
}

type Session struct {
	SessionId              string
	SessionEstimateOptions []EstimateOption
	ShowEstimates          bool
	Users                  map[int64]UserState
}
