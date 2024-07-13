package routers

import (
	"premier_league_application/handlers"

	"github.com/gin-gonic/gin"
)

func LeagueRouters(router *gin.RouterGroup) {
	router.GET("/", handlers.GetLeagues)
	router.POST("/add", handlers.CreateLeague)
	router.GET("/league-table", handlers.GetLeagueTable)
	router.GET("/play-one-week", handlers.PlayOneWeek)
	router.GET("/play-all-week", handlers.PlayAllWeek)
	router.GET("/prediction", handlers.Prediction)
	router.GET("/reset", handlers.Reset)
}
