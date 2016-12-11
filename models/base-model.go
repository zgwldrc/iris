package models

import "time"

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
