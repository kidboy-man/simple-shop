package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/kidboy-man/simple-shop/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
