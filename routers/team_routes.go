package routers

import (
	"premier_league_application/handlers"

	"github.com/gin-gonic/gin"
)

func TeamRouters(router *gin.RouterGroup) {
	router.GET("/", handlers.GetAllTeams)
	router.POST("/add", handlers.CreateTeam)
}
