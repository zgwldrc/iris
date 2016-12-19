//api for /account
package api

import (
	"gopkg.in/kataras/iris.v4"
	"iris/models"
	"fmt"
	"iris/error"
)

func init() {
	iris.Post("/account", createAccount)
}

func createAccount(ctx *iris.Context){
	var inputAccount  models.Account
	id, err := ctx.Session().GetInt("id")
	if err != nil {
		ctx.JSON(iris.StatusForbidden,models.JSONResponse{
			Message: fmt.Sprint(err),
		})
		return
	}
	inputAccount.UserID = id

	if iris.Config.IsDevelopment{
		fmt.Printf(
			"[Post /account]: \n操作用户:%s\n", ctx.Session().GetString("user"),
		)
	}
	//装载JSON数据
	if err := ctx.ReadJSON(&inputAccount); err != nil {
		iris.Logger.Println(err)
		ctx.JSON(iris.StatusBadRequest, models.JSONResponse{
			ErrorCode: error.ERR_ACCOUNT_BROKEN_INPUT,
			Message:fmt.Sprint(err),
		})
		return
	}

	//检验数据

	inputAccount.Create()
	if inputAccount.ID == 0 {

	}
}