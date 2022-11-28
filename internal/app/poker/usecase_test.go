package poker

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSession_Scenario(t *testing.T) {
	// Arrange
	user1 := NewUser("user_1")
	user2 := NewUser("user_2")
	user3 := NewUser("user_3")

	session := NewSession("unbeleivable-monkey")

	// Act & Assert
	session, _ = UserJoinSession(session, user1)
	session, _ = UserJoinSession(session, user2)
	session, _ = UserJoinSession(session, user3)
	require.Equal(t, 3, len(session.Users))

	session, _ = UserSetEstimate(session, user1, "1")
	session, _ = UserSetEstimate(session, user2, "2")
	session, _ = UserSetEstimate(session, user3, "18")
	require.Equal(t, EstimateOption("1"), *session.Users[user1.Token].Estimate)
	require.Equal(t, EstimateOption("2"), *session.Users[user2.Token].Estimate)
	require.Equal(t, (*EstimateOption)(nil), session.Users[user3.Token].Estimate)

	session = UserLeaveSession(session, user1)
	require.Equal(t, Offline, session.Users[user1.Token].Presence)

	session, _ = UserJoinSession(session, user1)
	require.Equal(t, Online, session.Users[user1.Token].Presence)

	session = ShowEstimatesToggle(session)
	require.Equal(t, true, session.ShowEstimates)

	session = ShowEstimatesToggle(session)
	require.Equal(t, false, session.ShowEstimates)

	session = UserLeaveSession(session, user1)
	session = ResetEstimates(session)
	require.Equal(t, 2, len(session.Users))
	require.Equal(t, (*EstimateOption)(nil), session.Users[user2.Token].Estimate)
	require.Equal(t, (*EstimateOption)(nil), session.Users[user3.Token].Estimate)
	require.Equal(t, false, session.ShowEstimates)
}
