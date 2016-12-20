package config

import (
    "gopkg.in/kataras/iris.v4"
    "os"
    "encoding/json"
    "fmt"
)

//该模块用于解析配置文件
func init() {
    //定义一个配置结构体的map，分别存储开发，测试，生产的不同配置
    configMap := map[string]iris.Configuration{}
    //从配置文件加载配置数据
    cf, errOpen := os.Open(`config.json`)
    if errOpen != nil {
        fmt.Println("Open config file custerr:", errOpen)
    }
    //关闭配置文件
    defer func(){
        if err := cf.Close(); err != nil {
            panic(err)
        }
    }()

    //配置文件格式为json，需要一个json Decoder，并调用其Decode方法
    errDecode := json.NewDecoder(cf).Decode(&configMap)
    if errDecode != nil {
        fmt.Println("Decoding config.json:", errDecode)
    }

    //加载配置之后，先合并基础配置，再合并特定环境的配置
    configMap["BASIC"].Set(iris.Config)

    //从环境变量获取相应的环境配置
    env := os.Getenv("IRIS_CONFIG_ENV")
    if  env == "" {
        env = "DEV"
    }
    configMap[env].Set(iris.Config)
    fmt.Printf("########当前使用配置%s#########\n",env)
    fmt.Println("iris.Config.Gzip:",iris.Config.Gzip)
    fmt.Println("iris.Config.IsDevelopment:",iris.Config.IsDevelopment)
    fmt.Println(`iris.Config.Other["MySQLDSN"]:`,iris.Config.Other["MySQLDSN"])
}

