package http

import (
	"github.com/gin-gonic/gin"
)

func dummy(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func InitRoutes(ginEngine *gin.Engine) *gin.Engine {
	api := ginEngine.Group("/api")
	api.GET("/ping", dummy)
	return ginEngine
}
