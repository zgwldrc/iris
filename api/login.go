package api

import (
    "github.com/kataras/iris"
    "../models"
)

func init(){
    iris.Post("/login", func(c *iris.Context) {
        c.SetHeader("Access-Control-Allow-Origin","*")
        // store user input info
        var ruser= models.User{}

        // store db query result
        var user = models.User{}
        err := c.ReadJSON(&ruser)
        if err != nil {
            c.JSON(405, "Input String can not be recognised")
            return
        }

        models.DB.Where("name = ?", ruser.Name).First(&user)
        if user.Name == "" {
            c.JSON(200, "UserName do not exists")
            return
        } else {
            if user.Password == ruser.Password {

            } else {
                c.Write("Password is not correct!")
            }
        }
    })
}

