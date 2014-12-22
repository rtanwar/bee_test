package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "auth/login.tpl"
	// jsoninfo := c.GetString("jsoninfo")
}

func (c *LoginController) Post() {
	var sessionName = beego.AppConfig.String("SessionName")
	identity := c.GetString("identity")
	password := c.GetString("password")

	// c.Data["Website"] = "beego.me"
	// c.Data["email"] = "astaxie@gmail.com"
	c.Data["identity"] = identity
	c.Data["password"] = password

	fmt.Printf("Login POST %s %s", identity, password)
	if (identity == "demo") && (password == "demo") {
		c.SetSession(sessionName, 123456)
	}

	c.TplNames = "auth/login.tpl"
	// this.Ctx.Redirect(302, "/admin/index")
}

func init() {
	// v := this.GetSession("user")
	// fmt.Printf("From Login init session[user]=%+v\n", v)

}
