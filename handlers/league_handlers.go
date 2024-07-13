package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premier_league_application/dtos"
	"premier_league_application/models"
	"premier_league_application/services"
)

// GetLeagues godoc
// @Summary Get all leagues
// @Description Get a list of all leagues
// @Tags leagues
// @Accept  json
// @Produce  json
// @Success 200 {array} models.League
// @Failure 500 {object} dtos.ErrorResponse
// @Router /leagues/ [get]
func GetLeagues(c *gin.Context) {
	leagues, err := services.GetAllLeagues()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, leagues)
}

// CreateLeague godoc
// @Summary Create a new league
// @Description Create a new league
// @Tags leagues
// @Accept  json
// @Produce  json
// @Param league body models.League true "League to create"
// @Success 200 {object} models.League
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /leagues/add [post]
func CreateLeague(c *gin.Context) {

	var league models.League

	if err := c.ShouldBindJSON(&league); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Error: err.Error()})
		return
	}

	if err := services.CreateLeague(league); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, league)
}

// GetLeagueTable godoc
// @Summary Get league table
// @Description Get the current league table
// @Tags leagues
// @Accept  json
// @Produce  json
// @Success 200 {array} dtos.TeamPosition
// @Router /leagues/league-table [get]
func GetLeagueTable(c *gin.Context) {
	positions := services.GetLeagueTable()

	c.JSON(http.StatusOK, positions)
}

// PlayOneWeek godoc
// @Summary Play one week of matches
// @Description Simulate one week of matches and update the league
// @Tags leagues
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.WeeklyResult
// @Router /leagues/play-one-week [get]
func PlayOneWeek(c *gin.Context) {
	weekly_result, _ := services.PlayOneWeek()

	c.JSON(http.StatusOK, weekly_result)
}

// PlayAllWeek godoc
// @Summary Play all remaining weeks
// @Description Simulate all remaining weeks and update the league
// @Tags leagues
// @Accept  json
// @Produce  json
// @Success 200 {array} dtos.TeamPosition
// @Router /leagues/play-all-week [get]
func PlayAllWeek(c *gin.Context) {
	positions := services.PlayAllWeek()

	c.JSON(http.StatusOK, positions)
}

// Prediction godoc
// @Summary Get championship predictions
// @Description Get the current championship predictions for each team
// @Tags leagues
// @Accept  json
// @Produce  json
// @Success 200 {array} dtos.TeamPrediction
// @Router /leagues/prediction [get]
func Prediction(c *gin.Context) {
	predictions := services.Prediction()

	c.JSON(http.StatusOK, predictions)
}

// Reset godoc
// @Summary Reset the league
// @Description Reset all league data
// @Tags leagues
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ErrorResponse
// @Router /leagues/reset [get]
func Reset(c *gin.Context) {
	services.ResetAllLeague()
}
