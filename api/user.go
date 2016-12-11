// api for /user
package api

import (
	"fmt"
	"gopkg.in/kataras/iris.v4"
	"iris/models"
	"iris/modules"
)

func init() {
	//增
	iris.Post("/user", createUser)("createUser")
	//查
	iris.Get("/user/:id", getUserById)
	//改
	iris.Put("/user/:id", updateUser)
	//删
	iris.Delete("/user/:id", deleteUser)
} //end of func init()


func createUser(ctx *iris.Context) {
	fmt.Printf(
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
	//对输入用户模型数据进行有效性检测
	checkResult := modules.NewUserChecker(&inputUser).CheckUserValidity()

	if checkResult.InputValid {
		//用户输入完整
		if checkResult.Exist {
			//输入用户名已存在，不可用
			ctx.JSON(iris.StatusBadRequest, models.JSONResponse{
				Message:"用户已存在",
			})
		} else {
			//输入用户名可用


			//首先hash用户明文密码再存储
			inputUser.Password = modules.HashPassword(inputUser.Password)
			modules.DB.Create(&inputUser)
			//建立完毕后orm会再次查询数据库，将record数据加载到模型

			if inputUser.ID != 0 {
				ctx.JSON(iris.StatusCreated, models.JSONResponse{
					Message:"创建用户(" +inputUser.Name + ") 成功",
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
			Message:"输入完整性不合法",
		})
	}
}
func deleteUser(ctx *iris.Context) {
		var inputUser = models.User{}
		var tempUser = models.User{}
		if err:= ctx.ReadJSON(&inputUser) ;err != nil {
			iris.Logger.Println(err)
			ctx.JSON(iris.StatusBadRequest,models.JSONResponse{
				Message:fmt.Sprint(err),
			})
		}

		//判断当前会话是否为admin
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
	}
func updateUser(ctx *iris.Context) {}
func getUserById(ctx *iris.Context) {}
