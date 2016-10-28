package api

import (
    "../models"
    "github.com/kataras/iris"
)

func init() {
    iris.Get("/initDB", func(c *iris.Context) {
        models.Initdb(models.DB)
        c.Write("InitDB Ok")
    })
}
