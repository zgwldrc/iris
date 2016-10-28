package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "fmt"
)

var DB *gorm.DB

type User struct {
    gorm.Model
    Name string
    Password string
    Account []Account        //has-many
}

type Account struct {
    gorm.Model

    UserID uint

    Endpoint Endpoint        //belong-to
    EndpointID uint

    AccountType AccountType  //belong-to
    AccountTypeID uint

    Account string
    Password string
    DescInfo string
}

type AccountType struct {
    gorm.Model
    Type string
}

type Endpoint struct {
    gorm.Model
    Endpoint string
}

func init(){
    var err error
    //DB, err = gorm.Open("mysql", iris.Config.Other["MysqlDSN"])
    DB, err = gorm.Open("mysql", "root:test123@tcp(192.168.0.62:3306)/app?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        fmt.Println(err)
    }
}


func Initdb(DB *gorm.DB) {
    
    DB.DropTableIfExists(&Account{}).DropTableIfExists(&Endpoint{}, &AccountType{},&User{})
    DB.CreateTable(&Endpoint{}, &AccountType{}, &User{}, &Account{})

    // Add foreign key
    // 1st param : foreignkey field
    // 2nd param : destination table(id)
    // 3rd param : ONDELETE
    // 4th param : ONUPDATE
    DB.Model(&Account{}).AddForeignKey("user_id", "users(id)", "SET NULL", "SET NULL")
    DB.Model(&Account{}).AddForeignKey("endpoint_id", "endpoints(id)", "SET NULL", "SET NULL")
    DB.Model(&Account{}).AddForeignKey("account_type_id", "account_types(id)", "SET NULL", "SET NULL")

    DB.Model(User{}).Related(Account{})
    DB.Model(Account{}).Related(Endpoint{})


    accountTypeList := []AccountType{
        {Type: "Mysql"},
        {Type: "Web"},
        {Type: "SSH"},
        {Type: "VSFTP"},
    }


    for _,v := range accountTypeList {
        DB.Create(&v)
    }

    admin := User{
        Name: "admin",
        Password: "admin",
    }

    DB.Create(&admin)
}
