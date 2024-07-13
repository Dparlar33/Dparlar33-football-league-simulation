package models

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	ID               uint     `gorm:"primaryKey" json:"id"`
	Player           []Player `json:"players"`
	LeagueID         uint     `json:"league_id"`
	Name             string   `json:"names"`
	ValueOfPlayers   float32  `json:"value_of_players"`
	Title            int      `json:"titles"`
	Won              int      `json:"wins"`
	Lost             int      `json:"loses"`
	Drawn            int      `json:"deuces"`
	CapacityOfPlayer int      `json:"capacity_of_players"`
	Point            int      `json:"points"`
	Score            int      `json:"scores"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
