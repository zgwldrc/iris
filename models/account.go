package models

type Account struct {
    Model

    UserID uint
    Account string           `gorm:"not null"`
    Password string          `gorm:"not null"`
	Type string              `gorm:"not null"`
    DescInfo string          `gorm:"not null"`
    Endpoint string          `gorm:"not null"`
}
