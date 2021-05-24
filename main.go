package main

import (
	"github.com/kidboy-man/simple-shop/database"
	_ "github.com/kidboy-man/simple-shop/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	database.InitDB()

	// option := cors.Options{
	// 	AllowMethods:     []string{"*"},
	// 	AllowHeaders:     []string{"*"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowOrigins:     []string{"*"},
	// 	AllowCredentials: true,
	// }

	// beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&option))

	beego.Run()
}
