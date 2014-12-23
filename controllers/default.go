package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/rtanwar/bee_test/models"
)

type User struct {
	name string
	id   int
}

type MainController struct {
	beego.Controller
}

var sessionName = beego.AppConfig.String("SessionName")

func (c *MainController) Get() {
	v := c.GetSession(sessionName)
	fmt.Printf("Session %s", v)
	if v == nil {
		c.Ctx.Redirect(302, "/login")
		// this.SetSession("asta", int(1))
		// this.Data["num"] = 0
	}
	c.Data["Countries"], _ = models.GetAllCountry()
	fmt.Println(c.Data["Countries"])
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["user"] = v
	c.TplNames = "index.tpl"
}

// var sessionName = beego.AppConfig.String("SessionName")

// var FilterUser = func(ctx *context.Context) {
// 	_, ok := ctx.Input.Session(sessionName).(int)
// 	if !ok && ctx.Input.Uri() != "/login" && ctx.Input.Uri() != "/register" {
// 		ctx.Redirect(302, "/login")
// 	}
// }
