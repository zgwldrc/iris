package models

import "time"

type Model struct {
	ID        int         `json:"id,omitempty"          gorm:"primary_key"`
	CreatedAt *time.Time  `json:"created_at,omitempty"  gorm:"index"`
	UpdatedAt *time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time  `json:"deleted_at,omitempty"`
}
