package middleware

import "gopkg.in/kataras/iris.v4"
import "gopkg.in/iris-contrib/middleware.v4/cors"

func init(){
	crs := cors.New(cors.Options{})
	iris.Use(crs)
}