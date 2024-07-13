package repositories

import (
	"premier_league_application/config"
	"premier_league_application/models"
)

func CreateTeam(team models.Team) error {
	return config.DB.Create(&team).Error
}

func GetAllTeams() ([]models.Team, error) {
	var teams []models.Team
	err := config.DB.Find(&teams).Error
	return teams, err
}

func GetTeamsByLeagueId(leagueId int) ([]models.Team, error) {

	var teams []models.Team

	err := config.DB.Where("league_id = ?", leagueId).Find(&teams).Error

	if err != nil {
		return nil, err
	}

	return teams, nil
}

func GetTeamById(id int) (models.Team, error) {
	var team models.Team

	err := config.DB.Where("id = ?", id).First(&team).Error
	if err != nil {
		return models.Team{}, err
	}

	return team, nil
}

func UpdateTeam(team models.Team) error {
	return config.DB.Save(&team).Error
}
