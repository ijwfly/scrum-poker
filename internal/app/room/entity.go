package room

type EstimateOption string

type Room struct {
	RoomId           int64
	EstimateOptions  []EstimateOption
	CurrentSessionId *int64
}
