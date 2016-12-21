package models

import (
    "iris/services/mysql"
    "gopkg.in/kataras/iris.v4"
)

type Account struct {
    Model

    UserID int               `json:"user_id,omitempty" gorm:"unique_index:uix_user_id_account_account_type_id_endpoint_id"`
    Account string           `json:"account" gorm:"unique_index:uix_user_id_account_account_type_id_endpoint_id;not null"`
    Password string          `json:"password" gorm:"not null"`

	AccountType AccountType  `json:"account_type"`
    AccountTypeID int        `json:"account_type_id,omitempty" gorm:"unique_index:uix_user_id_account_account_type_id_endpoint_id"`

    Endpoint Endpoint        `json:"endpoint"`
    EndpointID int           `json:"endpoint_id,omitempty" gorm:"unique_index:uix_user_id_account_account_type_id_endpoint_id"`

    DescInfo string          `json:"desc_info" gorm:"not null"`

}

type CheckAccountResult struct {
    //帐户重复
	IsDup        bool           `json:"is_dup,omitempty"`
    //输入信息不完整
	InputIsBroken   bool        `json:"input_is_broken,omitempty"`
}

func (a *Account)Load(uid, aid int){
    row:= mysql.DB.Raw(`
    SELECT a.id,a.account,a.password,a.desc_info,at.type,e.url
    FROM accounts as a
    INNER JOIN account_types as at ON a.account_type_id = at.id
    INNER JOIN endpoints     as e  ON a.endpoint_id = e.id
    WHERE a.user_id = ? AND a.id = ?
    `,uid,aid).Row()

    err := row.Scan(
        &a.ID,
        &a.Account,
        &a.Password,
        &a.DescInfo,
        &a.AccountType.Type,
        &a.Endpoint.URL,
    )
    if err != nil {
        iris.Logger.Println(err)
    }
}

func (a *Account)GetList(uid int, orderby string,limit int )

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



