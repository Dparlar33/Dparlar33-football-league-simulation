package models

import (
	"time"

	"gorm.io/gorm"
)

type League struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Team        []Team `json:"teams"`
	Name        string `json:"names"`
	LimitOfTeam int    `json:"limit_of_teams"`
	Week        int    `json:"weeks"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
