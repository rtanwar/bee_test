package controllers

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/rtanwar/bee_test/models"
	"strings"
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
	country := c.GetString("country")
	c.Data["Countries"], _ = models.GetAllCountry(country)
	// fmt.Println(c.Data["Countries"])
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["country"] = country

	c.Data["user"] = c.user
	c.TplNames = "index.tpl"
}

func (c *CountryController) Get_country() {
	country := c.Ctx.Input.Param(":id")
	fmt.Printf("Country Controller %s", country)
	// country := c.GetString("id")
	if country != "" {
		cdetail, err := models.GetCountry(country)
		if err == nil {
			c.Data["country"] = cdetail
		}
		c.Data["user"] = c.user
	}
}

func (c *CountryController) Get_countries_json() {
	country := c.Ctx.Input.Param(":id")

	fmt.Printf("Country Controller %s", country)
	c.Data["user"] = c.user
	// country := c.GetString("id")
	var cdetail interface{}
	var err error
	if country != "" {

		if strings.ToUpper(country) == "ALL" {
			fmt.Println("ALL DATA")
			cdetail, err = models.GetAllCountry("")
		} else {
			cdetail, err = models.GetCountry(country)
		}
		if err == nil {
			c.Data["json"] = cdetail
		} else {
			c.Data["json"] = "{}"
		}
		c.ServeJson()
	}

	// country := ""
	// var err error
	// c.Data["json"], err = models.GetAllCountry(country)
	// if err != nil {

	// } else {
	// 	c.ServeJson()
	// }
}
