package routers

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/plugins/auth"
	"github.com/rtanwar/bee_test/controllers"
)

func init() {
	beego.Router("/", &controllers.CountryController{}, "get:Get_countries")
	beego.Router("/json", &controllers.CountryController{}, "get:Get_countries_json")

	beego.Router("/country/:id:string", &controllers.CountryController{}, "get:Get_country")
	// beego.Router("/login", &controllers.LoginController{})
	// beego.Router("/login", &controllers.LoginController{}, "get:Get;post:Post")
	beego.Router("/login", &controllers.LoginController{}, "get:Get;post:Post")
	beego.Router("/logout", &controllers.LoginController{}, "get:LogOut")
	// beego.InsertFilter("/*", beego.BeforeRouter, controllers.FilterUser)
	// beego.InsertFilter("*", beego.BeforeRouter, auth.Basic("username", "secretpassword"))

	// beego.InsertFilter("*", beego.BeforeRouter, auth.Basic("username", "secretpassword"))
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser) //working
	beego.InsertFilter("/", beego.BeforeRouter, FilterUser)  //working

	// beego.AutoRouter(&controllers.LoginController{})
}
