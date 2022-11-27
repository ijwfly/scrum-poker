package room

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

func NewRoom(roomId int64) Room {
	var room Room
	room.RoomId = roomId
	room.EstimateOptions = defaultEstimateOptions()
	return room
}
