package handlers

import (
	"net/http"
	"premier_league_application/dtos"
	"premier_league_application/services"

	"github.com/gin-gonic/gin"
)

// GetPlayers godoc
// @Summary Get all players
// @Description Get a list of all players
// @Tags players
// @Accept  json
// @Produce  json
// @Success 200 {array} dtos.Player
// @Failure 500 {object} dtos.ErrorResponse
// @Router /players/ [get]
func GetPlayers(c *gin.Context) {
	players, err := services.GetAllPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, players)
}

// CreatePlayer godoc
// @Summary Create a new player
// @Description Create a new player
// @Tags players
// @Accept  json
// @Produce  json
// @Param player body dtos.Player true "Player to create"
// @Success 200 {object} dtos.Player
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /players/add [post]
func CreatePlayer(c *gin.Context) {

	var player dtos.Player

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreatePlayer(player); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}
