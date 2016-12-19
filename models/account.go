package models

import "iris/services/mysql"

type Account struct {
    Model

    UserID int               `gorm:"unique_index:uix_user_id_account_account_type_id_endpoint_id;"`
    Account string           `gorm:"unique_index:uix_user_id_account_account_type_id_endpoint_id;not null"`
    Password string          `gorm:"not null"`

	AccountType AccountType
    AccountTypeID int        `gorm:"unique_index:uix_user_id_account_account_type_id_endpoint_id;"`

    Endpoint Endpoint
    EndpointID int           `gorm:"unique_index:uix_user_id_account_account_type_id_endpoint_id;"`

    DescInfo string          `gorm:"not null"`

}

type CheckAccountResult struct {
    //帐户不存在
	NotExist        bool           `json:"not_exist,omitempty"`
    //输入信息不完整
	InputIsBroken   bool      `json:"input_is_broken,omitempty"`
    //account已存在时从数据库返回的id标志
	ID              int                 `json:"id,omitempty"`
}

func (a *Account)Create(){
    a.ID = 0
    mysql.DB.Create(a)
}
func (a *Account)Check() (r CheckAccountResult) {

    return r
}
func (a *Account)Delete(){
    if a.ID == 0 {
        return
    }
    mysql.DB.Delete(a)
}


