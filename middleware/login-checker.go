package middleware

import (
	"fmt"

	"gopkg.in/kataras/iris.v4"
	"iris/models"
)
func init() {
	iris.UseFunc(LoginStatusCheck)
}
func LoginStatusCheck(ctx *iris.Context) {
	message := "[全局登录状态检测中间件]:\n"
	message += "请求URI: " + ctx.PathString() + "\n"
	message += "请求携带Cookie: " + ctx.GetCookie("irissessionid") + "\n"

	switch ctx.PathString() {
	case "/login":
		ctx.Next()
	case "/createUser":
		ctx.Next()
	default:
		logined,err:=ctx.Session().GetInt("logined")
		if err != nil {
			iris.Logger.Println(err)
		}

		if logined != 1 {
			message += "Error: Session 失效或未登录"
			ctx.JSON(iris.StatusForbidden, models.JSONResponse{
				Message:"Session 失效或未登录",
			})
		} else {
			ctx.Next()
		}
	}

	fmt.Println(message)
}
