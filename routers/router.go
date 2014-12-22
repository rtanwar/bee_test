package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/auth"
	"github.com/rtanwar/bee_test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/login", &controllers.LoginController{})
	// beego.Router("/login", &controllers.LoginController{}, "get:Get;post:Post")
	beego.Router("/login", &controllers.LoginController{}, "get:Get;post:Post")
	// beego.InsertFilter("/*", beego.BeforeRouter, controllers.FilterUser)
	// beego.InsertFilter("*", beego.BeforeRouter, auth.Basic("username", "secretpassword"))

	beego.InsertFilter("*", beego.BeforeRouter, auth.Basic("username", "secretpassword"))
	// beego.InsertFilter("/*", beego.BeforeRouter, FilterUser) //woring

	// beego.AutoRouter(&controllers.LoginController{})
}
