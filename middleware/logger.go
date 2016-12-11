package middleware

import (
	"gopkg.in/kataras/iris.v4"
	"gopkg.in/iris-contrib/middleware.v4/logger"
)
func init() {
	    //给iris 配置访问日志的打印
    errorLogger := logger.New()
    iris.Use(errorLogger)

    iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
        errorLogger.Serve(ctx)
        ctx.JSON(404,"{Not Found}")
    })
}
