package middleware

import (
	"github.com/kataras/iris"
	"fmt"
	"iris/models"
)

type LoginedChecker struct {

}

func (lc LoginedChecker) Serve(ctx *iris.Context) {
	message := "[全局登录状态检测过滤器]:\n"
	message += "请求URI: " + ctx.PathString() + "\n"
	message += "请求携带Cookie: " + ctx.GetCookie("irissessionid") + "\n"

	switch ctx.PathString() {
	case "/login":
		ctx.Next()
	case "/createUser":
		ctx.Next()
	default:
		if ctx.Session().GetInt("logined") != 1 {
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

func init(){
	iris.Use(&LoginedChecker{})
}