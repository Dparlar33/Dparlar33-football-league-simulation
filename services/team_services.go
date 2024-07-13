package services

import (
	"premier_league_application/dtos"
	"premier_league_application/models"
	"premier_league_application/repositories"

	"github.com/mitchellh/mapstructure"
)

func CreateTeam(team dtos.Team) error {
	var newTeam models.Team

	mapstructure.Decode(team, &newTeam)

	return repositories.CreateTeam(newTeam)
}

func GetAllTeams() ([]dtos.Team, error) {
	teams, _ := repositories.GetAllTeams()

	var teamDtoList []dtos.Team

	for _, src := range teams {
		var teamDto dtos.Team
		err := mapstructure.Decode(src, &teamDto)

		if err != nil {
			return nil, err
		}
		teamDtoList = append(teamDtoList, teamDto)
	}

	return teamDtoList, nil
}

func GetTeamById(id int) (models.Team, error) {
	return repositories.GetTeamById(id)
}
