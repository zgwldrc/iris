package api

import (
    "github.com/kataras/iris"
    "fmt"
)

func init(){
    iris.Get("/accountTypeList", func(c *iris.Context) {
        a := "123"
        c.SetHeader("Access-Control-Allow-Origin","*")
        err := c.JSON(iris.StatusOK, a)
        if err != nil {
            fmt.Println(err)
        }
    })
}
