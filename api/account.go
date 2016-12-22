//api for /account
package api

import (
	"gopkg.in/kataras/iris.v4"
	"iris/models"
	"fmt"
	"iris/custerr"
)

func init() {
	//新增帐户
	iris.Post("/account", createAccount)
	//获取账户列表
	iris.Get("/account", getAccountList)
	//获取账户信息
	iris.Get("/account/:id", getAccount)
	//更新账户
	iris.Put("/account/:id", updateAccount)
	//删除账户
	iris.Delete("/account/:id", deleteAccount)
}

func getAccount(ctx *iris.Context){
	var Account  models.Account
	aid ,_:= ctx.ParamInt("id")
	uid, err := ctx.Session().GetInt("id")
	if err != nil {
		ctx.JSON(iris.StatusForbidden,models.JSONResponse{
			Message: fmt.Sprint(err),
		})
		return
	}
    Account.Load(uid,aid)
	ctx.JSON(iris.StatusOK, models.JSONResponse{
		Data: Account,
	})
}

func getAccountList(ctx *iris.Context){
	var a models.Account
	var AccountList  []models.Account
	uID, err := ctx.Session().GetInt("id")
	if err != nil {
		ctx.JSON(iris.StatusForbidden,models.JSONResponse{
			Message: fmt.Sprint(err),
		})
		return
	}

	AccountList = a.GetListByUserId(uID,"created_at")
	ctx.JSON(iris.StatusOK, models.JSONResponse{
		Data: AccountList,
	})
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
			ErrorCode: custerr.ERR_ACCOUNT_BROKEN_INPUT,
			Message:fmt.Sprint(err),
		})
		return
	}
	inputAccount.Create()
	if inputAccount.ID == 0 {
		ctx.JSON(iris.StatusBadRequest, models.JSONResponse{
			ErrorCode: custerr.ERR_ACCOUNT_CREATE_FAILED,
			Message: "账户创建失败",
		})
		return
	}

	ctx.JSON(iris.StatusOK, models.JSONResponse{
		Message: "创建成功",
	})
}

func updateAccount(ctx *iris.Context){
	var inputAccount  models.Account

	if aid, err := ctx.ParamInt("id"); err != nil {
		iris.Logger.Println(err)
		return
	} else {
		inputAccount.ID = aid
	}

	if uid, err := ctx.Session().GetInt("id"); err != nil {
		iris.Logger.Println(err)
		ctx.JSON(iris.StatusForbidden,models.JSONResponse{
			Message: fmt.Sprint(err),
		})
		return
	} else {
		inputAccount.UserID = uid
	}

	//装载JSON数据
	if err := ctx.ReadJSON(&inputAccount); err != nil {
		iris.Logger.Println(err)
		ctx.JSON(iris.StatusBadRequest, models.JSONResponse{
			ErrorCode: custerr.ERR_ACCOUNT_BROKEN_INPUT,
			Message:fmt.Sprint(err),
		})
		return
	}

	inputAccount.Update()
}

func deleteAccount(ctx *iris.Context){
	var inputAccount  models.Account
	aid ,_:= ctx.ParamInt("id")
	inputAccount.ID = aid
	inputAccount.Delete()
}