package api

import (
    "github.com/kataras/iris"
    "iris/modules"
	"iris/models"
)

func init() {
    iris.Post("/db/init", func(ctx *iris.Context) {
		//判断是否是admin用户，否则返回403
        if ctx.Session().GetString("isAdmin") == "true" {
			modules.Initdb(modules.DB)
			ctx.JSON(iris.StatusOK,models.JSONResponse{
				Message:"InitDB Complete",
			})
		} else {
			ctx.JSON(iris.StatusForbidden, models.JSONResponse{
				Message: "Your credential is not Allowed to do this! :-)",
			})
		}
    })
}
