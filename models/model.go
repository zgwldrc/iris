package models

import "time"

type Model struct {
	ID        int         `json:"id,omitempty"`
	CreatedAt *time.Time  `json:"-" gorm:"index"`
	UpdatedAt *time.Time  `json:"-"`
}
