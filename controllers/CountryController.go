package controllers

import (
	"encoding/json"
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/rtanwar/bee_test/models"
	"log"
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

func (c *CountryController) Delete_contry() {
	c.Data["json"] = "{}"
	country := c.Ctx.Input.Param(":id")
	if len(country) != 0 {
		models.DeleteCountry(country)
	}
	c.ServeJson()
}

func (c *CountryController) Get_countries_json() {

	country := c.Ctx.Input.Param(":id")
	fmt.Printf("Country Controller %s", country)
	c.Data["user"] = c.user
	c.Data["json"] = "{}"
	var err error
	var str_err string
	// country := c.GetString("id")

	if c.Ctx.Input.IsPost() {
		// jsoninfo := c.GetString("jsoninfo")
		fmt.Printf("JSON %s", c.Ctx.Input.CopyBody())
		new_c := models.Country{}
		json.Unmarshal(c.Ctx.Input.CopyBody(), &new_c)
		valid := validation.Validation{}
		valid.Required(new_c.Id, "Id")
		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
				str_err += err.Key + ": " + err.Message + "\n"
			}
			c.CustomAbort(500, str_err)
			// c.Data["json"] = str_err
			// c.ServeJson()
			return
			// c.Data["json"] = validationerr
		}
		fmt.Printf("Struct %s", new_c)
		insert := len(country) == 0
		_, err := models.SaveCountry(new_c, insert)
		if err != nil {
			c.Data["json"] = err.Error()
		}

		c.ServeJson()

		// 	v := Country{Id: m.Id}
		//
	} else {
		var cdetail interface{}
		if len(country) != 0 {
			if strings.ToUpper(country) == "ALL" {
				fmt.Println("ALL DATA")
				cdetail, err = models.GetAllCountry("")
			} else {
				cdetail, err = models.GetCountry(country)
			}
			if err == nil {
				c.Data["json"] = cdetail
			}
			c.ServeJson()
		}
	}

	// country := ""
	// var err error
	// c.Data["json"], err = models.GetAllCountry(country)
	// if err != nil {

	// } else {
	// 	c.ServeJson()
	// }
}
