package models

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Name string              `gorm:"not null;unique_index"`
    Password string          `gorm:"not null"`
    Account []Account        //has-many
}

type Account struct {
    gorm.Model

    UserID uint

    Endpoint Endpoint        //belong-to
    EndpointID uint

    AccountType AccountType  //belong-to
    AccountTypeID uint

    Account string           `gorm:"not null"`
    Password string          `gorm:"not null"`
    DescInfo string          `gorm:"not null"`
}

type AccountType struct {
    gorm.Model
    Type string              `gorm:"not null;unique_index"`
}

type Endpoint struct {
    gorm.Model
    Endpoint string          `gorm:"not null;unique_index"`
}
