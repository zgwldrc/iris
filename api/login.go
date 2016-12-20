// api for /login
package api

import (
    "gopkg.in/kataras/iris.v4"
    "iris/models"
    "iris/custerr"
    "fmt"
)

func init(){
    iris.Post("/login", login)("login")
    iris.Delete("/login", logout)("logout")
}

func login(ctx *iris.Context) {

    //读取传入数据到模型，用于判断用户合法性
    var inputUser = models.User{}

    if err := ctx.ReadJSON(&inputUser); err != nil {
        ctx.JSON(iris.StatusBadRequest, models.JSONResponse{
            ErrorCode: custerr.ERR_USER_BROKEN_INPUT,
            Message: "Input can not be recognised",
        })
        return
    }

    if iris.Config.IsDevelopment {
        fmt.Println("[login] user input :", inputUser.Name, inputUser.Password)
    }

    //返回一个CheckResult结构体存储检查结果
    checkResult := inputUser.CheckUserValidity()

    //用户不存在
    if checkResult.NotExist {
        ctx.JSON(iris.StatusBadRequest, models.JSONResponse{
            ErrorCode: custerr.ERR_USER_NOT_EXISTS,
            Message: "User not exists",
        })
        return
    }
    //密码无效
    if checkResult.PasswordInvalid {
        ctx.JSON(iris.StatusBadRequest, models.JSONResponse{
            ErrorCode: custerr.ERR_USER_PASSWD_INVALID,
            Message:"Password Invalid",
        })
        return
    }

    sess := ctx.Session()
    sess.Set("logined", 1)
    sess.Set("user", inputUser.Name)
    sess.Set("id", checkResult.ID)

    if iris.Config.IsDevelopment {
        name := sess.GetString("user")
        id,_:=sess.GetInt("id")
        fmt.Println("Session for:",name,id)
    }
    //响应登录成功
    ctx.JSON(iris.StatusOK, models.JSONResponse{
        Message:"Login Success",
        Data:checkResult,
    })
}

func logout(ctx *iris.Context) {
    ctx.SessionDestroy()
    ctx.JSON(iris.StatusOK, models.JSONResponse{
        Message:"You have been loged out",
    })
}

