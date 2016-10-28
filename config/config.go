package config

import (
    "github.com/kataras/iris"
    "github.com/iris-contrib/middleware/logger"
    "os"
    "encoding/json"
    "fmt"
)


type ConfigFile struct {
    Gzip bool
}
var Config = ConfigFile{Gzip: true}
func init(){

    cf, errOpen := os.Open(`C:\Users\xiayu\PycharmProjects\iris2\config\config.json`)
    if errOpen != nil {
        fmt.Println("Open config file error:", errOpen)
    }
    defer func(){
        if err := cf.Close(); err != nil {
            panic(err)
        }
    }()

    //buf := make([]byte, 1024)
    //n, errRead := cf.Read(buf)
    //if errRead != nil {
    //    fmt.Println("Read config file error:", errRead)
    //}
    //
    //fmt.Println("has read %s and the buf like %v",n,buf)

    decoder := json.NewDecoder(cf)

    errDecode := decoder.Decode(&Config)
    if errDecode != nil {
        fmt.Println("myerror:", errDecode)
    }


    //  ENCODE TEST
    //encoder := json.NewEncoder(cf)
    //errEncode := encoder.Encode(Config)
    //if errEncode != nil {
    //    fmt.Println("errEncode: ",errEncode)
    //}

    errorLogger := logger.New(logger.Config{
        // Status displays status code
        Status: true,
        // IP displays request's remote address
        IP: true,
        // Method displays the http method
        Method: true,
        // Path displays the request path
        Path: true,
    })

    iris.Use(errorLogger)

    iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
        errorLogger.Serve(ctx)
        ctx.Write("My Custom 404 error page ")
    })
}

