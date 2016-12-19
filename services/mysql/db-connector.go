package mysql

import (
	"fmt"

	_"github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jinzhu/gorm"
    "gopkg.in/kataras/iris.v4"

)

//这个模块提供一个DB对象的指针，在init()初始化，在main函数中释放
var DB *gorm.DB

func init(){
    var err error
    DB, err = gorm.Open("mysql", iris.Config.Other["MySQLDSN"])
    if err != nil {
        fmt.Println(err)
    }
}



