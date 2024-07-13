package dtos

type Team struct {
	ID               uint
	LeagueID         uint
	Name             string
	ValueOfPlayers   float32
	Title            int
	Won              int
	Lost             int
	Drawn            int
	CapacityOfPlayer int
	Point            int
	Score            int
}
