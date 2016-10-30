package api

import (
	"github.com/kataras/iris"
	"iris/models"
	"fmt"
	"iris/modules"
)

func init() {
	iris.Put("/User", func(ctx *iris.Context) {
		fmt.Println(
			"[Put /User]: \n",
			"操作用户:", ctx.Session().GetString("name"), "\n",
		)
		var inputUser = models.User{}
		if err := ctx.ReadJSON(&inputUser); err != nil {
			iris.Logger.Println(err)
			ctx.JSON(iris.StatusBadRequest, models.JSONResponse{
				Message:fmt.Sprint(err),
			})
		}
		checkResult := modules.NewUserChecker(&inputUser).CheckUserValidity()

		if checkResult.InputValid {
			//用户输入完整
			if checkResult.Exist {
				//用户存在
				ctx.JSON(iris.StatusBadRequest, models.JSONResponse{
					Message:"It seems like you input a existing user name  :-)",
				})
			} else {
				//用户不存在
				//仅当用户不存在且输入有效的情况下建立用户
				var tempUser models.User

				//首先hash用户明文密码再存储
				inputUser.Password = modules.HashPassword(inputUser.Password)
				modules.DB.Create(&inputUser)
				//建立完毕后再查询数据库，检测是否创建成功
				modules.DB.Where("name = ?", inputUser.Name).First(&tempUser)
				if tempUser.Name != "" {
					ctx.JSON(iris.StatusCreated, models.JSONResponse{
						Message:"Create User OK",
					})
				} else {
					ctx.JSON(iris.StatusInternalServerError, models.JSONResponse{
						Message:"Sorry ,Request has not been handled.",
					})
				}
			}
		} else {
			//用户输入不完整
			ctx.JSON(iris.StatusBadRequest, models.JSONResponse{
				Message:"It seems like you input a null string :-)",
			})
		}
	})

	//删除用户操作
	iris.Delete("/User", func(ctx *iris.Context) {
		var inputUser = models.User{}
		var tempUser = models.User{}
		if err:= ctx.ReadJSON(&inputUser) ;err != nil {
			iris.Logger.Println(err)
			ctx.JSON(iris.StatusBadRequest,models.JSONResponse{
				Message:fmt.Sprint(err),
			})
		}

		//当前回话用户是admin时才能执行该操作
		if ctx.Session().GetString("isAdmin") == "true" {
			//检查输入用户信息
			checkResult := modules.NewUserChecker(&inputUser).CheckUserValidity()
			//要求输入用户为非admin用户，并且输入用户必须存在
			if ! checkResult.IsAdmin && checkResult.Exist {
				//执行删除
				modules.DB.Where("name = ?",inputUser.Name).Delete(&inputUser)
				//检测删除结果
				modules.DB.Where("name = ?",inputUser.Name).First(&tempUser)
				if tempUser.Name == "" {
					//删除成功
					ctx.JSON(iris.StatusOK, models.JSONResponse{
						Message:"Delete Success",
					})
				} else {
					//删除失败
					ctx.JSON(iris.StatusInternalServerError, models.JSONResponse{
						Message:"Delete Failed",
					})
				}
			} else {
				//输入用户不满足上述条件
				ctx.JSON(iris.StatusForbidden, models.JSONResponse{
					Message:"Delete Failed",
				})
			}
		} else {
			//非Admin时
			ctx.JSON(iris.StatusForbidden, models.JSONResponse{
				Message:"Not Allow non admin do this operation",
			})
		}
	})
}
