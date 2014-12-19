package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	v := c.GetSession("user")
	if v == nil {
		c.Ctx.Redirect(302, "/login")
		// this.SetSession("asta", int(1))
		// this.Data["num"] = 0
	}

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["user"] = v
	c.TplNames = "index.tpl"
}
