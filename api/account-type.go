// api for /account-type
package api

import (
    "gopkg.in/kataras/iris.v4"
    "fmt"
)

func init(){
    iris.Get("/account-type", func(c *iris.Context) {
        a := "123"
        c.SetHeader("Access-Control-Allow-Origin","*")
        err := c.JSON(iris.StatusOK, a)
        if err != nil {
            fmt.Println(err)
        }
    })
}
