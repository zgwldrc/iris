package main

import (
    "github.com/kataras/iris"
    "./models"
    _ "./api"
    _"./config"
)

func main() {
    defer models.DB.Close()
    iris.Listen("localhost:8080")
}

