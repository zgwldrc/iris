package middleware

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/logger"
)
func init() {
	    //给iris 配置访问日志的打印
    errorLogger := logger.New(logger.Config{
        // Status displays status code
        Status: true,
        // IP displays request's remote address
        IP: true,
        // Method displays the http method
        Method: true,
        // Path displays the request path
        Path: true,
    })
    iris.Use(errorLogger)

    iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
        errorLogger.Serve(ctx)
        ctx.JSON(404,"{Not Found}")
    })
}
