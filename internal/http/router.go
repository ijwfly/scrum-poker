package http

import (
	"github.com/gin-gonic/gin"
	"scrum-poker/internal/app"
	"scrum-poker/internal/http/handlers"
)

type StrangeStruct struct {
	number int
}

func (s *StrangeStruct) Handle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": s.number,
	})
}

func New(number int) *StrangeStruct {
	return &StrangeStruct{
		number: number,
	}
}

func dummy(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func InitRoutes(ginEngine *gin.Engine, app *app.App) *gin.Engine {
	api := ginEngine.Group("/api")
	api.GET("/get_token", handlers.GetTokenHandler)
	api.GET("/session/:sessionId", handlers.NewGetSessionHandler(app).Handle)
	return ginEngine
}
