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
	UserName := c.GetSession(sessionName)
	fmt.Printf("User%s", UserName)
	if UserName != nil {
		c.Ctx.Redirect(302, "/")
	} else {
		c.TplNames = "auth/login.tpl"
	}

	// jsoninfo := c.GetString("jsoninfo")
}

func (c *LoginController) Post() {
	identity := c.GetString("identity")
	password := c.GetString("password")

	// c.Data["Website"] = "beego.me"
	// c.Data["email"] = "astaxie@gmail.com"
	c.Data["identity"] = identity
	c.Data["password"] = password

	fmt.Printf("Login POST %s %s", identity, password)
	if (identity == "demo") && (password == "demo") {
		m := make(map[string]interface{})
		//   	m["company"] = a.Company
		m["user"] = identity
		//   	m["timestamp"] = time.Now()
		//   	c.SetSession("foo", m)
		// c.SetSession(sessionName, identity)
		c.SetSession(sessionName, m)
		// UserName = identity
		c.Ctx.Redirect(302, "/")
	}
	c.TplNames = "auth/login.tpl"
	// this.Ctx.Redirect(302, "/admin/index")
}

func (c *LoginController) LogOut() {
	c.DestroySession()
	c.Ctx.Redirect(302, "/")
}
func init() {
	// v := this.GetSession("user")
	// fmt.Printf("From Login init session[user]=%+v\n", v)

}
