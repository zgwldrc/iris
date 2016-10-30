package main

import (
    "github.com/kataras/iris"
    _ "iris/config"
    _ "iris/modules/middleware"
    _ "iris/api"
    "iris/modules"
)

func main() {
    defer modules.DB.Close()
    iris.Listen("localhost:8080")
}

