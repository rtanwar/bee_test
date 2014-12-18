package routers

import (
	"github.com/astaxie/beego"
	"github.com/rtanwar/bee_test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/login", &controllers.LoginController{})
	beego.Router("/login", &controllers.LoginController{}, "get:Get;post:Post")
	// beego.AutoRouter(&controllers.LoginController{})
}
