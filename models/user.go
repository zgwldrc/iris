package models

import (
	"iris/services/mysql"
	"iris/services/bcrypt"
)

type User struct {
    Model
    Name string              `gorm:"not null;unique_index"  json:"name,omitempty"`
    Password string          `gorm:"not null"               json:"password,omitempty"`
    Account []Account        `json:"account,omitempty"`//has-many
}

//定义一个结构体，专门用于承载CheckUserValidity检查结果信息
type CheckUserResult struct {
	NotExist        bool           `json:"not_exist,omitempty"`
	PasswordInvalid bool    `json:"password_invalid,omitempty"`
	InputIsBroken   bool      `json:"input_is_broken,omitempty"`
	ID              int                 `json:"id,omitempty"`
}

//检查用户合法性，用于登录时，建立用户时，对用户输入的检测
//在返回结果中包含：账户是否存在，密码是否正确
func (this *User)CheckUserValidity() CheckUserResult {
    //临时存储mysql返回的结果
	tempUser := User{}
	result := CheckUserResult{}

	//CaseA:
	//空用户名 or 空密码情况
	if this.Name == "" || this.Password == "" {
		result.InputIsBroken = true
		return result
	}

	//CaseB:
	//用户输入合法则进入这里
	//根据数据库返回判断用户是否存在
	mysql.DB.Where("name = ?", this.Name).First(&tempUser)

	//CaseB-1
	//用户名不存在
	if  tempUser.Name == "" {
		result.NotExist = true
		return result
	}

	//CaseB-2
	//用户存在时进入这里
	//CaseB-2-1
	//密码不正确
	if !bcrypt.CompareHashAndPassword(tempUser.Password, this.Password) {
		result.PasswordInvalid = true
		return result
	}

	//CaseB-2-2
	//校验成功
	result.ID = tempUser.ID
	return result
}

func (this *User)Create() {
	this.Password = bcrypt.HashPassword(this.Password)
	mysql.DB.Create(this)
}

