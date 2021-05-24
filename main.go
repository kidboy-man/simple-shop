package main

import (
	"github.com/kidboy-man/simple-shop/database"
	_ "github.com/kidboy-man/simple-shop/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	database.InitDB()
	beego.Run()
}
