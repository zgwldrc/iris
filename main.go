package main

import (
    "gopkg.in/kataras/iris.v4"
    _ "iris/config"
    _ "iris/middleware"
    _ "iris/api"
    "iris/services/mysql"
    "flag"
    "iris/init-db"
)

func init() {
    var initDB = flag.Bool("initdb", false, "whether init db")
    if flag.Parse(); *initDB {
        init_db.Init(mysql.DB)
    }
}

func main() {

    defer mysql.DB.Close()
    iris.Listen("localhost:8080")
}

