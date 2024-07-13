package routers

import (
	"premier_league_application/handlers"

	"github.com/gin-gonic/gin"
)

func PlayerRouters(router *gin.RouterGroup) {
	router.GET("/", handlers.GetPlayers)
	router.POST("/add", handlers.CreatePlayer)
}
