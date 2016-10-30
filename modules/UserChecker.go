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

	//过滤掉空用户名情况
	if this.InputUser.Name == "" {
		return this.result
	}

	//检测密码字段是否为空
	if this.InputUser.Password != "" {
		//用户输入合法则进入这里
		//设置相应的flag
		this.result.InputValid = true
		//根据用户名查询数据库
		DB.Where("name = ?", this.InputUser.Name).First(this.tempUser)
		//根据数据库返回判断用户是否存在
		if  this.tempUser.Name != "" {
			//用户存在时进入这里
			this.result.Exist = true

			//检查输入用户是否为admin
			if this.tempUser.ID == 1 {
				this.result.IsAdmin =  true
			}

			//检查密码有效性
			if CompareHashAndPassword(this.tempUser.Password, this.InputUser.Password) {
				this.result.PasswordValid = true
			}
		}
		//用户不存在则跳到这里
		return this.result
	}
	//输入密码为空则跳到这里
	return this.result
}

