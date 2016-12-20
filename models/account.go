package models

import (
    "iris/services/mysql"
    "gopkg.in/kataras/iris.v4"
)

type Account struct {
    Model

    UserID int               `json:"-" gorm:"unique_index:uix_user_id_account_account_type_id_endpoint_id"`
    Account string           `json:"account" gorm:"unique_index:uix_user_id_account_account_type_id_endpoint_id;not null"`
    Password string          `json:"password" gorm:"not null"`

	AccountType AccountType  `json:"account_type"`
    AccountTypeID int        `json:"account_type_id" gorm:"unique_index:uix_user_id_account_account_type_id_endpoint_id"`

    Endpoint Endpoint        `json:"endpoint"`
    EndpointID int           `json:"endpoint_id" gorm:"unique_index:uix_user_id_account_account_type_id_endpoint_id"`

    DescInfo string          `json:"desc_info" gorm:"not null"`

}

type CheckAccountResult struct {
    //帐户重复
	IsDup        bool           `json:"is_dup,omitempty"`
    //输入信息不完整
	InputIsBroken   bool        `json:"input_is_broken,omitempty"`
}

func (a *Account)Load(uid, aid int){
    mysql.DB.
        Joins("JOIN account_types ON account_types.id = accounts.account_type_id").
        Joins("JOIN endpoints ON endpoints.id = accounts.endpoint_id").
        Where("accounts.user_id = ? AND accounts.id = ?", uid, aid).Find(a)
}
func (a *Account)Create() {
    r := a.Check()
    switch  {
    case r.IsDup:
        iris.Logger.Println("Input account failed due to detecting duplicated record")
        return
    case r.InputIsBroken:
        iris.Logger.Println("Input account failed due to broken input")
        return
    }
    mysql.DB.Create(a)
}

func (a *Account)Check() (r CheckAccountResult) {
    var ta Account

    mysql.DB.Where(
        "user_id = ? AND account = ? AND account_type_id = ? AND endpoint_id = ?",
        a.UserID,a.Account,a.AccountTypeID,a.EndpointID,
    ).First(&ta)

    if ta.ID !=0 {
        r.IsDup = true
       return r
    }
    return r
}

func (a *Account)Delete(){
    if a.ID == 0 {
        return
    }
    mysql.DB.Delete(a)
}

func (a *Account)Update(){
    mysql.DB.Save(a)
}



