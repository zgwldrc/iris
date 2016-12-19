// api for /account-type
package api

import (
    "gopkg.in/kataras/iris.v4"
    "iris/services/mysql"
    "iris/models"
)

func init(){
    iris.Get("/account-type", getAccountTypeList)
    iris.Delete("/account-type/:id", delAccount)
}

func getAccountTypeList (c *iris.Context) {
    var atl []models.AccountType
    mysql.DB.Select("id, type").Find(&atl)
    c.JSON(iris.StatusOK, models.JSONResponse{
        Data: atl,
    })
    return
}

func delAccount (c *iris.Context) {
    if id:= c.Param("id"); id == "" {
        return
    }

}