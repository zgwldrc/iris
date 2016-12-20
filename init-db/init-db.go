package init_db

import (
	"iris/models"
	"iris/services/bcrypt"
    "github.com/jinzhu/gorm"
)

func Init(DB *gorm.DB) {
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
    DB.Model(&models.Account{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
    DB.Model(&models.Account{}).AddForeignKey("endpoint_id", "endpoints(id)", "SET NULL", "CASCADE")
    DB.Model(&models.Account{}).AddForeignKey("account_type_id", "account_types(id)", "SET NULL", "CASCADE")

    //DB.Model(models.User{}).Related(models.Account{})
    //DB.Model(models.Account{}).Related(models.Endpoint{})
    //DB.Model(models.Account{}).Related(models.AccountType{})


    accountTypeList := []models.AccountType{
        {Type: "Mysql"},
        {Type: "Web"},
        {Type: "SSH"},
        {Type: "VSFTP"},
		{Type: "SVN"},
    }

    for _,v := range accountTypeList {
        DB.Create(&v)
    }

    admin := models.User{
        Name: "admin",
        Password: bcrypt.HashPassword("admin"),
    }

    DB.Create(&admin)
}

