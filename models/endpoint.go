package models


type Endpoint struct {
	Model

	URL string    `json:"url"     gorm:"unique_index;not null"`
}
