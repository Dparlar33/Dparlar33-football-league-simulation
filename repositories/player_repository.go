package repositories

import (
	"premier_league_application/config"
	"premier_league_application/models"
)

func CreatePlayer(player models.Player) error {
	return config.DB.Create(&player).Error
}

func GetAllPlayers() ([]models.Player, error) {
	var players []models.Player
	err := config.DB.Find(&players).Error
	return players, err
}
