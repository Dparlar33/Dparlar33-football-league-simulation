package models

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Value    int32  `json:"value"`
	Position string `json:"position"`
	TeamID   int32  `json:"team_id"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
