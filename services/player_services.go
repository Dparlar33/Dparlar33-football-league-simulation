package services

import (
	"premier_league_application/dtos"
	"premier_league_application/models"
	"premier_league_application/repositories"

	"github.com/mitchellh/mapstructure"
)

func CreatePlayer(player dtos.Player) error {

	var newPlayer models.Player

	mapstructure.Decode(player, &newPlayer)

	err := repositories.CreatePlayer(newPlayer)
	if err != nil {
		return err
	}

	team, err := repositories.GetTeamById(int(newPlayer.TeamID))
	if err != nil {
		return err
	}

	team.CapacityOfPlayer++
	team.ValueOfPlayers += float32(newPlayer.Value)
	return repositories.UpdateTeam(team)
}

func GetAllPlayers() ([]dtos.Player, error) {
	players, _ := repositories.GetAllPlayers()

	var playerDtoList []dtos.Player

	for _, src := range players {
		var playerDto dtos.Player
		err := mapstructure.Decode(src, &playerDto)

		if err != nil {
			return nil, err
		}
		playerDtoList = append(playerDtoList, playerDto)
	}

	return playerDtoList, nil
}
