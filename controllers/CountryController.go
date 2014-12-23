package controllers

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/rtanwar/bee_test/models"
)

type CountryController struct {
	MainController
}

func (c *CountryController) Get_countries() {
	// c.user = c.GetSession(sessionName)
	// fmt.Printf("Session %s", v)
	// if v == nil {
	// 	c.Ctx.Redirect(302, "/login")
	// 	// this.SetSession("asta", int(1))
	// 	// this.Data["num"] = 0
	// }
	c.Data["Countries"], _ = models.GetAllCountry()
	// fmt.Println(c.Data["Countries"])
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["user"] = UserName
	c.TplNames = "index.tpl"
}

func (c *CountryController) Get_country() {
	country := c.Ctx.Input.Param(":id")
	fmt.Printf("Country Controller %s", country)
	// country := c.GetString("id")
	if country != "" {

		c.Data["country"], _ = models.GetCountry(country)
		c.Data["user"] = UserName
	}
}