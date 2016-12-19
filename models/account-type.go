package models

type AccountType struct {
	Model
	Type string    `json:"type"        gorm:"not null;unique_index"`
}
