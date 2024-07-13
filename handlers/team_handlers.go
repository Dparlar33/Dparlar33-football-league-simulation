package handlers

import (
	"net/http"
	"premier_league_application/dtos"
	"premier_league_application/services"

	"github.com/gin-gonic/gin"
)

// GetAllTeams godoc
// @Summary Get all teams
// @Description Get a list of all teams
// @Tags teams
// @Accept  json
// @Produce  json
// @Success 200 {array} dtos.Team
// @Failure 500 {object} dtos.ErrorResponse
// @Router /teams/ [get]
func GetAllTeams(c *gin.Context) {
	teams, err := services.GetAllTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teams)
}

// CreateTeam godoc
// @Summary Create a new team
// @Description Create a new team
// @Tags teams
// @Accept  json
// @Produce  json
// @Param team body dtos.Team true "Team to create"
// @Success 200 {object} dtos.Team
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /teams/add [post]
func CreateTeam(c *gin.Context) {

	var team dtos.Team

	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateTeam(team); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}
