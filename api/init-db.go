// api for /init-db
package api

import (
    "gopkg.in/kataras/iris.v4"
    "iris/services/mysql"
	"iris/models"
	"iris/init-db"
)

func init() {
	iris.Post("/db/init", func(ctx *iris.Context) {
		init_db.Init(mysql.DB)
		ctx.JSON(iris.StatusOK, models.JSONResponse{
			Message:"InitDB Complete",
		})
	})
}
