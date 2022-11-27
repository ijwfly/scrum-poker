package session

import (
	room2 "scrum-poker/internal/app/room"
	"scrum-poker/internal/app/user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSession_Scenario(t *testing.T) {
	// Arrange
	user1 := user.NewUser(1, "user_1")
	user2 := user.NewUser(2, "user_2")
	user3 := user.NewUser(3, "user_3")

	roomObj := room2.NewRoom(1, "room_1")

	roomObj.EstimateOptions = []room2.EstimateOption{
		"?",
		"1",
		"2",
		"3",
		"5",
		"8",
		"13",
	}

	session := NewSession(1, roomObj)

	// Act & Assert
	session = UserJoinSession(session, user1)
	session = UserJoinSession(session, user2)
	session = UserJoinSession(session, user3)
	require.Equal(t, 3, len(session.Users))

	session = UserSetEstimate(session, user1, "1")
	session = UserSetEstimate(session, user2, "2")
	session = UserSetEstimate(session, user3, "18")
	require.Equal(t, room2.EstimateOption("1"), *session.Users[1].Estimate)
	require.Equal(t, room2.EstimateOption("2"), *session.Users[2].Estimate)
	require.Equal(t, (*room2.EstimateOption)(nil), session.Users[3].Estimate)

	session = UserLeaveSession(session, user1)
	require.Equal(t, Offline, session.Users[1].Presence)

	session = UserJoinSession(session, user1)
	require.Equal(t, Online, session.Users[1].Presence)

	session = SessionShowEstimatesToggle(session)
	require.Equal(t, true, session.ShowEstimates)

	session = SessionShowEstimatesToggle(session)
	require.Equal(t, false, session.ShowEstimates)

	session = SessionResetEstimates(session)
	require.Equal(t, (*room2.EstimateOption)(nil), session.Users[1].Estimate)
	require.Equal(t, (*room2.EstimateOption)(nil), session.Users[2].Estimate)
	require.Equal(t, (*room2.EstimateOption)(nil), session.Users[3].Estimate)
	require.Equal(t, false, session.ShowEstimates)
}
