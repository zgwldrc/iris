package models

import (
	"iris/modules"
)

type User struct {
    Model
    Name string              `gorm:"not null;unique_index"  json:"name"`
    Password string          `gorm:"not null"               json:"password"`
    Account []Account        //has-many
}

//定义一个结构体，专门用于承载CheckUserValidity检查结果信息
type CheckUserResult struct {
	Exist bool              `json:"exist"`
	PasswordValid bool      `json:"password_valid"`
	Id int                  `json:"id"`
}

//检查用户合法性，用于登录时，建立用户时，对用户输入的检测
//在返回结果中包含：账户是否存在，密码是否正确
func (this *User)CheckUserValidity() CheckUserResult {
    //临时存储mysql返回的结果
	tempUser := User{}
	result := CheckUserResult{}
	//空用户名 or 空密码情况
	if this.Name == "" || this.Password == "" {
		return CheckUserResult{}
	}

	//用户输入合法则进入这里
	//根据用户名查询数据库
	modules.DB.Where("name = ?", this.Name).First(tempUser)
	//根据数据库返回判断用户是否存在
	if  tempUser.Name != "" {
		//用户存在时进入这里
		result.Id = tempUser.ID
		result.Exist = true
		//检查密码有效性
		if modules.CompareHashAndPassword(tempUser.Password, this.Password) {
			result.PasswordValid = true
		}
	}
	return result
}

