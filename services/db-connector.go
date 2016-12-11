package modules

import (
	"fmt"

	_"github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jinzhu/gorm"
    "gopkg.in/kataras/iris.v4"

	"iris/models"
)

//这个模块提供一个DB对象的指针，在init()初始化，在main函数中释放
var DB *gorm.DB

func init(){
    var err error
    DB, err = gorm.Open("mysql", iris.Config.Other["MySQLDSN"])
    if err != nil {
        fmt.Println(err)
    }
}

func Initdb(DB *gorm.DB) {
    DB.DropTableIfExists(
		& models.Account{},
	).DropTableIfExists(
		& models.Endpoint{},
		& models.AccountType{},
		& models.User{},
	)
	
    DB.CreateTable(
		& models.Endpoint{},
		& models.AccountType{},
		& models.User{},
		& models.Account{},
	)

    // Add foreign key
    // 1st param : foreignkey field
    // 2nd param : destination table(id)
    // 3rd param : ONDELETE
    // 4th param : ONUPDATE
    DB.Model(&models.Account{}).AddForeignKey("user_id", "users(id)", "SET NULL", "SET NULL")
    DB.Model(&models.Account{}).AddForeignKey("endpoint_id", "endpoints(id)", "SET NULL", "SET NULL")
    DB.Model(&models.Account{}).AddForeignKey("account_type_id", "account_types(id)", "SET NULL", "SET NULL")

    DB.Model(models.User{}).Related(models.Account{})
    DB.Model(models.Account{}).Related(models.Endpoint{})
    DB.Model(models.Account{}).Related(models.AccountType{})


    accountTypeList := []models.AccountType{
        {Type: "Mysql"},
        {Type: "Web"},
        {Type: "SSH"},
        {Type: "VSFTP"},
    }


    for _,v := range accountTypeList {
        DB.Create(&v)
    }

    admin := models.User{
        Name: "admin",
        Password: HashPassword("admin"),
    }

    DB.Create(&admin)
}

