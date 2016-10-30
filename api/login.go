package api

import (
    "github.com/kataras/iris"
    "iris/models"
    "iris/modules"
    "fmt"
)

func init(){
    //主要根据用户合法性来Set－Cookie
    iris.Post("/login", func(ctx *iris.Context) {
        fmt.Println("Entered login api...")
        ctx.SetHeader("Access-Control-Allow-Origin","*")
        fmt.Println("Setting cors Header...Access-Control-Allow-Origin: *")
        //读取传入数据到模型，用于判断用户合法性
        var inputUser = models.User{}
        err := ctx.ReadJSON(&inputUser)
        if err != nil {
            ctx.JSON(405, "Input String can not be recognised")
            return
        }

        fmt.Println("user input :", inputUser)
        //返回一个CheckResult结构体存储检查结果
        checkResult := modules.NewUserChecker(&inputUser).CheckUserValidity()
        //若用户合法则设置相应的cookie
        if checkResult.PasswordValid {
            sess := ctx.Session()
            if checkResult.IsAdmin {
                sess.Set("isAdmin", "true")
            }
            sess.Set("logined", 1)
            sess.Set("name", inputUser.Name)
            sess.Set("id", inputUser.ID)
            fmt.Println("Session for:",sess.GetString("name"),sess.GetString("id"))
            //响应登录成功
            ctx.JSON(iris.StatusOK, models.JSONResponse{
                Message:"Login Success",
                Data:checkResult,
            })
            if err != nil {
                iris.Logger.Println(err)
            }
        } else {
            //密码无效，登录失败
            ctx.JSON(iris.StatusOK, models.JSONResponse{
                Message:"Login Failed",
                Data:checkResult,
            })
            if err != nil {
                iris.Logger.Println(err)
            }
        }
    })

    iris.Delete("/login", func(ctx *iris.Context) {
        ctx.SessionDestroy()
        ctx.JSON(iris.StatusOK, models.JSONResponse{
            Message:"You have been loged out",
        })
    })
}

