package models

import "time"

type Model struct {
	ID        int         `json:"id,omitempty"`
	CreatedAt *time.Time  `json:"created_at,omitempty"  gorm:"index"`
	UpdatedAt *time.Time  `json:"updated_at,omitempty"`
}
