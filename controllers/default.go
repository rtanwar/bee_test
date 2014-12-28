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
	user interface{}
}

func (c *MainController) Prepare() {
	fmt.Println("CONTROLLER PREPARE")
	sess := c.GetSession(sessionName)
	fmt.Printf("Session %s", sess)
	m := sess.(map[string]interface{})
	c.user = m["user"]
	//    c.Data["Username"] = m["company"]
	//c.user = c.GetSession(sessionName)
}

var sessionName = beego.AppConfig.String("SessionName")

func (c *MainController) Get() {

	// fmt.Printf("Session %s", v)
	// if v == nil {
	// 	c.Ctx.Redirect(302, "/login")
	// 	// this.SetSession("asta", int(1))
	// 	// this.Data["num"] = 0
	// }
	country := ""
	c.Data["Countries"], _ = models.GetAllCountry(country)
	// fmt.Println(c.Data["Countries"])
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["user"] = c.user
	c.TplNames = "index.tpl"
}

// var sessionName = beego.AppConfig.String("SessionName")

// var FilterUser = func(ctx *context.Context) {
// 	_, ok := ctx.Input.Session(sessionName).(int)
// 	if !ok && ctx.Input.Uri() != "/login" && ctx.Input.Uri() != "/register" {
// 		ctx.Redirect(302, "/login")
// 	}
// }
