package main

import (
    "gopkg.in/kataras/iris.v4"
    _ "iris/config"
    _ "iris/middleware"
    _ "iris/api"
    "iris/services/mysql"
)

func main() {
    defer mysql.DB.Close()
    iris.Listen("localhost:8080")
}

