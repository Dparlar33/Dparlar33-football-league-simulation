package repositories

import (
	"premier_league_application/config"
	"premier_league_application/models"
)

func CreateLeague(league models.League) error {
	return config.DB.Create(&league).Error
}

func GetAllLeagues() ([]models.League, error) {
	var leagues []models.League
	err := config.DB.Find(&leagues).Error
	return leagues, err
}

func GetLeagueById(id int) (models.League, error) {
	var league models.League
	err := config.DB.Where("id = ?", id).First(&league).Error

	return league, err
}

func UpdateLeague(league models.League) error {
	return config.DB.Save(&league).Error
}
