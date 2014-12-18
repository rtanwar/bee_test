package routers

import (
	"github.com/rtanwar/bee_test1/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
