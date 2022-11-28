package handlers

import (
	"github.com/gin-gonic/gin"
	"scrum-poker/internal/poker"
)

func GetTokenHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"token": poker.GenerateToken(),
	})
}
