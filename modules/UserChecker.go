package modules

import (
	"iris/models"
)

//定义一个结构体，专门用于承载CheckUserValidity检查结果信息
type CheckUserResult struct {
	Exist bool
	PasswordValid bool
	InputValid bool
	IsAdmin bool
}

type UserChecker struct {
	InputUser *models.User
	tempUser *models.User
	result *CheckUserResult
}

func NewUserChecker (inputUser *models.User) (*UserChecker) {
	return &UserChecker{
		InputUser: inputUser,
		tempUser: &models.User{},
		result: &CheckUserResult{},
	}
}

//检查用户合法性，用于登录时，建立用户时，对用户输入的检测
//在返回结果中包含：账户是否存在，密码是否正确，是否为admin
func (this *UserChecker) CheckUserValidity() (*CheckUserResult) {

	//若用户输入的账户名称存在空字符串则直接返回预定结果，不再继续判断
	if this.InputUser.Name == "" {
		return this.result
	}

	//输入用户名不为空时，继续检测密码字段是否为空，来设置InputValid属性
	if this.InputUser.Password != "" {
		this.result.InputValid = true
	}

	//输入用户名不为空时，检查数据库中用户存在性
	DB.Where("name = ?", this.InputUser.Name).First(this.tempUser)
	//用户存在时
	if  this.tempUser.Name != "" {
		//用户存在时进入这里
		this.result.Exist = true

		//检查输入用户是否为admin
		if this.tempUser.ID == 1 {
			this.result.IsAdmin =  true
		}

		//检查密码字段是否包含值（检查InputValid属性）,不为空时进一步检查密码有效性
		if this.result.InputValid {
			if CompareHashAndPassword(this.tempUser.Password, this.InputUser.Password) {
			    this.result.PasswordValid = true
		    }
		}
	}
	return this.result
}

