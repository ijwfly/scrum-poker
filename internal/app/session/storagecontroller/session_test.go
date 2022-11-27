package storagecontroller

import (
	"github.com/stretchr/testify/require"
	"scrum-poker/internal/app/session"
	"testing"
)

func TestSessionStorageController_Scenario(t *testing.T) {
	sessionStorageController := NewSessionStorageController()
	sessionName := "some-session"
	user1 := session.NewUser("user1")
	user2 := session.NewUser("user2")
	user3 := session.NewUser("user3")

	_, _ = sessionStorageController.UserJoinSession(user1, sessionName)
	_, _ = sessionStorageController.UserJoinSession(user2, sessionName)
	sessionObj, _ := sessionStorageController.UserJoinSession(user3, sessionName)
	require.Equal(t, 3, len(sessionObj.Users))

	_, _ = sessionStorageController.UserSetEstimate(user1, sessionName, "1")
	_, _ = sessionStorageController.UserSetEstimate(user2, sessionName, "2")
	sessionObj, _ = sessionStorageController.UserSetEstimate(user3, sessionName, "18")
	require.Equal(t, session.EstimateOption("1"), *sessionObj.Users[user1.Token].Estimate)
	require.Equal(t, session.EstimateOption("2"), *sessionObj.Users[user2.Token].Estimate)
	require.Equal(t, (*session.EstimateOption)(nil), sessionObj.Users[user3.Token].Estimate)

	sessionObj, _ = sessionStorageController.UserLeaveSession(user1, sessionName)
	require.Equal(t, session.Offline, sessionObj.Users[user1.Token].Presence)

	sessionObj, _ = sessionStorageController.UserJoinSession(user1, sessionName)
	require.Equal(t, session.Online, sessionObj.Users[user1.Token].Presence)

	sessionObj, _ = sessionStorageController.SessionShowEstimatesToggle(sessionName)
	require.Equal(t, true, sessionObj.ShowEstimates)

	sessionObj, _ = sessionStorageController.SessionShowEstimatesToggle(sessionName)
	require.Equal(t, false, sessionObj.ShowEstimates)

	sessionObj, _ = sessionStorageController.UserLeaveSession(user1, sessionName)
	sessionObj, _ = sessionStorageController.SessionResetEstimates(sessionName)
	require.Equal(t, 2, len(sessionObj.Users))
	require.Equal(t, (*session.EstimateOption)(nil), sessionObj.Users[user2.Token].Estimate)
	require.Equal(t, (*session.EstimateOption)(nil), sessionObj.Users[user3.Token].Estimate)
	require.Equal(t, false, sessionObj.ShowEstimates)
}
