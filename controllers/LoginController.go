package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/rtanwar/bee_test/models"
	"github.com/rtanwar/bee_test/utils"
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

func (this *LoginController) Register() {
	this.Data["message"] = ""
	if this.Ctx.Request.Method == "POST" {
		username := this.GetString("identity")
		password := this.GetString("password")
		passwordre := this.GetString("passwordre")
		test := models.RegisterForm{Username: username, Password: password, PasswordRe: passwordre}
		valid := validation.Validation{}
		b, _ := valid.Valid(&test)
		// if err != nil {
		// }
		if !b {
			m := ""
			for _, err := range valid.Errors {
				fmt.Println(err.Key, err.Message)
				m += "<br>" + err.Key + ":" + err.Message
			}
			this.Data["message"] = m
		} else {
			salt := utils.GetRandomString(10)
			encodedPwd := salt + "$" + utils.EncodePassword(password, salt)

			user := new(models.Users)
			user.Username = username
			user.Password = encodedPwd
			user.Salt = salt
			user.Email = username
			models.AddUsers(user)
			// o.Insert(user)

			this.Redirect("/", 302)
		}
		fmt.Printf("%s", this.Data["message"])
	}
	this.TplNames = "auth/register.tpl"
}

func (c *LoginController) Post() {
	identity := c.GetString("identity")
	password := c.GetString("password")

	// c.Data["Website"] = "beego.me"
	// c.Data["email"] = "astaxie@gmail.com"
	c.Data["identity"] = identity
	// c.Data["password"] = password
	// fmt.Printf("Login POST %s %s", identity, password)

	if models.CheckUser(identity, password) {
		// if (identity == "demo") && (password == "demo") {
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
	c.Data["message"] = "Invalid combination!!"
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
