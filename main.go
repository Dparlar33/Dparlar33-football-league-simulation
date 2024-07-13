package main

import (
	"premier_league_application/config"
	_ "premier_league_application/docs"

	"premier_league_application/models"
	"premier_league_application/routers"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title League Management API
// @version 1.0
// @description This is a sample server for managing leagues.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	config.DatabaseConnection()

	config.DB.AutoMigrate(&models.League{}, &models.Team{}, &models.Player{})

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	playerRouter := router.Group("/players")
	routers.PlayerRouters(playerRouter)

	teamRouter := router.Group("/teams")
	routers.TeamRouters(teamRouter)

	leagueRouter := router.Group("/leagues")
	routers.LeagueRouters(leagueRouter)

	router.Run(":8080")
}
