package pokersession

import (
	"github.com/stretchr/testify/require"
	"scrum-poker/internal/poker"
	"testing"
)

func TestSessionStorageController_Scenario(t *testing.T) {
	sessionStorageController := NewMemPokerSession()
	sessionName := "some-session"
	user1 := poker.NewUser("user1")
	user2 := poker.NewUser("user2")
	user3 := poker.NewUser("user3")

	_, _ = sessionStorageController.UserJoinSession(user1, sessionName)
	_, _ = sessionStorageController.UserJoinSession(user2, sessionName)
	sessionObj, _ := sessionStorageController.UserJoinSession(user3, sessionName)
	require.Equal(t, 3, len(sessionObj.Users))

	_, _ = sessionStorageController.UserSetEstimate(user1, sessionName, "1")
	_, _ = sessionStorageController.UserSetEstimate(user2, sessionName, "2")
	sessionObj, _ = sessionStorageController.UserSetEstimate(user3, sessionName, "18")
	require.Equal(t, poker.EstimateOption("1"), *sessionObj.Users[user1.Token].Estimate)
	require.Equal(t, poker.EstimateOption("2"), *sessionObj.Users[user2.Token].Estimate)
	require.Equal(t, (*poker.EstimateOption)(nil), sessionObj.Users[user3.Token].Estimate)

	sessionObj, _ = sessionStorageController.UserLeaveSession(user1, sessionName)
	require.Equal(t, poker.Offline, sessionObj.Users[user1.Token].Presence)

	sessionObj, _ = sessionStorageController.UserJoinSession(user1, sessionName)
	require.Equal(t, poker.Online, sessionObj.Users[user1.Token].Presence)

	sessionObj, _ = sessionStorageController.ShowEstimatesToggle(sessionName)
	require.Equal(t, true, sessionObj.ShowEstimates)

	sessionObj, _ = sessionStorageController.ShowEstimatesToggle(sessionName)
	require.Equal(t, false, sessionObj.ShowEstimates)

	sessionObj, _ = sessionStorageController.UserLeaveSession(user1, sessionName)
	sessionObj, _ = sessionStorageController.ResetEstimates(sessionName)
	require.Equal(t, 2, len(sessionObj.Users))
	require.Equal(t, (*poker.EstimateOption)(nil), sessionObj.Users[user2.Token].Estimate)
	require.Equal(t, (*poker.EstimateOption)(nil), sessionObj.Users[user3.Token].Estimate)
	require.Equal(t, false, sessionObj.ShowEstimates)
}
