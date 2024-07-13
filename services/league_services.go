package services

import (
	"fmt"
	"math/rand"
	"premier_league_application/dtos"
	"premier_league_application/models"
	"premier_league_application/repositories"
	"sort"

	"github.com/mitchellh/mapstructure"
)

func CreateLeague(league models.League) error {
	return repositories.CreateLeague(league)
}

func GetAllLeagues() ([]dtos.League, error) {
	leagues, _ := repositories.GetAllLeagues()

	var leagueDtoList []dtos.League

	for _, src := range leagues {
		var leagueDto dtos.League
		err := mapstructure.Decode(src, &leagueDto)

		if err != nil {
			return nil, err
		}
		leagueDtoList = append(leagueDtoList, leagueDto)
	}

	return leagueDtoList, nil
}

func GetLeagueTable() []dtos.TeamPosition {
	var teams []models.Team
	teams, _ = repositories.GetTeamsByLeagueId(1)

	sorting(teams)

	return mapping(teams)
}

func sorting(teams []models.Team) {
	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Point > teams[j].Point
	})
}

func mapping(teams []models.Team) []dtos.TeamPosition {
	teamPositionList := make([]dtos.TeamPosition, len(teams))

	for i, team := range teams {
		teamPositionList[i] = dtos.TeamPosition{
			Position: i + 1,
			Club:     team.Name,
			Drawn:    team.Drawn,
			Lost:     team.Lost,
			Points:   team.Point,
			Won:      team.Won,
		}
	}

	return teamPositionList
}

func PlayOneWeek() (dtos.WeeklyResult, error) {
	var teams []models.Team
	teams, _ = repositories.GetTeamsByLeagueId(1)

	return Play(teams)
}

func Play(teams []models.Team) (dtos.WeeklyResult, error) {

	league, _ := repositories.GetLeagueById(1)

	if league.Week <= len(teams)*2-2 {

		shuffleTeams(teams)

		matches := make([]dtos.Match, len(teams)/2)
		for i := 0; i < len(teams)/2; i++ {
			hostTeam := teams[2*i]
			awayTeam := teams[2*i+1]
			hostTeamScore, awayTeamScore := calculateScore(hostTeam, awayTeam)

			if hostTeamScore > awayTeamScore {
				hostTeam.Won++
				hostTeam.Point += 3
				awayTeam.Lost++
			} else if hostTeamScore < awayTeamScore {
				awayTeam.Won++
				awayTeam.Point += 3
				hostTeam.Lost++
			} else {
				hostTeam.Drawn++
				awayTeam.Drawn++
				hostTeam.Point++
				awayTeam.Point++
			}

			matches[i] = dtos.Match{
				HostTeam:      hostTeam.Name,
				AwayTeam:      awayTeam.Name,
				HostTeamScore: hostTeamScore,
				AwayTeamScore: awayTeamScore,
			}

			if err := repositories.UpdateTeam(hostTeam); err != nil {
				return dtos.WeeklyResult{}, err
			}

			if err := repositories.UpdateTeam(awayTeam); err != nil {
				return dtos.WeeklyResult{}, err
			}
		}

		league.Week++

		weeklyResult := dtos.WeeklyResult{
			Week:    league.Week,
			Matches: matches,
		}

		repositories.UpdateLeague(league)

		return weeklyResult, nil
	} else {

		weeklyResult := dtos.WeeklyResult{
			Week: league.Week,
		}

		league.Week = 0
		repositories.UpdateLeague(league)
		return weeklyResult, fmt.Errorf("no more available weeks")
	}
}

func calculateScore(HostTeam, AwayTeam models.Team) (int, int) {
	HostTeamScore := float32(HostTeam.Title)*0.05 + HostTeam.ValueOfPlayers*0.05 + float32(HostTeam.Won)*0.25 + float32(HostTeam.Drawn)*0.10 - float32(HostTeam.Lost)*0.20
	AwayTeamScore := float32(AwayTeam.Title)*0.05 + AwayTeam.ValueOfPlayers*0.05 + float32(AwayTeam.Won)*0.25 + float32(AwayTeam.Drawn)*0.10 - float32(AwayTeam.Lost)*0.20
	return int(HostTeamScore), int(AwayTeamScore)
}

func shuffleTeams(teams []models.Team) {
	for i := len(teams) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		teams[i], teams[j] = teams[j], teams[i]
	}
}

func PlayAllWeek() []dtos.TeamPosition {
	var teams []models.Team
	teams, _ = repositories.GetTeamsByLeagueId(1)

	league, _ := repositories.GetLeagueById(1)
	totalWeeks := len(teams)*2 - 2

	for league.Week <= totalWeeks {
		_, err := Play(teams)
		if err != nil {
			break
		}
		league, _ = repositories.GetLeagueById(1)
		teams, _ = repositories.GetTeamsByLeagueId(1)
	}

	return GetLeagueTable()
}

func Prediction() []dtos.TeamPrediction {

	var teams []models.Team
	teams, _ = repositories.GetTeamsByLeagueId(1)
	league, _ := repositories.GetLeagueById(1)

	sorting(teams)

	totalWeeks := len(teams)*2 - 2
	remainingWeeks := totalWeeks - league.Week
	predictions := make([]dtos.TeamPrediction, len(teams))

	for i, team := range teams {

		maxPossiblePoints := team.Point + (remainingWeeks * 3)

		odds := (float64(team.Point) + 1) / (float64(maxPossiblePoints)*100 + float64(team.Title)*1)

		predictions[i] = dtos.TeamPrediction{
			TeamName: team.Name,
			Odds:     odds,
		}
	}

	var totalOdds float64
	for _, prediction := range predictions {
		totalOdds += prediction.Odds
	}

	for i := range predictions {
		predictions[i].Odds = (predictions[i].Odds / totalOdds) * 100
	}

	return predictions
}

func ResetAllLeague() {
	var teams []models.Team
	teams, _ = repositories.GetTeamsByLeagueId(1)
	league, _ := repositories.GetLeagueById(1)

	league.Week = 0

	for _, team := range teams {
		team.Lost = 0
		team.Point = 0
		team.Drawn = 0
		team.Won = 0
		repositories.UpdateTeam(team)
	}
	repositories.UpdateLeague(league)
}
