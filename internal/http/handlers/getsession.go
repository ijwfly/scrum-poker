package handlers

import (
	"github.com/gin-gonic/gin"
	"scrum-poker/internal/app"
	"scrum-poker/internal/poker"
)

type GetSesionHandler struct {
	App *app.App
}

func NewGetSessionHandler(app *app.App) *GetSesionHandler {
	return &GetSesionHandler{
		App: app,
	}
}

func (h *GetSesionHandler) Handle(c *gin.Context) {
	sessionId := c.Param("sessionId")
	username := c.DefaultQuery("username", "")
	token := c.DefaultQuery("token", "")
	if len(username) != 0 && len(token) != 0 {
		user := poker.User{Name: username, Token: token}
		session, err := h.App.Poker.UserJoinSession(user, sessionId)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"session": session,
		})
		return
	} else {
		session := h.App.Poker.GetOrCreateSession(sessionId)
		c.JSON(200, gin.H{
			"session": session,
		})
	}
}
