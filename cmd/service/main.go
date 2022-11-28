package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"scrum-poker/internal/app"
	"scrum-poker/internal/http"
	"scrum-poker/internal/poker/pokersession"
)

func main() {
	appObj := app.NewApp(pokersession.NewMemPokerSession())

	ginEngine := gin.Default()
	ginEngine = http.InitRoutes(ginEngine, appObj)
	err := ginEngine.Run()
	fmt.Println(fmt.Errorf("error: %v", err))
}
