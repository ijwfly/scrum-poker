package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"scrum-poker/internal/app/http"
)

func main() {
	ginEngine := gin.Default()
	ginEngine = http.InitRoutes(ginEngine)
	err := ginEngine.Run()
	fmt.Println(fmt.Errorf("error: %v", err))
}
